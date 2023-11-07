package main

import (
	"flag"
	"fmt"
	"github.com/AlMkin/metricsalert/internal/agent"
	"github.com/AlMkin/metricsalert/internal/metrics"
	"github.com/AlMkin/metricsalert/internal/sender"
	"github.com/AlMkin/metricsalert/pkg/runtimeinfo"
	"os"
	"strconv"
	"time"
)

func getEnvOrFlag(envKey string, flagVal *int, defaultVal int) int {
	if value, exists := os.LookupEnv(envKey); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		fmt.Printf("Warning: could not parse %s, using default %d\n", envKey, defaultVal)
	}
	return *flagVal
}

func main() {
	serverAddressFlag := flag.String("a", "localhost:8080", "address of the metrics server")
	reportIntervalFlag := flag.Int("r", 10, "report interval in seconds")
	pollIntervalFlag := flag.Int("p", 2, "poll interval in seconds")

	flag.Parse()

	if flag.NArg() > 0 {
		fmt.Println("Error: unknown flags provided")
		flag.Usage()
		os.Exit(1)
	}

	serverAddress := os.Getenv("ADDRESS")
	if serverAddress == "" {
		serverAddress = *serverAddressFlag
	}
	reportIntervalSeconds := getEnvOrFlag("REPORT_INTERVAL", reportIntervalFlag, 10)
	pollIntervalSeconds := getEnvOrFlag("POLL_INTERVAL", pollIntervalFlag, 2)

	newSender := sender.NewSender(serverAddress)
	metricsGetter := &runtimeinfo.Getter{}

	collector := metrics.NewCollector(metricsGetter)
	a := agent.NewAgent(newSender, collector, time.Duration(pollIntervalSeconds)*time.Second, time.Duration(reportIntervalSeconds)*time.Second)
	a.Run()
}
