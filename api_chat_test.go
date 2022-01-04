package lark

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChatInfo(t *testing.T) {
	resp, err := bot.GetChat(testGroupChatID)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.Equal(t, "go-lark-ci", resp.Data.Name)
		assert.Equal(t, "group", resp.Data.ChatMode)
		t.Log(resp.Data)
	}
}

func TestChatCRUD(t *testing.T) {
	resp, err := bot.CreateChat(CreateChatRequest{
		Name: fmt.Sprintf("go-lark-ci-%d", time.Now().Unix()),
	})
	if assert.NoError(t, err) {
		chatID := resp.Data.ChatID
		assert.NotEmpty(t, chatID)
		_, err = bot.DeleteChat(chatID)
		assert.NoError(t, err)
	}
}
