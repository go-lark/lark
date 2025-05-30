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

func TestAPIError(t *testing.T) {
	resp, err := bot.PostText(t.Context(), "failing", WithChatID("1231"))
	t.Log(resp, err)
	assert.Error(t, err)
	assert.NotZero(t, resp.Code)
}
