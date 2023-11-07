package main

import (
	"flag"
	"fmt"
	"github.com/AlMkin/metricsalert/internal/agent"
	"github.com/AlMkin/metricsalert/internal/metrics"
	"github.com/AlMkin/metricsalert/internal/sender"
	"github.com/AlMkin/metricsalert/pkg/runtimeinfo"
	"os"
	"time"
)

func main() {
	serverAddress := flag.String("a", "localhost:8080", "address of the metrics server")
	reportIntervalSeconds := flag.Int("r", 10, "report interval in seconds")
	pollIntervalSeconds := flag.Int("p", 2, "poll interval in seconds")

	flag.Parse()

	if flag.NArg() > 0 {
		fmt.Println("Error: unknown flags provided")
		flag.Usage()
		os.Exit(1)
	}

	newSender := sender.NewSender(*serverAddress)
	metricsGetter := &runtimeinfo.Getter{}

	collector := metrics.NewCollector(metricsGetter)
	a := agent.NewAgent(newSender, collector, time.Duration(*pollIntervalSeconds)*time.Second, time.Duration(*reportIntervalSeconds)*time.Second)
	a.Run()
}
