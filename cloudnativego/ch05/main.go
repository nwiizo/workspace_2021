package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func helloMuxHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello gorilla/mux!\n"))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", helloMuxHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}
