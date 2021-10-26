package main

import "fmt"

func main() {
	and := true && false
	fmt.Println(and) // false

	or := true || false
	fmt.Println(or) // true

	not := !true
	fmt.Println(not) // false
}
