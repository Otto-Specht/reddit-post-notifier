package logger

import (
	"fmt"
	"os"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DBG"
	case INFO:
		return "INF"
	case WARN:
		return "WRN"
	case ERROR:
		return "ERR"
	case FATAL:
		return "FTL"
	default:
		return "UKN"
	}
}

func color(l LogLevel) string {
	switch l {
	case DEBUG:
		return "\033[37m"
	case INFO:
		return "\033[36m"
	case WARN:
		return "\033[33m"
	case ERROR:
		return "\033[31m"
	case FATAL:
		return "\033[31m"
	default:
		return ""
	}
}

func log(level LogLevel, msg string) {
	fmt.Printf("%s[%s] %s\033[0m\n", color(level), level.String(), msg)
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
