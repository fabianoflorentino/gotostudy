package config

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load("config/.env"); err != nil {
		log.Fatalf("Error loading .env file")
	}
}
