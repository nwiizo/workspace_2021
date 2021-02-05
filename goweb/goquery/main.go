package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://www.conoha.jp/"

	// Getリクエスト
	res, _ := http.Get(url)
	defer res.Body.Close()

	// 読み取り
	buf, _ := ioutil.ReadAll(res.Body)

	// HTMLパース
	reader := bytes.NewReader(buf)
	doc, _ := goquery.NewDocumentFromReader(reader)

	// titleを抜き出し
	rslt := doc.Find("title").Text()
	fmt.Println(rslt)
}
