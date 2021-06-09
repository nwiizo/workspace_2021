package main

import (
	"fmt"
	"regexp"
)

func main() {
	re := regexp.MustCompile("^iPhone(\\d+),(\\d+)$")
	fmt.Println(re.FindStringSubmatch("iPhone10,6"))
}
