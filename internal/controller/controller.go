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
	updateLatestPostIdsForUsers(userNames)

	interval := getPollInterval()
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	logger.Info(fmt.Sprintf("Checking %v user(s) every %s.", len(userNames), util.PrettyPrintDuration(interval)))

	for range ticker.C {
		logger.Info("Checking for new posts...")

		var newPostsAllUsers []api.UserSubmittedEntry

		for _, userlastPost := range lastPostIdPerUser {
			newPosts := getNewEntries(userlastPost)

			if len(newPosts) == 0 {
				logger.Debug(fmt.Sprintf("No new posts from %s", userlastPost.User))
				continue
			}

			logger.Debug(fmt.Sprintf("%v new posts from %s", len(newPosts), userlastPost.User))
			newPostsAllUsers = append(newPostsAllUsers, newPosts...)

		}

		if len(newPostsAllUsers) == 0 {
			logger.Info("Found no new posts")
			continue
		}

		logger.Info(fmt.Sprintf("Found %v new posts", len(newPostsAllUsers)))

		// TODO: Handle new post
	}
}
