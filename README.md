# Parquet Search Engine
 A Simple Full-Stack App for File Upload and Search
                 This is a full-stack web application designed to facilitate the upload, storage, and search of records within Parquet files. It provides a user-friendly interface to allow users to seamlessly interact with data without needing complex setups or external dependencies.

## Key Features:
**File Upload:** Users can easily upload Parquet files through the web interface.
**Search Functionality:** After uploading, users can search through the data using keywords or specific criteria to retrieve relevant records instantly.
**Backend in Go:** The backend is powered by Go (Golang), which efficiently handles file uploads, parsing of Parquet files, and serves API requests for searching data.
**Frontend in React:** The frontend is built using React, providing a clean and intuitive UI for users to interact with the application. It allows for seamless uploading of files and execution of search queries.
**Data Handling:** The Parquet files are loaded into memory and parsed into appropriate Go structs, making searching fast and efficient.
**In-Memory Search:** The application does not require a database or external search engine, as it leverages in-memory structures to perform quick searches on uploaded Parquet files.

## Benefits:
**Lightweight:** No heavy database systems are required. The entire search mechanism works within the application’s memory, making it lightweight and fast.
**Easy Integration:** The application is simple to set up and can be easily customized to work with other types of data files, making it a versatile tool for various use cases.
**Scalable:** While initially designed for Parquet files, the backend could be extended to support other file formats, and the search capabilities could be enhanced based on project needs.

## Tech Stack
- Backend: Golang
- FrontEnd: React
- Libraries:
   - Golang: `parquet-go`
   - React: `axios`
--------------------------------------------------------------------------------------------------------------------------------------------------------------------


## Architecture 
- The backend loads uploaded files in-memory and parses them into Go structs.
- The frontend allows uploading files and keyword searches.
- No Database or external search engine used.
--------------------------------------------------------------------------------------------------------------------------------------------------------------------


## Setup
clone the git repository in local using command
- git clone https://github.com/Anushreepd/parquet_search_engine.git

## Backend 
- cd backend
- go mod tidy
- go run main.go

## Frontend
- cd frontend
- npm install
- npm start

--------------------------------------------------------------------------------------------------------------------------------------------------------------------

## Images
![Screenshot 2025-04-27 152316](https://github.com/user-attachments/assets/4f3af14e-40aa-4e96-9cd0-b826940686dd)
![Screenshot 2025-04-27 152400](https://github.com/user-attachments/assets/4e6e6fb7-12e0-475c-9fd8-4e47c8afa039)

