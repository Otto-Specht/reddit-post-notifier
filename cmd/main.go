package main

import (
	"os"
	"strings"

	"github.com/Otto-Specht/reddit-post-notifier/internal/api"
	"github.com/Otto-Specht/reddit-post-notifier/pkg/logger"
	"github.com/Otto-Specht/reddit-post-notifier/pkg/util"
)

func main() {
	logger.Info("Initilizing...")

	util.LoadEnv()
	util.VerifyEnv()

	userNames := GetUserNames()
	logger.Debug("Users: " + strings.Join(userNames, ", "))

	userNames = api.CheckIfUsersExistOrRemove(userNames)

	for _, value := range userNames {
		api.GetLatestPostId(value)
	}
}

func GetUserNames() []string {
	userNamesRaw := os.Getenv("USER_NAMES")
	if userNamesRaw == "" {
		logger.FatalAndExit("No user names provided. Exiting...")
	}

	return strings.Split(userNamesRaw, ",")
}
