package kvprest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/FramnkRulez/keyvalueproj/pkg/keyvaluestore"
	"github.com/gorilla/mux"
)

// defines the format used for creating a kvp
type keyValuePair struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func addCorsHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
}

func SetupHandlers(router *mux.Router) {
	// handle CORS preflight requests
	router.HandleFunc("/", HandlePreflight).Methods("OPTIONS")
	router.HandleFunc("/kvp", HandlePreflight).Methods("OPTIONS")
	router.HandleFunc("/kvp/{key}", HandlePreflight).Methods("OPTIONS")

	// list key/value pairs
	router.HandleFunc("/", ListKeyValuePairs)

	// get value for given key
	router.HandleFunc("/kvp/{key}", GetValueForKey).Methods("GET")

	// set key=value
	router.HandleFunc("/kvp", SetValueForKey).Methods("POST")

	// delete key/value pair
	router.HandleFunc("/kvp/{key}", DeleteKey).Methods("DELETE")
}

// this allows our test server to deal with CORS pre-flight request since
// we will be requesting on localhost but a different port from javascript.
func HandlePreflight(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CORS preflight request")

	addCorsHeaders(w)
}

// returns a list all key/value pairs stored in our map
func ListKeyValuePairs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listKeyValuePairs called")

	addCorsHeaders(w)

	json.NewEncoder(w).Encode(keyvaluestore.GetKeyValuePairs())
}

// returns the value for the specified key
func GetValueForKey(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getValueForKey called")

	addCorsHeaders(w)

	vars := mux.Vars(r)
	key := vars["key"]

	if len(key) > 0 {
		json.NewEncoder(w).Encode(keyvaluestore.GetValueForKey(key))
	}
}

// sets the specified key/value pair
func SetValueForKey(w http.ResponseWriter, r *http.Request) {
	fmt.Println("setValueForKey called")

	addCorsHeaders(w)

	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))

	var kvp keyValuePair
	json.Unmarshal(reqBody, &kvp)

	if len(kvp.Key) > 0 {
		keyvaluestore.SetValueForKey(kvp.Key, kvp.Value)
	}
}

// deletes the key/value pair from our map
func DeleteKey(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteKey called")

	addCorsHeaders(w)

	vars := mux.Vars(r)
	key := vars["key"]

	if len(key) > 0 {
		keyvaluestore.DeleteKey(key)
	}
}
