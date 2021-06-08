package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	url := "https://3-shake.com/"
	re, err := regexp.Compile("http(.*):(.*)")
	if err != nil {
		return
	}
	resp, _ := http.Get(url)
	defer resp.Body.Close()

	byteArray, _ := ioutil.ReadAll(resp.Body)

	//fmt.Println(string(byteArray)) // htmlをstringで取得
	fmt.Println(re.FindAllString(string(byteArray), -1))
	words := re.FindAllString(string(byteArray), -1)
	for i := 0; i < len(words); i++ {
		fmt.Println(words[i])
	}
}
