package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

// Config holds all configuration values
type Config struct {
	Port              string
	GinMode           string
	CORSOrigins       []string
	FileBrowserRoot   string
	EnableFileBrowser bool
	LogLevel          string
}

// LoadConfig loads configuration from environment variables
func LoadConfig() *Config {
	// Try to load .env file from the project root (one level up from backend)
	envPath := filepath.Join("..", ".env")
	if err := godotenv.Load(envPath); err != nil {
		log.Printf("Warning: Could not load .env file from %s: %v", envPath, err)
		log.Println("Continuing with system environment variables...")
	}

	config := &Config{
		Port:              getEnvWithDefault("PORT", "8080"),
		GinMode:           getEnvWithDefault("GIN_MODE", "debug"),
		CORSOrigins:       parseCORSOrigins(getEnvWithDefault("CORS_ORIGINS", "http://localhost:5173,http://localhost:3000")),
		FileBrowserRoot:   getEnvWithDefault("FILE_BROWSER_ROOT", "/"),
		EnableFileBrowser: getEnvBool("ENABLE_FILE_BROWSER", true),
		LogLevel:          getEnvWithDefault("LOG_LEVEL", "info"),
	}

	return config
}

// getEnvWithDefault gets an environment variable with a default value
func getEnvWithDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// getEnvBool gets a boolean environment variable with a default value
func getEnvBool(key string, defaultValue bool) bool {
	if value := os.Getenv(key); value != "" {
		if parsed, err := strconv.ParseBool(value); err == nil {
			return parsed
		}
	}
	return defaultValue
}

// parseCORSOrigins parses the CORS_ORIGINS environment variable
func parseCORSOrigins(origins string) []string {
	if origins == "" {
		return []string{}
	}

	// Split by comma and trim whitespace
	parts := strings.Split(origins, ",")
	result := make([]string, len(parts))
	for i, part := range parts {
		result[i] = strings.TrimSpace(part)
	}

	return result
}
