// File: config/config.go
// Description: This package is responsible for loading environment variables from a .env file.
// It uses the godotenv library to load the variables into the application.
package config

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a .env file.
// It uses the godotenv library to read the file and set the variables in the environment.
// If the file cannot be loaded, it logs a fatal error and exits the application.
func LoadEnv() {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
}
