# Parquet Search Engine

A Simple fullstack app to upload files and search through the records.

## Tech Stack
- Backend: Golang
- FrontEnd: React
- Libraries:
   - Golang: `parquet-go`
   - React: `axios`
--------------------------------------------------------------------------------------------------------------------------------------------------------------------


## Architecture 
- The backend loads uploaded files in-memory and parses them into Go structs.
- The frontend allows uploading files and keyWord searches.
- No Database or external search engine used.
--------------------------------------------------------------------------------------------------------------------------------------------------------------------


## setup
clone

## Backend 
- cd backend
- go mod tidy
- go run main.go

## Frontend
- cd frontend
- npm install
- npm start
