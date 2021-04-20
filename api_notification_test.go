package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	ciHookURLV1 = "https://open.feishu.cn/open-apis/bot/hook/e197c5f5e65f4778b9e7a89bf23a5d4c"
	ciHookURLV2 = "https://open.feishu.cn/open-apis/bot/v2/hook/7b01451f-113b-4296-8f0d-9615499d6545"
)

func TestWebhookV1(t *testing.T) {
	bot := NewNotificationBot(ciHookURLV1)
	resp, err := bot.PostNotification("", "no title message")
	assert.NoError(t, err)
	assert.True(t, resp.Ok)
	_, err = bot.PostNotification("go-lark CI", "it works")
	assert.NoError(t, err)
}

// A weird case which sends V2 message body with V1 URL
func TestWebhookV1Error(t *testing.T) {
	bot := NewNotificationBot(ciHookURLV1)
	mbText := NewMsgBuffer(MsgText)
	mbText.Text("hello")
	resp, err := bot.PostNotificationV2(mbText.Build())
	assert.NoError(t, err)
	assert.Zero(t, resp.StatusCode)
}

func TestWebhookV2(t *testing.T) {
	bot := NewNotificationBot(ciHookURLV2)

	mbText := NewMsgBuffer(MsgText)
	mbText.Text("hello")
	resp, err := bot.PostNotificationV2(mbText.Build())
	assert.NoError(t, err)
	assert.Zero(t, resp.StatusCode)
	assert.Equal(t, "success", resp.StatusMessage)

	mbPost := NewMsgBuffer(MsgPost)
	mbPost.Post(NewPostBuilder().Title("hello").TextTag("world", 1, true).Render())
	resp, err = bot.PostNotificationV2(mbPost.Build())
	assert.NoError(t, err)
	assert.Zero(t, resp.StatusCode)
	assert.Equal(t, "success", resp.StatusMessage)

	mbImg := NewMsgBuffer(MsgImage)
	mbImg.Image("img_a97c1597-9c0a-47c1-9fb4-dd3e5e37ac9g")
	resp, err = bot.PostNotificationV2(mbImg.Build())
	assert.NoError(t, err)
	assert.Zero(t, resp.StatusCode)
	assert.Equal(t, "success", resp.StatusMessage)

	mbShareGroup := NewMsgBuffer(MsgShareCard)
	mbShareGroup.ShareChat(testGroupChatID)
	resp, err = bot.PostNotificationV2(mbShareGroup.Build())
	assert.NoError(t, err)
	assert.Zero(t, resp.StatusCode)
	assert.Equal(t, "success", resp.StatusMessage)
}

func TestWebhookV2Errors(t *testing.T) {
	bot := NewNotificationBot(ciHookURLV2)

	resp, err := bot.PostNotification("go-lark CI", "it works")
	assert.NoError(t, err)
	assert.False(t, resp.Ok)

	msg := NewMsgBuffer(MsgInteractive)
	respV2, err := bot.PostNotificationV2(msg.Build())
	assert.Nil(t, respV2)
	assert.Error(t, ErrMessageTypeNotSuppored, err)
}
