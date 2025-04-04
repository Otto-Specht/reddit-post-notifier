package discordapi

import (
	"bytes"
	"encoding/json"
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

func SendServerMessage(message string) {
	discordBotToken := os.Getenv("DISCORD_BOT_TOKEN")
	discordChannelId := os.Getenv("DISCORD_CHANNEL_ID")
	if discordBotToken == "" || discordChannelId == "" {
		logger.FatalAndExit("Missing DISCORD_BOT_TOKEN and/or DISCORD_CHANNEL_ID")
	}

	url := fmt.Sprintf("https://discord.com/api/v10/channels/%s/messages", discordChannelId)
	jsonData, err := json.Marshal(map[string]string{"content": message})
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create discord message request data: %s", err.Error()))
		return
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Error(fmt.Sprintf("Failed to create discord message request: %s", err.Error()))
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bot "+discordBotToken)

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
