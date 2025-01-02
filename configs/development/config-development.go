package development

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ErrEnvNotFound = fmt.Errorf("environment file not found")
	ErrEnvLoading  = fmt.Errorf("failed to load environment")
)

// Init loads environment variables from .env file
func Init() error {
	// Try loading from multiple possible locations
	err1 := godotenv.Load()                // Try root directory
	err2 := godotenv.Load("../../.env")    // Try relative to service
	err3 := godotenv.Load("../../../.env") // Try one level up

	if err1 != nil && err2 != nil && err3 != nil {
		return fmt.Errorf("%w: .env file not found in any location", ErrEnvNotFound)
	}
	return nil
}

// GetEnvStr retrieves string value from environment variable
func GetEnvStr(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists && value != "" {
		return value
	}
	return fallback
}

// GetEnvInt retrieves integer value from environment variable
func GetEnvInt(key string, fallback int) int {
	value, exists := os.LookupEnv(key)
	if !exists || value == "" {
		return fallback
	}

	intVal, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}

	return intVal
}
