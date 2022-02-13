# keyvalueproj - Demonstrates a React Key/Value frontend calling a REST backend written in Golang

This project consists of two main services: frontend (react app) and backend (go app).  When launched, the go service will be running at http://localhost:8080 and respond to REST GET, POST and DELETE requests for specified key/value pairs:

GET http://localhost:8080/ - list all key/value pairs currently in the key/value collection
GET http://localhost:8080/kvp/{key} - retrieve the value for the specified key
POST http://localhost:8080/kvp - create the key/value in the collection using the specified JSON body
DELETE http://localhost:8080/kbp/{key} - delete the key/value pair in the collection stored at key

# System Requirements
Requires Golang 1.11+ and Node.js 16.14+ to be installed locally on the system

# Installation 
In the backend folder, run 'go run main.go' and the first time it will download required modules

In the frontend folder, run 'npm install' to install required node modules.

# Running the backend (Go REST service)
To run the backend, run 'go run main.go' from the backend folder.  Backend uses port 8080 (this can be modified in main.go, however the frontend assumes port 8080)

# Running the frontend (React web interface)
To run the frontend, run 'npm start' from the frontend folder.  Frontend defaults to port 3000 but will prompt for an alternate port if unavailable.

# Testing
Go tests can be executed by running 'go test -v ./...' in the backend folder.
Additional testing of REST API can be done via Postman (GET, POST, DELETE)

# Usage
After running the backend and frontend modules, key/value pairs can be added via the web interface.  Each time a key/value pair is added the list will update and the key/value pair will be listed with the option to delete the pair.

If the server is stopped, or is not running when the web frontend tries to connect an error message will be displayed.  Additionally if the server is stopped as keys are added or deleted error messages will indicate the backend server could not be contacted.




