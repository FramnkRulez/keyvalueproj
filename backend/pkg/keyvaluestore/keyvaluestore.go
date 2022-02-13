package keyvaluestore

import (
	"sync"
)

// define a global in-memory map of key/value pairs we'll use as our data store.
var keyvaluepairs map[string]interface{}
var mapMutex sync.Mutex

func init() {
	keyvaluepairs = make(map[string]interface{})
}

func GetKeyValuePairs() map[string]interface{} {
	mapMutex.Lock()
	defer mapMutex.Unlock()
	return keyvaluepairs
}

func GetValueForKey(key string) interface{} {
	mapMutex.Lock()
	defer mapMutex.Unlock()
	return keyvaluepairs[key]
}

func SetValueForKey(key string, value interface{}) {
	mapMutex.Lock()
	defer mapMutex.Unlock()

	keyvaluepairs[key] = value
}

func DeleteKey(key string) {
	mapMutex.Lock()
	defer mapMutex.Unlock()
	delete(keyvaluepairs, key)
}
