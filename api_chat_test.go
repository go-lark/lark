package lark

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestChatInfo(t *testing.T) {
	bot.WithUserIDType(UIDOpenID)
	assert.Equal(t, UIDOpenID, bot.userIDType)
	resp, err := bot.GetChat(testGroupChatID)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.Equal(t, "go-lark-ci", resp.Data.Name)
		assert.Equal(t, "group", resp.Data.ChatMode)
		assert.Equal(t, testUserOpenID, resp.Data.OwnerID)
		t.Log(resp.Data)
	}
}

func TestChatCRUD(t *testing.T) {
	bot.WithUserIDType(UIDOpenID)
	resp, err := bot.CreateChat(CreateChatRequest{
		Name: fmt.Sprintf("go-lark-ci-%d", time.Now().Unix()),
	})
	if assert.NoError(t, err) {
		chatID := resp.Data.ChatID
		assert.NotEmpty(t, chatID)
		upResp, err := bot.UpdateChat(chatID, UpdateChatRequest{
			Description: "new description",
		})
		t.Log(upResp)
		if assert.NoError(t, err) {
			getResp, err := bot.GetChat(chatID)
			if assert.NoError(t, err) {
				assert.Equal(t, "new description", getResp.Data.Description)
				_, err = bot.DeleteChat(chatID)
				assert.NoError(t, err)
			}
		}
	}
}
