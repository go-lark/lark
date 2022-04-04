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
		Name:     fmt.Sprintf("go-lark-ci-%d", time.Now().Unix()),
		ChatMode: "group",
		ChatType: "public",
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
				// join chat
				joinResp, err := bot.JoinChat(chatID)
				assert.Zero(t, joinResp.Code)
				assert.NoError(t, err)

				// add chat member
				addMemberResp, err := bot.AddChatMember(chatID, []string{testUserOpenID})
				if assert.NoError(t, err) {
					assert.Equal(t, 0, addMemberResp.Code)
					assert.Empty(t, addMemberResp.Data.InvalidIDList)
				}
				// remove chat member
				removeMemberResp, err := bot.RemoveChatMember(chatID, []string{testUserOpenID})
				if assert.NoError(t, err) {
					assert.Equal(t, 0, removeMemberResp.Code)
					assert.Empty(t, removeMemberResp.Data.InvalidIDList)
				}

				// delete
				_, err = bot.DeleteChat(chatID)
				assert.NoError(t, err)
			}
		}
	}
}

func TestIsInChat(t *testing.T) {
	resp, err := bot.IsInChat(testGroupChatID)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.True(t, resp.Data.IsInChat)
	}
}

func TestGetChatMembers(t *testing.T) {
	bot.WithUserIDType(UIDOpenID)
	resp, err := bot.GetChatMembers(testGroupChatID, "", 1)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.Items)
		assert.NotEmpty(t, resp.Data.PageToken)
		assert.NotEmpty(t, resp.Data.MemberTotal)
		if assert.True(t, resp.Data.HasMore) {
			nextResp, err := bot.GetChatMembers(testGroupChatID, resp.Data.PageToken, 1)
			if assert.NoError(t, err) {
				assert.Equal(t, 0, nextResp.Code)
				t.Log(nextResp.Data.Items)
			}
		}
	}
}
