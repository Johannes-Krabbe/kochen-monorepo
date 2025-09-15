package internal

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	DatabaseURL  string
	PostgresUser string
	PostgresPass string
	PostgresDB   string
	PostgresHost string
	PostgresPort string
	Port         int
}

func LoadConfig() (*Config, error) {
	config := &Config{}

	// Load individual postgres variables
	config.PostgresUser = os.Getenv("POSTGRES_USER")
	if config.PostgresUser == "" {
		return nil, fmt.Errorf("POSTGRES_USER environment variable is required")
	}

	config.PostgresPass = os.Getenv("POSTGRES_PASSWORD")
	if config.PostgresPass == "" {
		return nil, fmt.Errorf("POSTGRES_PASSWORD environment variable is required")
	}

	config.PostgresDB = os.Getenv("POSTGRES_DB")
	if config.PostgresDB == "" {
		return nil, fmt.Errorf("POSTGRES_DB environment variable is required")
	}

	// Default postgres host and port
	config.PostgresHost = os.Getenv("POSTGRES_HOST")
	if config.PostgresHost == "" {
		config.PostgresHost = "kochen-postgres"
	}

	config.PostgresPort = os.Getenv("POSTGRES_PORT")
	if config.PostgresPort == "" {
		config.PostgresPort = "5432"
	}

	// Build database URL from components
	config.DatabaseURL = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		config.PostgresUser, config.PostgresPass, config.PostgresHost, config.PostgresPort, config.PostgresDB)

	// Load server port
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