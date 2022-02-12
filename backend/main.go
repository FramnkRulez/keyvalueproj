package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/FramnkRulez/keyvalueproj/pkg/kvprest"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("keyvalue backend server")

	router := mux.NewRouter().StrictSlash(true)

	// list key/value pairs
	router.HandleFunc("/", kvprest.ListKeyValuePairs)

	// get value for given key
	router.HandleFunc("/kvp/{key}", kvprest.GetValueForKey).Methods("GET")

	// set key=value
	router.HandleFunc("/kvp", kvprest.SetValueForKey).Methods("POST")

	// delete key/value pair
	router.HandleFunc("/kvp/{key}", kvprest.DeleteKey).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
