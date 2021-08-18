package mlog

import (
	"fmt"
	"log"
	"time"
)

const (
	LogLevelDebug = iota
	LoglevelInfo
	LogLevelWarn
	LogLevelError
	LogLevelFatal
	LogLevelOff
)

var (
	DebugColor = ""
	InfoColor  = ""
	WarnColor  = ""
	ErrorColor = ""
	FatalColor = ""
)

func init() {
	DebugColor = fmt.Sprintf("%c[1;0;34m[DEBUG]%c[0m", 0x1B, 0x1B)
	InfoColor = fmt.Sprintf("%c[1;0;32m[INFO]%c[0m", 0x1B, 0x1B)
	WarnColor = fmt.Sprintf("%c[1;0;33m[WARNING]%c[0m", 0x1B, 0x1B)
	ErrorColor = fmt.Sprintf("%c[1;0;31m[ERROR]%c[0m", 0x1B, 0x1B)
	FatalColor = fmt.Sprintf("%c[1;0;30m[FATAL]%c[0m", 0x1B, 0x1B)
}

type ConsoleLog struct {
	Level int
}

var logger ConsoleLog

func NewLoggerWithOptions(level string) {
	logger = ConsoleLog{Level: levelAdapt(level)}
}

func Debug(format string, ps ...interface{}) {
	if logger.Level <= LogLevelDebug {
		fmt.Printf("%s[%s]|%s\n", DebugColor, curHumanTime(), fmt.Sprintf(format, ps...))
	}
}

func Info(format string, ps ...interface{}) {
	if logger.Level <= LoglevelInfo {
		fmt.Printf("%s[%s]|%s\n", InfoColor, curHumanTime(), fmt.Sprintf(format, ps...))
	}
}

func Warn(format string, ps ...interface{}) {
	if logger.Level <= LogLevelWarn {
		fmt.Printf("%s[%s]|%s\n", WarnColor, curHumanTime(), fmt.Sprintf(format, ps...))
	}
}

func Error(format string, ps ...interface{}) {
	if logger.Level <= LogLevelError {
		fmt.Printf("%s[%s]|%s\n", ErrorColor, curHumanTime(), fmt.Sprintf(format, ps...))
	}
}

func Fatal(format string, ps ...interface{}) {
	if logger.Level <= LogLevelFatal {
		fmt.Printf("%s[%s]|%s\n", FatalColor, curHumanTime(), fmt.Sprintf(format, ps...))
	}
}

func curHumanTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// levelAdapt adapts m logger level to zap logger level
func levelAdapt(level string) int {
	switch level {
	case "debug":
		return LogLevelDebug
	case "info":
		return LoglevelInfo
	case "warn":
		return LogLevelWarn
	case "error":
		return LogLevelError
	case "fatal":
		return LogLevelFatal
	default:
		log.Fatalf("unknown level %s", level)
	}
	return 0
}
