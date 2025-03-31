package util

import (
	"fmt"
	"os"

	"github.com/Otto-Specht/reddit-post-notifier/pkg/logger"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Allow skipping of .env file eg. for Docker
	if os.Getenv("SKIP_ENV") != "true" {
		logger.Debug("Loading env file...")
		err := godotenv.Load(".env")
		if err != nil {
			logger.Warn(fmt.Sprintf("Failed to load env file: %s", err.Error()))
		}
	}
}
