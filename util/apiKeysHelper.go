package util

import (
	"os"
	"log"
	"strings"
)

// Helps in getting API Keys
func GetAPIKeys() []string {
	apiKeyStr, exists := os.LookupEnv("API_KEYS")
	if !exists {
		log.Fatal("API_KEYS Environment Variable not found")
	}

	// Create a array o api keys
	keys := strings.Split(apiKeyStr, ",")

	// Keeping the limit of alteast 3 api keys
	if len(keys) < 3 {
		log.Fatal("Sufficient Api Keys not available")
	}
	return keys
}