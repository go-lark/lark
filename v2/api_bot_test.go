package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBotInfo(t *testing.T) {
	resp, err := bot.GetBotInfo(t.Context())
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.Equal(t, "go-lark-bot", resp.Bot.AppName)
		t.Log(resp)
	}
}
