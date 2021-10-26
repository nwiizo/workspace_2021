package main

import (
	"fmt"
	"os"

	"github.com/bojand/ghz/printer"
	"github.com/bojand/ghz/runner"
)

func main() {
	report, err := runner.Run(
		"radiko.ActorService.SearchActors",
		"server-actor-and-article-grpc-gm32druggq-an.a.run.app:443",
		runner.WithDataFromJSON("{\"actor_name\":\"佐藤\"}"),
		//runner.WithInsecure(true),
	)

	if err != nil {
		fmt.Println("hiiii")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	printer := printer.ReportPrinter{
		Out:    os.Stdout,
		Report: report,
	}

	//printer.Print("pretty")
	printer.Print("summary")
}
