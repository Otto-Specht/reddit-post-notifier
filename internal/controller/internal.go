package controller

import (
	"fmt"
	"os"
	"strconv"
	"time"

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
