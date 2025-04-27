package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/reader"
)

// EventData represents the structure for each event
type EventData struct {
	Message       string `json:"message"`
	Sender        string `json:"sender"`
	Event         string `json:"event"`
	EventId       string `json:"event_id"`
	NanoTimeStamp string `json:"nano_timestamp"`
}

var mockData = []EventData{
	{"User logged in", "John", "Login", "e001", "1627471500000"},
	{"User logged out", "Jane", "Logout", "e002", "1627471560000"},
	{"Password change", "John", "Update", "e003", "1627471620000"},
	{"Account created", "Alice", "Create", "e004", "1627471680000"},
}

// CORS middleware
func enableCORS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}
}

func parseParquetFile(filepath string) ([]EventData, error) {
	var events []EventData

	fr, err := local.NewLocalFileReader(filepath)
	if err != nil {
		return nil, err
	}
	defer fr.Close()

	pr, err := reader.NewParquetReader(fr, new(EventData), 4)
	if err != nil {
		return nil, err
	}
	defer pr.ReadStop()

	num := int(pr.GetNumRows())
	if num > 0 {
		if err := pr.Read(&events); err != nil {
			return nil, err
		}
	}
	return events, nil
}

func getAllEventsHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(mockData)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	query := r.URL.Query().Get("query")
	if query == "" {
		http.Error(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	var results []EventData
	for _, event := range mockData {
		if strings.Contains(strings.ToLower(event.Message), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(event.Sender), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(event.Event), strings.ToLower(query)) {
			results = append(results, event)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	// Enable CORS
	enableCORS(w, r)

	// Handle pre-flight OPTIONS request
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	// Only allow POST method for file upload
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Check if the request is JSON or multipart form
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "application/json") {
		// Handle JSON upload
		var events []EventData
		if err := json.NewDecoder(r.Body).Decode(&events); err != nil {
			http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
			return
		}
		// Add events to mockData
		mockData = append(mockData, events...)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf("Added %d events from JSON", len(events))))
		return
	}

	// Handle file upload (multipart form-data)
	err := r.ParseMultipartForm(10 << 20) // Limit the file size to 10MB
	if err != nil {
		http.Error(w, "Error parsing multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File not found", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Save the uploaded file temporarily
	tempPath := "./uploaded_" + handler.Filename
	tempFile, err := os.Create(tempPath)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}
	defer tempFile.Close()
	io.Copy(tempFile, file)

	// Parse the Parquet file to extract event data
	events, err := parseParquetFile(tempPath)
	if err != nil {
		http.Error(w, "Failed to parse Parquet: "+err.Error(), http.StatusInternalServerError)
		os.Remove(tempPath) // Remove the uploaded file in case of error
		return
	}
	os.Remove(tempPath) // Remove the uploaded file after processing

	// Add parsed events to mockData
	mockData = append(mockData, events...)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Added %d events from Parquet", len(events))))
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	enableCORS(w, r)
	if r.Method != http.MethodDelete {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	eventID := r.URL.Query().Get("event_id")
	if eventID == "" {
		http.Error(w, "Missing event_id", http.StatusBadRequest)
		return
	}

	found := false
	var updated []EventData
	for _, e := range mockData {
		if e.EventId != eventID {
			updated = append(updated, e)
		} else {
			found = true
		}
	}

	if !found {
		http.Error(w, "Event not found", http.StatusNotFound)
		return
	}

	mockData = updated
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Deleted event with ID %s", eventID)))
}

func main() {
	http.HandleFunc("/events", getAllEventsHandler)
	http.HandleFunc("/search", searchHandler)
	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/delete", deleteHandler)

	fmt.Println("🚀 Server started at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
