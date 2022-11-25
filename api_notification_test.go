package lark

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWebhookV1(t *testing.T) {
	bot := NewNotificationBot(testWebhookV1)
	resp, err := bot.PostNotification("", "no title message")
	assert.NoError(t, err)
	assert.True(t, resp.Ok)
	_, err = bot.PostNotification("go-lark CI", "it works")
	assert.NoError(t, err)
}

// A weird case which sends V2 message body with V1 URL
func TestWebhookV1Error(t *testing.T) {
	bot := NewNotificationBot(testWebhookV1)
	mbText := NewMsgBuffer(MsgText)
	mbText.Text("hello")
	resp, err := bot.PostNotificationV2(mbText.Build())
	assert.NoError(t, err)
	assert.Zero(t, resp.StatusCode)
}

func TestWebhookV2(t *testing.T) {
	bot := NewNotificationBot(testWebhookV2)

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

func TestWebhookV2CardMessage(t *testing.T) {
	bot := NewNotificationBot(testWebhookV2)

	b := NewCardBuilder()
	card := b.Card(
		b.Div(
			b.Field(b.Text("左侧内容")).Short(),
			b.Field(b.Text("右侧内容")).Short(),
			b.Field(b.Text("整排内容")),
			b.Field(b.Text("整排**Markdown**内容").LarkMd()),
		),
		b.Div().
			Text(b.Text("Text Content")).
			Extra(b.Img("img_a7c6aa35-382a-48ad-839d-d0182a69b4dg")),
		b.Note().
			AddText(b.Text("Note **Text**").LarkMd()).
			AddImage(b.Img("img_a7c6aa35-382a-48ad-839d-d0182a69b4dg")),
	).
		Wathet().
		Title("Notification Card")
	msgV4 := NewMsgBuffer(MsgInteractive)
	omV4 := msgV4.Card(card.String()).Build()
	resp, err := bot.PostNotificationV2(omV4)

	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.StatusCode)
		assert.NotEmpty(t, resp.StatusMessage)
	}
}

func TestWebhookV2Signed(t *testing.T) {
	bot := NewNotificationBot(testWebhookV2Signed)

	mbText := NewMsgBuffer(MsgText)
	mbText.Text("hello sign").WithSign("LIpnNexV7rwOyOebKoqSdb", time.Now().Unix())
	resp, err := bot.PostNotificationV2(mbText.Build())
	assert.NoError(t, err)
	assert.Zero(t, resp.StatusCode)
	assert.Equal(t, "success", resp.StatusMessage)
}

func TestWebhookV2SignedError(t *testing.T) {
	bot := NewNotificationBot("https://open.feishu.cn/open-apis/bot/v2/hook/749be902-6eaa-4cc3-9325-be4126164b02")

	mbText := NewMsgBuffer(MsgText)
	mbText.Text("hello sign").WithSign("LIpnNexV7rwOyOebKoqSdb", time.Now().Unix())
	resp, err := bot.PostNotificationV2(mbText.Build())
	assert.NoError(t, err)
	assert.Zero(t, resp.StatusCode)
	assert.Equal(t, "sign match fail or timestamp is not within one hour from current time", resp.Msg)
}
