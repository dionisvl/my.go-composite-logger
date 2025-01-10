package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found, using default environment variables")
	} else {
		log.Println(".env file loaded successfully")
	}

	// Load .env.override file (if exists)
	if err := godotenv.Overload(".env.override"); err != nil {
		log.Println("Info: .env.override file not found, no overrides applied")
	} else {
		log.Println(".env.override file loaded successfully, overrides applied")
	}
}

func GetSentryDSN() string {
	return os.Getenv("SENTRY_DSN")
}

func IsSentryEnabled() bool {
	return os.Getenv("USE_SENTRY") == "true"
}
