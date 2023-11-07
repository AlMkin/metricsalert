package config

import (
	"flag"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultAddress        = "http://localhost:8080"
	DefaultReportInterval = 10
	DefaultPollInterval   = 2
)

type Config struct {
	Address        string
	ReportInterval time.Duration
	PollInterval   time.Duration
}

func getEnvAsInt(name string, defaultValue int) int {
	valueStr := os.Getenv(name)
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultValue
}

func ensureHTTPPrefix(serverAddress string) string {
	if !strings.HasPrefix(serverAddress, "http://") && !strings.HasPrefix(serverAddress, "https://") {
		return "http://" + serverAddress
	}
	return serverAddress
}

func GetConfig() Config {
	serverAddressPtr := flag.String("a", "", "address of the metrics server")
	reportIntervalPtr := flag.Int("r", DefaultReportInterval, "report interval in seconds")
	pollIntervalPtr := flag.Int("p", DefaultPollInterval, "poll interval in seconds")

	flag.Parse()

	address := os.Getenv("ADDRESS")
	if address == "" {
		address = DefaultAddress
	}
	reportInterval := time.Duration(getEnvAsInt("REPORT_INTERVAL", DefaultReportInterval)) * time.Second
	pollInterval := time.Duration(getEnvAsInt("POLL_INTERVAL", DefaultPollInterval)) * time.Second

	if *serverAddressPtr != "" {
		address = ensureHTTPPrefix(*serverAddressPtr)
	}
	if *reportIntervalPtr != DefaultReportInterval {
		reportInterval = time.Duration(*reportIntervalPtr) * time.Second
	}
	if *pollIntervalPtr != DefaultPollInterval {
		pollInterval = time.Duration(*pollIntervalPtr) * time.Second
	}

	return Config{
		Address:        address,
		ReportInterval: reportInterval,
		PollInterval:   pollInterval,
	}
}

func GetServerConfig(cmdLineAddr string) Config {
	config := GetConfig()

	if cmdLineAddr != "" {
		config.Address = cmdLineAddr
	}

	return config
}
