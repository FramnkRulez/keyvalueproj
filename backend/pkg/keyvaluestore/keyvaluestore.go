// Package keyvaluestore provides an implementation of an in-memory key/value store
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
	if len(key) > 0 {
		mapMutex.Lock()
		defer mapMutex.Unlock()
		return keyvaluepairs[key]
	}

	return nil
}

func SetValueForKey(key string, value interface{}) {
	if len(key) > 0 {
		mapMutex.Lock()
		defer mapMutex.Unlock()
		keyvaluepairs[key] = value
	}
}

func DeleteKey(key string) {
	if len(key) > 0 {
		mapMutex.Lock()
		defer mapMutex.Unlock()
		delete(keyvaluepairs, key)
	}
}
