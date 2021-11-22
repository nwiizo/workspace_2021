package main

import (
	"bufio"
	"fmt"
	"log"
	"net/url"
	"os"
)

func main() {
	data, _ := os.Open("urls.txt")

	defer data.Close()
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		baseURL, err := url.Parse(scanner.Text())

		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("URL: %s\n", baseURL.String())
		fmt.Printf("Host: %s\n", baseURL.Host)
		// HOST 名が
		fmt.Printf("Hostname(): %s\n", baseURL.Hostname())
		fmt.Printf("Path: %s\n", baseURL.Path)
		fmt.Printf("RawPath: %s\n", baseURL.RawPath)
		fmt.Printf("RawQuery: %s\n", baseURL.RawQuery)
		fmt.Printf("Fragment: %s\n", baseURL.Fragment)
		for key, values := range baseURL.Query() {
			fmt.Printf("Query Key: %s\n", key)
			for i, v := range values {
				fmt.Printf("Query Value[%d]: %s\n", i, v)
			}
		}
	}
}
