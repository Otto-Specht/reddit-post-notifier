package controller

import (
	"fmt"
	"time"

	"github.com/Otto-Specht/reddit-post-notifier/internal/api"
	"github.com/Otto-Specht/reddit-post-notifier/pkg/logger"
	"github.com/Otto-Specht/reddit-post-notifier/pkg/util"
)

var lastPostIdPerUser = []UserPostId{}

func Start(userNames []string) {
	for _, userName := range userNames {
		var lastPost api.UserSubmittedEntry

		feed, err := api.GetUserFeed(userName, 1)
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

	interval := getPollInterval()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	logger.Info(fmt.Sprintf("Checking %v user(s) every %s.", len(userNames), util.PrettyPrintDuration(interval)))
	for {
		select {
		case <-ticker.C:
			logger.Debug("Starting job..")
		}
	}
}
