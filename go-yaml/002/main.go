package main

import (
	"bytes"
	"fmt"

	yaml "github.com/goccy/go-yaml"
)

var v struct {
	A struct {
		B int
		C string
		D int
	}
}

func main() {

	buf := bytes.NewBufferString("a: *a\n")
	dec := yaml.NewDecoder(buf, yaml.ReferenceDirs("testdata"))

	if err := dec.Decode(&v); err != nil {
		//...
	}
	fmt.Printf("%+v\n", v)
}
