package logger

import (
	"fmt"
	"os"
	"strings"
)

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

var logLevel LogLevel = DEBUG

func SetLogLevel(level string) {
	switch strings.ToLower(level) {
	case "":
	case "dbg":
	case "debug":
		logLevel = DEBUG
		return
	case "inf":
	case "info":
		logLevel = INFO
		return
	case "wrn":
	case "warn":
		logLevel = WARN
		return
	case "err":
	case "error":
		logLevel = ERROR
		return
	default:
		Warn(fmt.Sprintf("Unknown log level %s. Using DEBUG", level))
		return
	}
}

func Debug(msg string) {
	log(DEBUG, msg)
}

func Info(msg string) {
	log(INFO, msg)
}

func Warn(msg string) {
	log(WARN, msg)
}

func Error(msg string) {
	log(ERROR, msg)
}

func FatalAndExit(msg string) {
	log(ERROR, msg)
	os.Exit(1)
}
