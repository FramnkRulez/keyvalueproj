# keyvalueproj - Demonstrates a React Key/Value frontend calling a REST backend written in Golang

This project consists of two main services: frontend (react app) and backend (go app).  When launched, the go service will be running at http://localhost:8080 and respond to REST GET, POST and DELETE requests for specified key/value pairs:

GET http://localhost:8080/ - list all key/value pairs currently in the key/value collection
GET http://localhost:8080/kvp/{key} - retrieve the value for the specified key
POST http://localhost:8080/kvp - create the key/value in the collection using the specified JSON body
DELETE http://localhost:8080/kbp/{key} - delete the key/value pair in the collection stored at key


