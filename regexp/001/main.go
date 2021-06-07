package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	for _, n := range []int{1, 10, 100, 1000, 10000} {
		// Generate target string
		var s1 string
		for i := 0; i < n; i++ {
			s1 += "a?"
		}
		var s2 string
		for i := 0; i < n; i++ {
			s2 += "a"
		}
		re := s1 + s2 // a?a?a?aaa (ex. n==3)
		fmt.Println(re)
		fmt.Println("n:", n)
		start := time.Now()
		// Compile to NFA
		regOjb := regexp.MustCompile(re)
		// Do matching
		regOjb.MatchString(s2)
		end := time.Now()
		fmt.Printf("%f [sec]\n", (end.Sub(start)).Seconds())
	}
}
