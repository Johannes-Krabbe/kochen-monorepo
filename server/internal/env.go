package internal

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DatabaseURL string
	Port        int
}

func LoadConfig() (*Config, error) {
	config := &Config{}

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable is required")
	}
	config.DatabaseURL = databaseURL

	portStr := os.Getenv("PORT")
	if portStr == "" {
		config.Port = 8080
	} else {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, fmt.Errorf("PORT must be a valid integer: %w", err)
		}
		if port < 1 || port > 65535 {
			return nil, fmt.Errorf("PORT must be between 1 and 65535")
		}
		config.Port = port
	}

	return config, nil
}