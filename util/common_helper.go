package util

import "os"

// GetEnvWithDefault retrieves an environment variable with a default value.
func GetEnvWithDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
