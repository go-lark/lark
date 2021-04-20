package lark

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGetGroupList(t *testing.T) {
	bot := newTestBot()
	_, _ = bot.GetTenantAccessTokenInternal(true)
	resp, err := bot.GetGroupList(1, 10)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		t.Log(resp.Chats)
		assert.NotEmpty(t, resp.Chats)
	}
}

func TestGetGroupInfo(t *testing.T) {
	bot := newTestBot()
	_, _ = bot.GetTenantAccessTokenInternal(true)
	resp, err := bot.GetGroupInfo("oc_51e1f7ff2b3dbe8f74b7b22d5d44d14e")
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.Equal(t, "go-lark开发者交流群", resp.Data.Name)
		assert.NotEmpty(t, resp.Data.Members)
	}
}

func TestGroupCreateDisbandAddDelete(t *testing.T) {
	bot := newTestBot()
	_, _ = bot.GetTenantAccessTokenInternal(true)
	createResp, err := bot.CreateGroup("go-lark-ci create-group", "group create", []string{testUserOpenID})
	if assert.NoError(t, err) {
		assert.Equal(t, 0, createResp.Code)
		t.Log(createResp.OpenChatID)
	}

	addResp, err := bot.AddGroupMember(createResp.OpenChatID, []string{"ou_b4852f3f88e454522180a5fcb346648d"})
	if assert.NoError(t, err) {
		assert.Equal(t, 0, addResp.Code)
	}
	delResp, err := bot.DeleteGroupMember(createResp.OpenChatID, []string{"ou_b4852f3f88e454522180a5fcb346648d"})
	if assert.NoError(t, err) {
		assert.Equal(t, 0, delResp.Code)
	}
	addV4Resp, err := bot.AddGroupMemberByUserID(testGroupChatID, []string{"55677932"})
	if assert.NoError(t, err) {
		assert.Equal(t, 0, addV4Resp.Code)
	}

	disbandResp, err := bot.DisbandGroup(createResp.OpenChatID)
	assert.Nil(t, err)
	assert.Zero(t, disbandResp.Code)
}

func TestUpdateGroupInfo(t *testing.T) {
	bot := newTestBot()
	_, _ = bot.GetTenantAccessTokenInternal(true)
	_, err := bot.UpdateGroupInfo(&UpdateGroupInfoReq{
		OpenChatID: "oc_dae15ce1b55b00ecf853c50336a51e55",
		Name:       "test 1",
	})
	assert.Nil(t, err)
}

func TestAddRemoveBot(t *testing.T) {
	bot := newTestBot()
	_, _ = bot.GetTenantAccessTokenInternal(true)
	openChatID := "oc_dae15ce1b55b00ecf853c50336a51e55"

	_, err := bot.RemoveBotFromGroup(openChatID)
	assert.Nil(t, err)
	time.Sleep(time.Second)

	_, err = bot.AddBotToGroup(openChatID)
	assert.Nil(t, err)
}
