package main

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/v12/lib"
)

func main() {
	rate := vegeta.Rate{Freq: 100, Per: time.Second}
	duration := 500 * time.Second
	urls := [...]string{"https://3-shake.com/", "https://relance.jp/"}
	for _, url := range urls {
		targeter := vegeta.NewStaticTargeter(vegeta.Target{
			Method: "GET",
			URL:    url,
		})
		attacker := vegeta.NewAttacker()

		var metrics vegeta.Metrics
		fmt.Println("start:\n", url)
		for res := range attacker.Attack(targeter, rate, duration, "Big Bang!") {
			metrics.Add(res)
		}
		metrics.Close()

		fmt.Printf("Latencies:\n")
		fmt.Printf("- total %d [ms]\n", metrics.Latencies.Total.Milliseconds())
		fmt.Printf("- mean  %d [ms]\n", metrics.Latencies.Mean.Milliseconds())
		fmt.Printf("- 50th  %d [ms]\n", metrics.Latencies.P50.Milliseconds())
		fmt.Printf("- 90th  %d [ms]\n", metrics.Latencies.P90.Milliseconds())
		fmt.Printf("- 95th  %d [ms]\n", metrics.Latencies.P95.Milliseconds())
		fmt.Printf("- 99th  %d [ms]\n", metrics.Latencies.P99.Milliseconds())
		fmt.Printf("- max   %d [ms]\n", metrics.Latencies.Max.Milliseconds())
		fmt.Printf("- min   %d [ms]\n", metrics.Latencies.Min.Milliseconds())
		fmt.Printf("Duration:\n")
		fmt.Printf("- duration %d [ms]\n", metrics.Duration.Milliseconds())
		fmt.Printf("Wait:\n")
		fmt.Printf("- wait %d [ms]\n", metrics.Wait.Milliseconds())
	}
}
