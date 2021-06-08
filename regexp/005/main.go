package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
)

func privateIP(ip string) (bool, error) {
	var err error
	private := false
	IP := net.ParseIP(ip)
	if IP == nil {
		err = errors.New("Invalid IP")
	} else {
		_, private24BitBlock, _ := net.ParseCIDR("10.0.0.0/8")
		_, private20BitBlock, _ := net.ParseCIDR("172.16.0.0/12")
		_, private16BitBlock, _ := net.ParseCIDR("192.168.0.0/16")
		private = private24BitBlock.Contains(IP) || private20BitBlock.Contains(IP) || private16BitBlock.Contains(IP)
	}
	return private, err
}

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
		fmt.Println(privateIP(links[i]))
	}
}
