package logger

import (
	"fmt"
	"time"
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

func (l LogLevel) Color() string {
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
	if level < logLevel {
		return
	}

	t := time.Now().Format(time.RFC3339)
	fmt.Printf("%s%s [%s] %s\033[0m\n", level.Color(), t, level.String(), msg)
}
