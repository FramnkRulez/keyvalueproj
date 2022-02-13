package keyvaluestore

import (
	"testing"
)

// Tests getting the list of key value pairs from the store
func TestGetKeyValuePairs(t *testing.T) {
	if len(GetKeyValuePairs()) > 0 {
		t.Fatalf(`Failed: expected initial length of keyvaluestore to be empty`)
	}

	SetValueForKey("test", "123")

	if len(GetKeyValuePairs()) != 1 {
		t.Fatalf(`Failed: expected single element in keyvaluestore`)
	}

	DeleteKey("test")

	if len(GetKeyValuePairs()) > 0 {
		t.Fatalf(`Failed: expected new length of keyvaluestore to be empty`)
	}
}

// tests getting an individual value by key
func TestGetValueForKey(t *testing.T) {
	val := GetValueForKey("test")
	if val != nil {
		t.Fatalf(`Failed: expected nil value for non-existant key`)
	}

	SetValueForKey("test", "123")

	val = GetValueForKey("test")
	if val != "123" {
		t.Fatalf(`Failed: expected '123' value for 'test'`)
	}

	DeleteKey("test")

	val = GetValueForKey("test")
	if val != nil {
		t.Fatalf(`Failed: expected nil value for deleted key`)
	}

	val = GetValueForKey("")
	if val != nil {
		t.Fatalf(`Failed: expected nil value for empty key`)
	}
}

// verifies setting key values
func TestSetValueForKey(t *testing.T) {
	SetValueForKey("", "123")

	val := GetValueForKey("")
	if val != nil {
		t.Fatalf(`Failed: expected nil value for invalid key`)
	}

	SetValueForKey("test", "123")

	val = GetValueForKey("test")
	if val != "123" {
		t.Fatalf(`Failed: expected '123' value for key 'test'`)
	}

	DeleteKey("test")
}

// tests delete key functionality
func TestDeleteKey(t *testing.T) {
	DeleteKey("")

	SetValueForKey("test", "123")
	DeleteKey("test")

	if len(GetKeyValuePairs()) > 0 {
		t.Fatalf(`Failed: expected length of keyvaluestore after delete to be zero`)
	}
}
