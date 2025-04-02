package main

import (
	"os"
	"strings"

	"github.com/Otto-Specht/reddit-post-notifier/internal/controller"
	"github.com/Otto-Specht/reddit-post-notifier/pkg/logger"
	"github.com/Otto-Specht/reddit-post-notifier/pkg/util"
)

func main() {
	logger.Info("Starting...")

	util.LoadEnv()
	util.VerifyEnv()

	userNames := GetUserNames()
	logger.Debug("Users: " + strings.Join(userNames, ", "))

	//userNames = api.CheckIfUsersExistOrRemove(userNames)

	controller.Start([]string{})
}

func GetUserNames() []string {
	userNamesRaw := os.Getenv("USER_NAMES")
	if userNamesRaw == "" {
		logger.FatalAndExit("No user names provided. Exiting...")
	}

	return strings.Split(userNamesRaw, ",")
}
