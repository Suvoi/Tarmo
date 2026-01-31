package logger

import (
	"fmt"
	"log"
)

func Info(msg string, args ...any) {
	log.Println(formatLog("INFO", msg, args...))
}

func Warn(msg string, args ...any) {
	log.Println(formatLog("WARN", msg, args...))
}

func Error(msg string, args ...any) {
	log.Println(formatLog("ERROR", msg, args...))
}

func Fatal(msg string, args ...any) {
	log.Fatal(formatLog("FATAL", msg, args...))
}

func formatLog(level string, msg string, args ...any) string {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}
	return fmt.Sprintf("[###%s###] %s", level, msg)
}
