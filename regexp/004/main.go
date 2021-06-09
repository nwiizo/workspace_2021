package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	// url の指定
	url := "https://3-shake.com/"
	// 正規表現の作成
	re, err := regexp.Compile("http(.*)://(.*)")
	if err != nil {
		return
	}
	// net/http でのリクエストの発射
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	// []byte でリクエストの中身を取得
	byteArray, _ := ioutil.ReadAll(resp.Body)
	// 正規表現にあったものを全てlinks に入れる
	links := re.FindAllString(string(byteArray), -1)
	for i := 0; i < len(links); i++ {
		fmt.Println(links[i])
	}
}
