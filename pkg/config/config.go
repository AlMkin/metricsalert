package config

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Address        string
	ReportInterval time.Duration
	PollInterval   time.Duration
}

func GetEnvOrDefault(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func GetEnvOrFlagInt(envKey string, flagVal *int, defaultVal int) int {
	if value, exists := os.LookupEnv(envKey); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
		fmt.Printf("Warning: could not parse %s, using default %d\n", envKey, defaultVal)
	}
	return *flagVal
}

func LoadConfig() (*Config, error) {
	serverAddressFlag := flag.String("a", "localhost:8080", "address of the metrics server")
	reportIntervalFlag := flag.Int("r", 10, "report interval in seconds")
	pollIntervalFlag := flag.Int("p", 2, "poll interval in seconds")

	flag.Parse()

	if flag.NArg() > 0 {
		return nil, fmt.Errorf("unknown flags provided")
	}

	cfg := &Config{
		Address:        GetEnvOrDefault("ADDRESS", *serverAddressFlag),
		ReportInterval: time.Second * time.Duration(GetEnvOrFlagInt("REPORT_INTERVAL", reportIntervalFlag, 10)),
		PollInterval:   time.Second * time.Duration(GetEnvOrFlagInt("POLL_INTERVAL", pollIntervalFlag, 2)),
	}

	return cfg, nil
}
