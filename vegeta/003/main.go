package main

import (
	"fmt"
	"os"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	rate := vegeta.Rate{Freq: 1000, Per: time.Second}
	duration := 5 * time.Second
	//targeter := vegeta.NewStaticTargeter()

	file, err := os.Open("./single_http.list")
	if err != nil {
		fmt.Printf("url file: %s", err)
		os.Exit(1)
	}
	defer file.Close()
	//scanner := bufio.NewScanner(file)
	targeter := vegeta.NewHTTPTargeter(file, nil, nil)
	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics
	for res := range attacker.Attack(targeter, rate, duration, "Load Test: ./single_http.list") {
		metrics.Add(res)
		if metrics.Latencies.Total.Milliseconds() != 0 {
			fmt.Printf("start: %s \n", res.Attack)
			fmt.Printf("- total %d [ms]\n", metrics.Latencies.Total.Milliseconds())
			fmt.Printf("- 99th percentile: %s\n", metrics.Latencies.P99)
		}
		metrics.Close()
	}

}
