package controller

import (
	"fmt"
	"time"

	"github.com/Otto-Specht/reddit-post-notifier/pkg/logger"
	"github.com/Otto-Specht/reddit-post-notifier/pkg/util"
)

func Start(userNames []string) {
	logger.Debug("Starting controller...")

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
