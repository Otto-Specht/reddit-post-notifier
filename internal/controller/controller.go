package controller

import (
	"fmt"
	"time"

	"github.com/Otto-Specht/reddit-post-notifier/internal/discordapi"
	"github.com/Otto-Specht/reddit-post-notifier/internal/redditapi"
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

		var newPostsAllUsers []redditapi.UserSubmittedEntry

		for idx, userlastPost := range lastPostIdPerUser {
			newPosts := getNewEntries(userlastPost)

			if len(newPosts) == 0 {
				logger.Debug(fmt.Sprintf("No new posts from %s", userlastPost.User))
				continue
			}

			logger.Debug(fmt.Sprintf("%v new posts from %s", len(newPosts), userlastPost.User))
			newPostsAllUsers = append(newPostsAllUsers, newPosts...)

			lastPostIdPerUser[idx] = UserPostId{User: userlastPost.User, PostId: newPosts[0].Id}
		}

		if len(newPostsAllUsers) == 0 {
			logger.Info("Found no new posts")
			continue
		}

		logger.Info(fmt.Sprintf("Found %v new posts. Sending discord message...", len(newPostsAllUsers)))

		discordapi.SendServerMessage(util.BuildNotifyMessage(newPostsAllUsers))
	}
}
