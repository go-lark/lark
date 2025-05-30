package lark

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWebhook(t *testing.T) {
	bot := NewNotificationBot(testWebhook)

	mbText := NewMsgBuffer(MsgText)
	mbText.Text("hello")
	resp, err := bot.PostNotification(t.Context(), mbText.Build())
	assert.NoError(t, err)
	assert.Zero(t, resp.StatusCode)
	assert.Equal(t, "success", resp.StatusMessage)

	mbPost := NewMsgBuffer(MsgPost)
	mbPost.Post(NewPostBuilder().Title("hello").TextTag("world", 1, true).Render())
	resp, err = bot.PostNotification(t.Context(), mbPost.Build())
	assert.NoError(t, err)
	assert.Zero(t, resp.StatusCode)
	assert.Equal(t, "success", resp.StatusMessage)

	mbImg := NewMsgBuffer(MsgImage)
	mbImg.Image("img_a97c1597-9c0a-47c1-9fb4-dd3e5e37ac9g")
	resp, err = bot.PostNotification(t.Context(), mbImg.Build())
	assert.NoError(t, err)
	assert.Zero(t, resp.StatusCode)
	assert.Equal(t, "success", resp.StatusMessage)

	mbShareGroup := NewMsgBuffer(MsgShareCard)
	mbShareGroup.ShareChat(testGroupChatID)
	resp, err = bot.PostNotification(t.Context(), mbShareGroup.Build())
	assert.NoError(t, err)
	assert.Zero(t, resp.StatusCode)
	assert.Equal(t, "success", resp.StatusMessage)
}

func TestWebhookCardMessage(t *testing.T) {
	bot := NewNotificationBot(testWebhook)

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
	msg := NewMsgBuffer(MsgInteractive)
	om := msg.Card(card.String()).Build()
	resp, err := bot.PostNotification(t.Context(), om)

	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.StatusCode)
		assert.NotEmpty(t, resp.StatusMessage)
	}
}

func TestWebhookSigned(t *testing.T) {
	bot := NewNotificationBot(testWebhookSigned)

	mbText := NewMsgBuffer(MsgText)
	mbText.Text("hello sign").WithSign("FT1dnAgPYYTcpafMTkhPjc", time.Now().Unix())
	resp, err := bot.PostNotification(t.Context(), mbText.Build())
	assert.NoError(t, err)
	assert.Zero(t, resp.StatusCode)
	assert.Equal(t, "success", resp.StatusMessage)
}

func TestWebhookSignedError(t *testing.T) {
	bot := NewNotificationBot("https://open.feishu.cn/open-apis/bot/v2/hook/749be902-6eaa-4cc3-9325-be4126164b02")

	mbText := NewMsgBuffer(MsgText)
	mbText.Text("hello sign").WithSign("LIpnNexV7rwOyOebKoqSdb", time.Now().Unix())
	resp, err := bot.PostNotification(t.Context(), mbText.Build())
	assert.NoError(t, err)
	assert.Zero(t, resp.StatusCode)
	assert.Equal(t, "sign match fail or timestamp is not within one hour from current time", resp.Msg)
}
