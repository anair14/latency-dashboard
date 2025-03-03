package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// LoadConfig loads environment variables from a .env file
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: No .env file found, using default values.")
	}
}

// GetEnv retrieves an environment variable or returns a default value
func GetEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

