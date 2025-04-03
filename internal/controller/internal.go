package controller

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Otto-Specht/reddit-post-notifier/internal/redditapi"
	"github.com/Otto-Specht/reddit-post-notifier/pkg/logger"
	"github.com/Otto-Specht/reddit-post-notifier/pkg/util"
)

func getPollInterval() time.Duration {
	defaultInterval := time.Minute * 15

	envInterval := os.Getenv("POLL_INTERVAL")
	if envInterval == "" {
		return defaultInterval
	}

	envIntervalValue, err := strconv.Atoi(envInterval[:len(envInterval)-1])
	envIntervalUnit := envInterval[len(envInterval)-1:]

	if envIntervalValue < 1 || err != nil || envIntervalUnit == "" {
		logger.Error(fmt.Sprintf("Invalid POLL_INTERVAL value (eg. 10s, 30m, 1h), using default: %s", util.PrettyPrintDuration(defaultInterval)))
		return defaultInterval
	}

	switch envIntervalUnit {
	case "s":
		return time.Second * time.Duration(envIntervalValue)
	case "m":
		return time.Minute * time.Duration(envIntervalValue)
	case "h":
		return time.Hour * time.Duration(envIntervalValue)
	default:
		logger.Error(fmt.Sprintf("Invalid POLL_INTERVAL unit (eg. 10s, 30m, 1h), using default: %s", util.PrettyPrintDuration(defaultInterval)))
		return defaultInterval
	}
}

func updateLatestPostIdsForUsers(userNames []string) {
	for _, userName := range userNames {
		var lastPost redditapi.UserSubmittedEntry

		feed, err := redditapi.GetUserFeed(userName, 1)
		if err == nil && len(feed.Entries) != 0 {
			lastPost = feed.Entries[0]
		}

		logger.Debug(fmt.Sprintf("Last post at %s from u/%s (%s)", lastPost.Published, userName, lastPost.Title))

		lastPostIdPerUser = append(lastPostIdPerUser,
			UserPostId{
				User:   userName,
				PostId: lastPost.Id,
			})
	}
}

func getNewEntries(userlastPost UserPostId) []redditapi.UserSubmittedEntry {
	limit := 5

	feed, err := redditapi.GetUserFeed(userlastPost.User, limit)
	if err != nil {
		logger.Warn(fmt.Sprintf("Skipping user %s", userlastPost.User))
		return []redditapi.UserSubmittedEntry{}
	}

	if len(feed.Entries) == 0 || feed.Entries[0].Id == userlastPost.PostId {
		return []redditapi.UserSubmittedEntry{}
	}

	var newEntries []redditapi.UserSubmittedEntry

	for _, entry := range feed.Entries {
		if entry.Id != userlastPost.PostId {
			newEntries = append(newEntries, entry)
		} else {
			return newEntries
		}
	}

	logger.Info(fmt.Sprintf("Found %v new posts for user %s. There might be more because the script only checks the %v latest posts.", limit, userlastPost.User, limit))
	return newEntries
}
