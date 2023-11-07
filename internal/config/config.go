package config

import (
	"os"
	"strconv"
	"time"
)

const (
	DefaultAddress        = ":8080"
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

func GetConfig() Config {
	address := os.Getenv("ADDRESS")
	if address == "" {
		address = DefaultAddress
	}
	reportInterval := time.Duration(getEnvAsInt("REPORT_INTERVAL", DefaultReportInterval)) * time.Second
	pollInterval := time.Duration(getEnvAsInt("POLL_INTERVAL", DefaultPollInterval)) * time.Second

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
