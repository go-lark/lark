package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetLogger(t *testing.T) {
	bot := newTestBot()
	newLogger := initDefaultLogger()
	bot.SetLogger(newLogger)
	assert.Equal(t, newLogger, bot.logger)
	assert.Equal(t, newLogger, bot.Logger())
}

func TestLogLevel(t *testing.T) {
	var logLevel LogLevel = LogLevelDebug
	assert.Equal(t, "DEBUG", logLevel.String())
	logLevel = LogLevelError
	assert.Equal(t, "ERROR", logLevel.String())
	logLevel = LogLevelTrace
	assert.Equal(t, "TRACE", logLevel.String())
	logLevel = LogLevelWarn
	assert.Equal(t, "WARN", logLevel.String())
	logLevel = LogLevelInfo
	assert.Equal(t, "INFO", logLevel.String())
	logLevel = 1000
	assert.Equal(t, "", logLevel.String())
}
