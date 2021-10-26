package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	// 処理
	time.Sleep(time.Second * 3)
	end := time.Now()
	fmt.Printf("%f秒\n", (end.Sub(start)).Seconds())
}
