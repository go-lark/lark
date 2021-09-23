package lark

// go-lark doesn't have a specific logger
// instead, we use log.Logger

import (
	"bytes"
	"log"
	"os"
)

// LogPrefix for lark
const LogPrefix = "[go-lark] "

func initDefaultLogger() *log.Logger {
	// create a default std logger
	logger := log.New(os.Stderr, LogPrefix, log.LstdFlags)
	return logger
}

// SetLogger set a new logger
func (bot *Bot) SetLogger(logger *log.Logger) {
	bot.logger = logger
}

// Logger returns current logger
func (bot Bot) Logger() *log.Logger {
	return bot.logger
}

func (bot *Bot) captureOutput(f func()) string {
	var buf bytes.Buffer
	bot.logger.SetOutput(&buf)
	f()
	bot.logger.SetOutput(os.Stderr)
	return buf.String()
}
