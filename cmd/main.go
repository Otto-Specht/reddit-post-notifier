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

	userNames := GetUserNames()
	logger.Debug("Users: " + strings.Join(userNames, ", "))

	clientId := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientId == "" || clientSecret == "" {
		logger.FatalAndExit("Missing CLIENT_ID and/or CLIENT_SECRET")
	}
	api.GenerateAccessToken(clientId, clientSecret)

}

func GetUserNames() []string {
	userNamesRaw := os.Getenv("USER_NAMES")
	if userNamesRaw == "" {
		logger.FatalAndExit("No user names provided. Exiting...")
	}

	return strings.Split(userNamesRaw, ",")
}
