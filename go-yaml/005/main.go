package main

import (
	"fmt"

	"github.com/goccy/go-yaml"
)

func main() {
	yml := `
  a: 1
  b: "hello"
  `
	var v struct {
		A int
		B string
	}
	if err := yaml.Unmarshal([]byte(yml), &v); err != nil {
		panic(err)
	}
	if v.A != 2 {
		// output error with YAML source
		path, err := yaml.PathString("$.a")
		if err != nil {
			panic(err)
		}
		source, err := path.AnnotateSource([]byte(yml), true)
		if err != nil {
			panic(err)
		}
		fmt.Printf("a value expected 2 but actual %d:\n%s\n", v.A, string(source))
	}
}
