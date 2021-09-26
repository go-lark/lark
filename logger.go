package lark

import (
	"bytes"
	"io"
	"log"
	"os"
)

// LogLevel defs
type LogLevel int

// LogLevels
const (
	LogLevelTrace = iota + 1
	LogLevelDebug
	LogLevelInfo
	LogLevelWarn
	LogLevelError
)

// LogWrapper interface
type LogWrapper interface {
	// for log print
	Log(LogLevel, string)
	// for test redirection
	SetOutput(io.Writer)
}

// String .
func (ll LogLevel) String() string {
	switch ll {
	case LogLevelTrace:
		return "TRACE"
	case LogLevelDebug:
		return "DEBUG"
	case LogLevelInfo:
		return "INFO"
	case LogLevelWarn:
		return "WARN"
	case LogLevelError:
		return "ERROR"
	}
	return ""
}

type stdLogger struct {
	*log.Logger
}

func (sl stdLogger) Log(level LogLevel, msg string) {
	sl.Printf("[%s] %s\n", level, msg)
}

const logPrefix = "[go-lark] "

func initDefaultLogger() LogWrapper {
	// create a default std logger
	logger := stdLogger{
		log.New(os.Stderr, logPrefix, log.LstdFlags),
	}
	return logger
}

// SetLogger set a new logger
func (bot *Bot) SetLogger(logger LogWrapper) {
	bot.logger = logger
}

// Logger returns current logger
func (bot Bot) Logger() LogWrapper {
	return bot.logger
}

func (bot Bot) captureOutput(f func()) string {
	var buf bytes.Buffer
	bot.logger.SetOutput(&buf)
	f()
	bot.logger.SetOutput(os.Stderr)
	return buf.String()
}
