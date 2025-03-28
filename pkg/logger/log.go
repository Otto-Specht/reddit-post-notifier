package logger

import (
	"fmt"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

func (l LogLevel) String() string {
	switch l {
	case INFO:
		return "INF"
	case WARN:
		return "WRN"
	case ERROR:
		return "ERR"
	case DEBUG:
		return "DBG"
	default:
		return "UKN"
	}
}

func log(level LogLevel, msg string) {
	fmt.Printf("[%s] %s\n", level.String(), msg)
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
