package main

import (
	"testing"
)

func TestinsertMap(t *testing.T) {
	// test the InsertMap function
	store := make(map[string]pair)

	// clean insert on a blank map should succeed
	code, err := insertMap(pair{key: "Hello", value: "World"}, store)
	if err != nil || code != 1 {
		t.Errorf("Inserting {%s, %s} on an empty map failed with code [%d].", "hello", "world", code)
	}

	// inserting the same value should fail
	code_1, err_1 := insertMap(pair{key: "Hello", value: "Moon"}, store)
	if code_1 == 1 || err_1 == nil {
		t.Errorf("Duplicate Values Entered without Breaking!")
	}
}
