package main

import (
	"errors"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

var store = make(map[string]string)

var ErrorNoSuchKey = errors.New("no such key")

func Get(key string) (string, error) {
	value, ok := store[key]

	if !ok {
		return "", ErrorNoSuchKey
	}

	return value, nil
}

func Put(key string, value string) error {
	store[key] = value

	return nil
}

func Delete(key string) error {
	delete(store, key)

	return nil
}

// keyValuePutHandler expects to be called with a PUT request for
// the "/v1/key/{key}" resource.
func keyValuePutHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Retrieve "key" from the request
	key := vars["key"]

	value, err := io.ReadAll(r.Body) // The request body has our value
	defer r.Body.Close()

	if err != nil { // If we have an error, report it
		http.Error(w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}

	err = Put(key, string(value)) // Store the value as a string
	if err != nil {               // If we have an error, report it
		http.Error(w,
			err.Error(),
			http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated) // All good! Return StatusCreated
}

func main() {
	r := mux.NewRouter()

	// Register keyValuePutHandler as the handler function for PUT
	// requests matching "/v1/{key}"
	r.HandleFunc("/v1/{key}", keyValuePutHandler).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}
