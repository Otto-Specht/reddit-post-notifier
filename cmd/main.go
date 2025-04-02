package main

import (
	"os"
	"strings"

	"github.com/Otto-Specht/reddit-post-notifier/internal/api"
	"github.com/Otto-Specht/reddit-post-notifier/internal/controller"
	"github.com/Otto-Specht/reddit-post-notifier/pkg/logger"
	"github.com/Otto-Specht/reddit-post-notifier/pkg/util"
)

func main() {
	logger.Info("Starting...")

	util.LoadEnv()
	util.VerifyEnv()

	userNames := api.CheckIfUsersExistOrRemove(GetUserNames())

	if len(userNames) == 0 {
		logger.FatalAndExit("No usernames to check...")
	}

	controller.Start(userNames)
}

func GetUserNames() []string {
	userNamesRaw := os.Getenv("USER_NAMES")
	if userNamesRaw == "" {
		logger.FatalAndExit("No user names provided. Exiting...")
	}

	return strings.Split(userNamesRaw, ",")
}
