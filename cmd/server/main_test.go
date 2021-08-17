package main

import "testing"

func TestInsertMap(t *testing.T) {
	// test the InsertMap function
	store := make(map[string]pair)

	// clean insert on a blank map should succeed
	code, err := insertMap(pair{key: "Hello", value: "World"}, store)
	if err != nil {
		t.Errorf("Inserting {%s, %s} on an empty map failed with code [%d].", "hello", "world", code)
	}

	// inserting the same value should fail
	code_1, err_1 := insertMap(pair{key: "Hello", value: "Moon"}, store)
	if code_1 == 1 || err_1 == nil {
		t.Errorf("Duplicate Values Entered without Breaking!")
	}
}

func TestQueryMap(t *testing.T) {

	store := make(map[string]pair)
	insertMap(pair{key: "Hello", value: "World"}, store)

	// checking for a pre-existing value
	_, err := queryMap("Hello", store)
	if err != nil {
		t.Errorf("Querying Existing Value Failed")
	}

	// check for a value that doesn't exist
	_, err1 := queryMap("XY", store)
	if err1 == nil {
		t.Errorf("Value that doesn't exist queried succesfully")
	}

}
