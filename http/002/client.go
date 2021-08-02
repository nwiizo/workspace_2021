package main

import (
	"fmt"
	"net/http"
)

func main() {
	response, _ := http.Get("http://localhost:8000")
	fmt.Println(response.Header.Get("content-type"))
}
