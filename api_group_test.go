package lark

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetGroupList(t *testing.T) {
	resp, err := bot.GetGroupList(1, 10)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		t.Log(resp.Chats)
		assert.NotEmpty(t, resp.Chats)
	}
}

func TestGetGroupInfo(t *testing.T) {
	resp, err := bot.GetGroupInfo(testGroupChatID)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.Equal(t, "go-lark-ci", resp.Data.Name)
		assert.NotEmpty(t, resp.Data.Members)
	}
}

func TestGroupCRUD(t *testing.T) {
	// create group
	createResp, err := bot.CreateGroup(
		fmt.Sprintf("go-lark-ci group %d", time.Now().Unix()),
		"group create",
		[]string{testUserOpenID},
	)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, createResp.Code)
	}
	groupChatID := createResp.OpenChatID
	// delete member
	delResp, err := bot.DeleteGroupMember(groupChatID, []string{testUserOpenID})
	if assert.NoError(t, err) {
		assert.Equal(t, 0, delResp.Code)
	}
	// add member
	addResp, err := bot.AddGroupMember(groupChatID, []string{testUserOpenID})
	if assert.NoError(t, err) {
		assert.Equal(t, 0, addResp.Code)
	}
	// delete again
	delResp, err = bot.DeleteGroupMember(groupChatID, []string{testUserOpenID})
	if assert.NoError(t, err) {
		assert.Equal(t, 0, delResp.Code)
	}
	// add by user id
	addByUserResp, err := bot.AddGroupMemberByUserID(groupChatID, []string{testUserID})
	if assert.NoError(t, err) {
		assert.Equal(t, 0, addByUserResp.Code)
	}
	// update info
	updateResp, err := bot.UpdateGroupInfo(&UpdateGroupInfoReq{
		OpenChatID: groupChatID,
		Name:       "test 1",
	})
	if assert.NoError(t, err) {
		assert.Equal(t, 0, updateResp.Code)
	}
	// disband
	disbandResp, err := bot.DisbandGroup(groupChatID)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, disbandResp.Code)
	}
}

func TestBotAddRemove(t *testing.T) {
	// rm bot
	rmBotResp, err := bot.RemoveBotFromGroup(testGroupChatID)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, rmBotResp.Code)
	}
	// add bot
	addBotResp, err := bot.AddBotToGroup(testGroupChatID)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, addBotResp.Code)
	}
}
