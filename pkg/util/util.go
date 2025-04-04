package util

import (
	"fmt"
	"os"
	"time"

	"github.com/Otto-Specht/reddit-post-notifier/internal/redditapi"
	"github.com/Otto-Specht/reddit-post-notifier/pkg/logger"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Allow skipping of .env file eg. for Docker
	if os.Getenv("SKIP_ENV") != "true" {
		err := godotenv.Load(".env")
		logger.SetLogLevel(os.Getenv("LOG_LEVEL"))
		if err != nil {
			logger.Warn(fmt.Sprintf("Failed to load env file: %s", err.Error()))
		}
	} else {
		logger.SetLogLevel(os.Getenv("LOG_LEVEL"))
	}
}

func VerifyEnv() {
	redditClientId := os.Getenv("REDDIT_CLIENT_ID")
	redditClientSecret := os.Getenv("REDDIT_CLIENT_SECRET")
	if redditClientId == "" || redditClientSecret == "" {
		logger.FatalAndExit("Missing REDDIT_CLIENT_ID and/or REDDIT_CLIENT_SECRET")
	}

	discordBotToken := os.Getenv("DISCORD_BOT_TOKEN")
	discordChannelId := os.Getenv("DISCORD_CHANNEL_ID")
	if discordBotToken == "" || discordChannelId == "" {
		logger.FatalAndExit("Missing DISCORD_BOT_TOKEN and/or DISCORD_CHANNEL_ID")
	}
}

func PrettyPrintDuration(d time.Duration) string {
	hours := d / time.Hour
	d -= hours * time.Hour
	minutes := d / time.Minute
	d -= minutes * time.Minute
	seconds := d / time.Second

	if hours > 0 {
		return fmt.Sprintf("%dh", hours)
	} else if minutes > 0 {
		return fmt.Sprintf("%dm", minutes)
	} else {
		return fmt.Sprintf("%ds", seconds)
	}
}

func BuildNotifyMessage(posts []redditapi.UserSubmittedEntry) string {
	if len(posts) == 0 {
		return ""
	}

	msg := fmt.Sprintf("@everyone: Found %v new post(s):\n", len(posts))

	for _, post := range posts {
		msg += fmt.Sprintf(
			"[%s](%s) posted in [%s](https://reddit.com/%s): [%s](%s)\n",
			post.Author.Name,
			post.Author.Uri,
			post.Category.Label,
			post.Category.Label,
			post.Title,
			post.Link.Href,
		)
	}

	return msg
}
