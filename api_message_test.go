package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostText(t *testing.T) {
	resp, err := bot.PostText("PostText: email hello, world", WithEmail(testUserEmail))
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
	resp, err = bot.PostText("PostText: open_id hello, world", WithOpenID(testUserOpenID))
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
	resp, err = bot.PostText("PostText: union_id hello, world", WithUnionID(testUserUnionID))
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
	resp, err = bot.PostText("PostText: chat_id hello, world", WithChatID(testGroupChatID))
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
}

func TestPostTextFailed(t *testing.T) {
	resp, err := bot.PostText("PostText: email hello, world", WithEmail("9999@example.com"))
	if assert.NoError(t, err) {
		assert.NotEqual(t, 0, resp.Code)
		assert.Contains(t, resp.Msg, "invalid receive_id")
	}
}

func TestPostTextMention(t *testing.T) {
	resp, err := bot.PostTextMention("PostTextMention", testUserOpenID, WithChatID(testGroupChatID))
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
}

func TestPostTextMentionAll(t *testing.T) {
	resp, err := bot.PostTextMentionAll("PostTextMentionAll", WithChatID(testGroupChatID))
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
}

func TestReplyMessage(t *testing.T) {
	resp, err := bot.PostText("Message to be replied", WithChatID(testGroupChatID))
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
	resp, err = bot.PostTextMentionAndReply("PostTextMentionAndReply", testUserOpenID, WithChatID(testGroupChatID), resp.Data.MessageID)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
}

func TestPostImage(t *testing.T) {
	resp, err := bot.PostImage("img_a97c1597-9c0a-47c1-9fb4-dd3e5e37ac9g", WithChatID(testGroupChatID))
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
}

func TestPostShareChat(t *testing.T) {
	resp, err := bot.PostShareChat(testGroupChatID, WithChatID(testGroupChatID))
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
}

func TestPostShareUser(t *testing.T) {
	resp, err := bot.PostShareUser(testUserOpenID, WithChatID(testGroupChatID))
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
}

func TestPostMessage(t *testing.T) {
	// text message
	msg := NewMsgBuffer(MsgText)
	om := msg.BindEmail(testUserEmail).Text("hello, world").Build()
	resp, err := bot.PostMessage(om)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
	// group text message
	msg = NewMsgBuffer(MsgText)
	om = msg.BindOpenChatID(testGroupChatID).Text("group: hello, world").Build()
	resp, err = bot.PostMessage(om)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
	// image
	msg = NewMsgBuffer(MsgImage)
	om = msg.BindOpenChatID(testGroupChatID).Image("96f394ba-fc6a-4f38-b515-7b8b98160012").Build()
	resp, err = bot.PostMessage(om)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
	// share chat
	msg = NewMsgBuffer(MsgShareCard)
	om = msg.BindOpenChatID(testGroupChatID).ShareChat(testGroupChatID).Build()
	resp, err = bot.PostMessage(om)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
	// share user
	msg = NewMsgBuffer(MsgShareUser)
	om = msg.BindChatID(testGroupChatID).ShareUser(testUserOpenID).Build()
	resp, err = bot.PostMessage(om)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
	// file
	msg = NewMsgBuffer(MsgFile)
	om = msg.BindChatID(testGroupChatID).File("file_v2_356de203-39c2-49fd-8357-70b9b311f44g").Build()
	resp, err = bot.PostMessage(om)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
}

func TestPostPostMessage(t *testing.T) {
	msg := NewMsgBuffer(MsgPost)
	postContent := NewPostBuilder().
		Title("中文标题").
		TextTag("你好世界", 1, true).
		TextTag("hello, world", 1, true).
		LinkTag("ByteDance", "https://bytedance.com/").
		AtTag("www", testGroupChatID).
		ImageTag("img_a7c6aa35-382a-48ad-839d-d0182a69b4dg", 300, 300).
		WithLocale(LocaleEnUS).
		Title("English Title").
		TextTag("hello, world", 1, true).
		Render()
	om := msg.BindOpenChatID(testGroupChatID).Post(postContent).Build()
	resp, err := bot.PostMessage(om)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
}

func TestPostCardMessage(t *testing.T) {
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
		b.Action(
			b.Button(b.Text("**Primary**").LarkMd()).Primary(),
			b.Button(b.Text("Confirm")).Confirm("Confirm", "Are you sure?"),
			b.Overflow(
				b.Option("Option 1"),
				b.Option("选项2"),
			).Value(map[string]interface{}{"k": "v"}),
		).TrisectionLayout(),
		b.Action(
			b.SelectMenu(
				b.Option("Option 1"),
				b.Option("选项2"),
			).
				Placeholder("select").
				Value(map[string]interface{}{"k": "v"}),
		),
		b.Note().
			AddText(b.Text("Note **Text**").LarkMd()).
			AddImage(b.Img("img_a7c6aa35-382a-48ad-839d-d0182a69b4dg")),
	).
		Wathet().
		Title("卡片标题 Card Title")
	msgV4 := NewMsgBuffer(MsgInteractive)
	omV4 := msgV4.BindEmail(testUserEmail).Card(card.String()).Build()
	resp, err := bot.PostMessage(omV4)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
}

func TestEphemeralMessage(t *testing.T) {
	b := NewCardBuilder()
	card := b.Card(
		b.Div().
			Text(b.Text("Text Content")),
	).
		Wathet().
		Title("Ephemeral Card")
	msg := NewMsgBuffer(MsgInteractive)
	om := msg.BindChatID(testGroupChatID).BindEmail(testUserEmail).Card(card.String()).Build()
	resp, err := bot.PostEphemeralMessage(om)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}

	delResp, err := bot.DeleteEphemeralMessage(resp.Data.MessageID)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, delResp.Code)
	}
}

func TestMessageCRUD(t *testing.T) {
	// create
	resp, err := bot.PostText("PostText: CRUD hello, world", WithChatID(testGroupChatID))
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
		t.Log(resp.Data)
	}
	// get
	getResp, err := bot.GetMessage(resp.Data.MessageID)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, getResp.Code)
		if assert.NotEmpty(t, getResp.Data.Items) {
			assert.Equal(t, getResp.Data.Items[0].MessageID, resp.Data.MessageID)
		}
	}
	// recall
	rcResp, err := bot.RecallMessage(resp.Data.MessageID)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, rcResp.Code)
	}
	// receipt unread
	receipt, err := bot.MessageReadReceipt(resp.Data.MessageID)
	if assert.NoError(t, err) {
		t.Log(receipt.Data.ReadUsers)
	}
	// receipt read
	receiptOld, err := bot.MessageReadReceipt(testMessageID)
	if assert.NoError(t, err) {
		assert.NotEmpty(t, receiptOld.Data.ReadUsers)
		t.Log(receiptOld.Data.ReadUsers)
	}
}
