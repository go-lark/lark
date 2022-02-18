package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandURL(t *testing.T) {
	bot := NewChatBot("test-id", "test-secret")
	bot.SetDomain("http://localhost")
	assert.Equal(t, bot.ExpandURL("/test"),
		"http://localhost/test")
}
