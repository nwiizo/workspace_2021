package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
)

//go:embed sample.json
var sampleBytes []byte

type sample struct {
	Key1 string `json:"key1"`
	Key2 string `json:"key2"`
}

func main() {
	var s sample
	if err := json.Unmarshal(sampleBytes, &s); err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", s)
}
