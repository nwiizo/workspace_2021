package main

import (
	"fmt"

	"github.com/goccy/go-yaml"
)

var v struct {
	A int
	B string
}

func main() {

	v.A = 1
	v.B = "hello"
	bytes, err := yaml.Marshal(v)
	if err != nil {
		//...
	}
	fmt.Println(string(bytes))
}
