package config

import (
	"os"
)

// Config struct holds the configuration settings for the application.
type Config struct {
	NodeAddress  string
	APIAddress   string
	InitialPeer  string
}

// LoadConfig loads configuration settings from environment variables.
func LoadConfig() *Config {
	return &Config{
		NodeAddress: getEnv("NODE_ADDRESS", "localhost:3001"),
		APIAddress:  getEnv("API_ADDRESS", "localhost:8080"),
		InitialPeer: getEnv("INITIAL_PEER", ""),
	}
}

// getEnv retrieves an environment variable or returns a default value if not set.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
