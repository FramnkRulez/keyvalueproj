# keyvalueproj - Demonstrates a React Key/Value frontend calling a REST backend written in Golang

This project consists of two main services: frontend (react app) and backend (go app).  When launched, the go service will be running at http://localhost:8080 and respond to REST GET, POST and DELETE requests for specified key/value pairs:

GET http://localhost:8080/ - list all key/value pairs currently in the key/value collection
GET http://localhost:8080/kvp/{key} - retrieve the value for the specified key
POST http://localhost:8080/kvp - create the key/value in the collection using the specified JSON body
DELETE http://localhost:8080/kbp/{key} - delete the key/value pair in the collection stored at key

# system requirements
Requires Golang 1.11+ and Node.js 16.14+ to be installed locally on the system

# Running the backend (Go REST service)
To run the backend, run 'go run main.go' from the backend folder.  Backend uses port 8080 (this can be modified in main.go, however the frontend assumes port 8080)

# Running the frontend (React web interface)
To run the frontend, run 'npm start' from the frontend folder.  Frontend defaults to port 3000 but will prompt for an alternate port if unavailable.




