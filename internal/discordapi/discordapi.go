package discordapi

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Otto-Specht/reddit-post-notifier/pkg/logger"
)

var api DiscordApi = DiscordApi{
	httpClient:        http.Client{},
	AccessToken:       "",
	AccessTokenExpire: 0,
}

func SendServerMessage() {
	discordBotToken := os.Getenv("DISCORD_BOT_TOKEN")
	discordChannelId := os.Getenv("DISCORD_CHANNEL_ID")
	if discordBotToken == "" || discordChannelId == "" {
		logger.FatalAndExit("Missing DISCORD_BOT_TOKEN and/or DISCORD_CHANNEL_ID")
	}

	req, err := http.NewRequest("post", fmt.Sprintf("https://discord.com/api/v10/channels/%s/messages", discordChannelId), nil)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create discord message request: %s", err.Error()))
		return
	}

	resp, err := api.httpClient.Do(req)
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to send discord message: %s", err.Error()))
		return
	}

	if resp.StatusCode != http.StatusOK {
		logger.Error(fmt.Sprintf("Failed to send discord message: %s", resp.Status))
	}

	logger.Info("Discord message sent successfully")
}
