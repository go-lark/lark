package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuzzMessage(t *testing.T) {
	resp, err := bot.PostText(t.Context(), "this text will be buzzed", WithEmail(testUserEmail))
	if assert.NoError(t, err) {
		bot.WithUserIDType(UIDOpenID)
		messageID := resp.Data.MessageID
		buzzResp, err := bot.BuzzMessage(t.Context(), BuzzTypeInApp, messageID, testUserOpenID)
		if assert.NoError(t, err) {
			assert.Equal(t, 0, buzzResp.Code)
		}
	}
}
