package main

import (
	"os"
	"strings"

	"github.com/Otto-Specht/reddit-post-notifier/pkg/logger"
	"github.com/Otto-Specht/reddit-post-notifier/pkg/util"
)

func main() {
	logger.Info("Initilizing...")

	util.LoadEnv()

	userNames := GetUserNames()

	logger.Debug("User names: " + strings.Join(userNames, ", "))
}

func GetUserNames() []string {
	userNamesRaw := os.Getenv("USER_NAMES")
	if userNamesRaw == "" {
		logger.FatalAndExit("No user names provided. Exiting...")
	}

	return strings.Split(userNamesRaw, ",")
}
