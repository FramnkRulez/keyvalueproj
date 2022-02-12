package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type KeyValuePair struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

var keyvaluepairs map[string]interface{}

func listKeyValuePairs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listKeyValuePairs called")
	json.NewEncoder(w).Encode(keyvaluepairs)
}

func getValueForKey(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getValueForKey called")

	vars := mux.Vars(r)
	key := vars["key"]

	json.NewEncoder(w).Encode(keyvaluepairs[key])
}

func setValueForKey(w http.ResponseWriter, r *http.Request) {
	fmt.Println("setValueForKey called")

	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))

	var kvp KeyValuePair
	json.Unmarshal(reqBody, &kvp)

	keyvaluepairs[kvp.Key] = kvp.Value
}

func main() {
	fmt.Println("keyvalue backend server")

	keyvaluepairs = make(map[string]interface{})
	keyvaluepairs["test"] = "one"
	keyvaluepairs["test1"] = "two"
	keyvaluepairs["test2"] = 1

	router := mux.NewRouter().StrictSlash(true)

	// list key/value pairs
	router.HandleFunc("/", listKeyValuePairs)

	// get value for given key
	router.HandleFunc("/kvp/{key}", getValueForKey)

	// set key=value
	router.HandleFunc("/kvp", setValueForKey).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}
