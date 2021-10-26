package main

import (
	"fmt"

	yaml "github.com/goccy/go-yaml"
)

func main() {

	type Person struct {
		*Person `yaml:",omitempty,inline,alias"` // embed Person type for default value
		Name    string                           `yaml:",omitempty"`
		Age     int                              `yaml:",omitempty"`
	}
	defaultPerson := &Person{
		Name: "John Smith",
		Age:  20,
	}
	people := []*Person{
		{
			Person: defaultPerson, // assign default value
			Name:   "Ken",         // override Name property
			Age:    10,            // override Age property
		},
		{
			Person: defaultPerson, // assign default value only
		},
		{
			Name: "Shuya",
			Age:  26,
		},
	}
	var doc struct {
		Default *Person   `yaml:"default,anchor"`
		People  []*Person `yaml:"people"`
	}
	doc.Default = defaultPerson
	doc.People = people
	bytes, err := yaml.Marshal(doc)
	if err != nil {
		//...
	}
	fmt.Println(string(bytes))
	/*
	   default: &default
	     name: John Smith
	       age: 20
	       people:
	       - <<: *default
	         name: Ken
	           age: 10
	           - <<: *default
	*/

}
