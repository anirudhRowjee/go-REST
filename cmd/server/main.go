// simple HTTP REST Key-Value Store
package main

import (
	"errors"
	"io"
	"log"
	"net/http"
)

// basic KV datatype
type pair struct {
	key   string `json:"key"`
	value string `json:"value"`
}

// utils - query the map, accessed by pointer

func insertMap(toInsert pair, store map[string]pair) (status int, err error) {
	// insert a pair into the map
	nullpair := pair{"", ""}
	key := toInsert.key

	// check if there's already a value for the key
	if store[key] != nullpair {
		// log.Println("Request Tried to insert key already present")
		return 0, errors.New("Key Already Present!")
	}

	// now we know it isn't there, we insert
	store[key] = toInsert
	return 1, nil
}

func queryMap(toQuery string, store map[string]pair) (response pair, err error) {
	// query the map for a pair
	nullpair := pair{"", ""}
	key := toQuery
	result := store[key]
	if result == nullpair {
		// We didn't find a value for this key
		return nullpair, errors.New("key not present")
	}
	return result, nil
}

/*
func assembleCompleteMap(store map[string]pair) map[string]string {
	// assemble the complete map a JSON-Serializable Map
}
*/

func parseRequestBody(req *http.Request) (key string, value string) {
	// parse the HTTP request body for the data

	// dummy return for now
	return "", ""
}

// function handlers

func main() {

	// map of URL Patterns and Handlers
	urlMap := make(map[string]func(http.ResponseWriter, *http.Request))

	// declare a basic data store - Key, {}
	store := make(map[string]pair)

	// test insertion
	_, err := insertMap(pair{key: "Hello", value: "World"}, store)
	if err != nil {
		log.Println("Element not present!")
	} else {
		log.Println("Success!")
	}

	// test duplicate insertion

	log.Println(queryMap("Hello", store))

	// failing cache access
	res1, err1 := queryMap("xyz", store)
	if err1 != nil {
		log.Println("Element not present!")
	} else {
		log.Println(res1.key, " : ", res1.value)
	}

	// hello world for REST!
	handler := func(w http.ResponseWriter, req *http.Request) {
		log.Println("Recieved a request!")
		io.WriteString(w, "Hello, World!\n")
	}

	urlMap["/"] = handler

	// register the handlers
	for url, handler := range urlMap {
		http.HandleFunc(url, handler)
	}

	// logging setup
	log.Println("Listening for Req at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
