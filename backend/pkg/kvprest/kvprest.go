package kvprest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// define a global in-memory map of key/value pairs we'll use as our data store.
var keyvaluepairs map[string]interface{}
var mapMutex sync.Mutex

func init() {
	keyvaluepairs = make(map[string]interface{})
}

// defines the format used for creating a kvp
type keyValuePair struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

// returns a list all key/value pairs stored in our map
func ListKeyValuePairs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("listKeyValuePairs called")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	mapMutex.Lock()
	defer mapMutex.Unlock()
	json.NewEncoder(w).Encode(keyvaluepairs)
}

// returns the value for the specified key
func GetValueForKey(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getValueForKey called")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")

	vars := mux.Vars(r)
	key := vars["key"]

	if len(key) > 0 {
		mapMutex.Lock()
		defer mapMutex.Unlock()
		json.NewEncoder(w).Encode(keyvaluepairs[key])
	}
}

// sets the specified key/value pair
func SetValueForKey(w http.ResponseWriter, r *http.Request) {
	fmt.Println("setValueForKey called")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")

	reqBody, _ := ioutil.ReadAll(r.Body)
	fmt.Fprintf(w, "%+v", string(reqBody))

	var kvp keyValuePair
	json.Unmarshal(reqBody, &kvp)

	if len(kvp.Key) > 0 {
		mapMutex.Lock()
		defer mapMutex.Unlock()
		keyvaluepairs[kvp.Key] = kvp.Value
	}
}

// deletes the key/value pair from our map
func DeleteKey(w http.ResponseWriter, r *http.Request) {
	fmt.Println("deleteKey called")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")

	vars := mux.Vars(r)
	key := vars["key"]

	if len(key) > 0 {
		mapMutex.Lock()
		defer mapMutex.Unlock()
		delete(keyvaluepairs, key)
	}
}