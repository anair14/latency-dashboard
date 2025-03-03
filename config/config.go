package config

import (
    "os"
)

// Database Configuration
var DBPath = "./database/data.db"

// Authentication Configuration
var SessionSecret = getEnv("SESSION_SECRET", "supersecretkey")

// Default Latency Threshold
var DefaultLatencyThreshold = 500

// Helper function to get environment variables with a default fallback
func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}

