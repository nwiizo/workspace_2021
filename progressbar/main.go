package main

import (
	"time"

	progressbar "github.com/schollz/progressbar/v3"
)

func main() {
	bar := progressbar.Default(100)
	for i := 0; i < 20; i++ {
		bar.Add(5)
		time.Sleep(20 * time.Millisecond)
	}
}
