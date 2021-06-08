package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	// url の指定
	url := "https://www.cman.jp/network/support/go_access.cgi"
	// 正規表現の作成
	var re = regexp.MustCompile(`((25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|1[0-9][0-9]|[1-9]?[0-9])`)

	// net/http でのリクエストの発射
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	// []byte でリクエストの中身を取得
	byteArray, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(byteArray))
	// 正規表現にあったものを全てlinks に入れる
	links := re.FindAllString(string(byteArray), -1)
	fmt.Println(links)
	for i := 0; i < len(links); i++ {
		fmt.Println(links[i])
	}
}
