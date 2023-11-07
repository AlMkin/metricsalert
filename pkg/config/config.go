package config

import (
	"fmt"
	"os"
	"strconv"
)

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
