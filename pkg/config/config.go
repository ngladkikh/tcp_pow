package config

import (
	"os"
	"strconv"
)

type Config struct {
	ServerAddr    string
	ServerPort    string
	PowComplexity int
	TimeoutMillis int
}

func NewConfig() (*Config, error) {
	complexity, err := strconv.Atoi(getEnv("POW_COMPLEXITY", "3"))
	if err != nil {
		return nil, err
	}

	timeout, err := strconv.Atoi(getEnv("SERVER_TIMEOUT_MILLIS", "5000"))
	if err != nil {
		return nil, err
	}

	return &Config{
		ServerAddr:    getEnv("SERVER_ADDR", "0.0.0.0"),
		ServerPort:    getEnv("SERVER_PORT", "9999"),
		PowComplexity: complexity,
		TimeoutMillis: timeout,
	}, nil
}

func getEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
