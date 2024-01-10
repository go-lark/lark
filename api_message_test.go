package lark

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

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
	resp, err = bot.PostText("PostText: with formats <b>Bold</b> <i>italic</i> <u>Underline</u> <s>strikethrough</s> [Link](https://bytedance.com/)", WithChatID(testGroupChatID))
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
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
	replyID := resp.Data.MessageID
	resp, err = bot.PostTextMentionAndReply("PostTextMentionAndReply", testUserOpenID, WithChatID(testGroupChatID), replyID)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
	// Reply with Outcoming Message
	mb := NewMsgBuffer(MsgText)
	tb := NewTextBuilder()
	om := mb.Text(tb.Text("Reply raw").Render()).BindReply(replyID).ReplyInThread(true).Build()
	resp, err = bot.ReplyMessage(om)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
}

func TestReplyMessageFailed(t *testing.T) {
	om := newMsgBufWithOptionalUserID(MsgText, &OptionalUserID{UIDOpenID, testUserOpenID}).Text("will fail").Build()
	_, err := bot.ReplyMessage(om)
	assert.ErrorIs(t, err, ErrParamMessageID)
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

func TestPostTextFailed(t *testing.T) {
	resp, err := bot.PostText("PostText: email hello, world", WithEmail("9999@example.com"))
	if assert.NoError(t, err) {
		assert.NotEqual(t, 0, resp.Code)
		assert.Contains(t, resp.Msg, "invalid receive_id")
	}
}

func TestPostFailedByUserID(t *testing.T) {
	_, err := bot.PostText("should fail", &OptionalUserID{"some id", ""})
	assert.ErrorIs(t, err, ErrParamUserID)
	_, err = bot.PostTextMention("should fail", "", &OptionalUserID{"some id", ""})
	assert.ErrorIs(t, err, ErrParamUserID)
	_, err = bot.PostTextMentionAll("should fail", &OptionalUserID{"some id", ""})
	assert.ErrorIs(t, err, ErrParamUserID)
	_, err = bot.PostTextMentionAndReply("should fail", "", &OptionalUserID{"some id", ""}, "")
	assert.ErrorIs(t, err, ErrParamUserID)
	_, err = bot.PostImage("should fail", &OptionalUserID{"some id", ""})
	assert.ErrorIs(t, err, ErrParamUserID)
	_, err = bot.PostShareChat("should fail", &OptionalUserID{"some id", ""})
	assert.ErrorIs(t, err, ErrParamUserID)
	_, err = bot.PostShareUser("should fail", &OptionalUserID{"some id", ""})
	assert.ErrorIs(t, err, ErrParamUserID)
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
	om = msg.BindOpenChatID(testGroupChatID).Image("img_v2_bd72e090-3e6a-42bf-b0b0-fead4a11c9cg").Build()
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
	om = msg.BindChatID(testGroupChatID).File("file_v3_0069_61085e72-3285-4ac8-82ec-31b060eeae7g").Build()
	resp, err = bot.PostMessage(om)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
	// audio - not actually tested
	msg = NewMsgBuffer(MsgAudio)
	om = msg.BindChatID(testGroupChatID).File("file_v3_0069_61085e72-3285-4ac8-82ec-31b060eeae7g").Build()
	resp, err = bot.PostMessage(om)
	assert.NotEqual(t, 0, resp.Code)
	assert.NotEmpty(t, resp.Msg)
	// media - not actually tested
	msg = NewMsgBuffer(MsgMedia)
	om = msg.BindChatID(testGroupChatID).Media("file_v3_0069_61085e72-3285-4ac8-82ec-31b060eeae7g", "img_v3_0269_ba0a83c6-21ff-48e2-a159-9061d2c8217g").Build()
	resp, err = bot.PostMessage(om)
	assert.NotEqual(t, 0, resp.Code)
	assert.NotEmpty(t, resp.Msg)
	// sticker
	msg = NewMsgBuffer(MsgSticker)
	om = msg.BindChatID(testGroupChatID).Sticker("4ba009df-2453-47b3-a753-444b152217bg").Build()
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

		newOM := NewMsgBuffer(MsgPost).
			BindOpenChatID(testGroupChatID).
			Post(
				NewPostBuilder().
					Title("modified title").
					TextTag("modified content", 1, true).
					Render(),
			).
			Build()
		_, err = bot.UpdateMessage(resp.Data.MessageID, newOM)
		if assert.NoError(t, err) {
			assert.Equal(t, 0, resp.Code)
		}
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
		Title("卡片标题 Card Title").
		UpdateMulti(true)
	msgV4 := NewMsgBuffer(MsgInteractive)
	omV4 := msgV4.BindEmail(testUserEmail).Card(card.String()).Build()
	resp, err := bot.PostMessage(omV4)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
		time.Sleep(time.Second * 1)
		newCard := NewCardBuilder().Card(
			b.Div(
				b.Field(b.Text("左侧内容")).Short(),
				b.Field(b.Text("右侧内容")).Short(),
				b.Field(b.Text("整排内容")),
				b.Field(b.Text("整排**Markdown**内容").LarkMd()),
			),
		).Title("Updated title")
		newOM := msgV4.BindEmail(testUserEmail).Card(newCard.String()).Build()
		resp, err := bot.UpdateMessage(resp.Data.MessageID, newOM)
		t.Log(err, resp)
	}
}

func TestColumnSet(t *testing.T) {
	b := NewCardBuilder()
	card := b.Card(
		b.ColumnSet(
			b.Column(
				b.Markdown("已审批单量\n**29单**\n<font color='green'>领先团队59%</font>").AlignCenter(),
			).Width("weighted").Weight(1),
			b.Column(
				b.Markdown("平均审批耗时\n**0.9小时**\n<font color='green'>领先团队100%</font>").AlignCenter(),
			).Width("weighted").Weight(1),
			b.Column(
				b.Markdown("待批率\n**25%**\n<font color='red'>落后团队29%</font>").AlignCenter(),
			).Width("weighted").Weight(1),
		).
			FlexMode("bisect").
			BackgroundStyle("grey").
			HorizontalSpacing("default"),
		b.Hr(),
		b.Markdown("**团队审批效率总览**"),
		b.ColumnSet(
			b.Column(
				b.Markdown("**审批人**\n王大明\n张军\n李小方"),
			).Width("weighted").Weight(1),
			b.Column(
				b.Markdown("**审批时长**\n小于1小时\n2小时\n3小时"),
			).Width("weighted").Weight(1),
			b.Column(
				b.Markdown("**对比上周变化**\n<font color='green'>↓12%</font>\n<font color='red'>↑5%</font>\n<font color='green'>↓25%</font>"),
			).Width("weighted").Weight(1),
		).Action(
			b.ColumnSetAction(
				b.URL().
					Href("https://open.feishu.cn").
					MultiHref(
						"https://developer.android.com/",
						"https://developer.apple.com/",
						"https://www.windows.com",
					),
			),
		),
	).
		Wathet().
		Title("Card with Column Set").
		UpdateMulti(true)
	msgV4 := NewMsgBuffer(MsgInteractive)
	omV4 := msgV4.BindEmail(testUserEmail).Card(card.String()).Build()
	resp, err := bot.PostMessage(omV4)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
}

func TestI18NCard(t *testing.T) {
	b := NewCardBuilder()
	card := b.I18N.
		Card(
			b.I18N.WithLocale(
				LocaleEnUS,
				b.Div(
					b.Field(b.Text("English Content")),
				),
			),
			b.I18N.WithLocale(
				LocaleZhCN,
				b.Div(
					b.Field(b.Text("中文内容")),
				),
			),
			b.I18N.WithLocale(
				LocaleJaJP,
				b.Div(
					b.Field(b.Text("日本語コンテンツ")),
				),
			),
		).
		Title(
			b.I18N.LocalizedText(LocaleEnUS, "English Title"),
			b.I18N.LocalizedText(LocaleZhCN, "中文标题"),
			b.I18N.LocalizedText(LocaleJaJP, "日本語タイトル"),
		).
		Red().
		UpdateMulti(true)
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
}

func TestIdempotentMessage(t *testing.T) {
	uuid := fmt.Sprintf("%s-%d", "abc", rand.Intn(999))
	msg := NewMsgBuffer(MsgText)
	om := msg.BindChatID(testGroupChatID).WithUUID(uuid).Text("hello, UUID").Build()
	resp, err := bot.PostMessage(om)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
	msg.Clear()
	om = msg.BindChatID(testGroupChatID).WithUUID(uuid).Text("goodbye, UUID").Build()
	_, _ = bot.PostMessage(om)
	// you will then only receive the first one
	msg.Clear()
	om = msg.BindChatID(testGroupChatID).Text("goodbye, without UUID").Build()
	_, _ = bot.PostMessage(om)
}

func TestPinMessages(t *testing.T) {
	msg := NewMsgBuffer(MsgText)
	om := msg.BindEmail(testUserEmail).Text("hello, world").Build()
	resp, err := bot.PostMessage(om)
	if assert.NoError(t, err) {
		messageID := resp.Data.MessageID
		resp, err := bot.PinMessage(messageID)
		if assert.NoError(t, err) {
			assert.Equal(t, 0, resp.Code)
			assert.Equal(t, messageID, resp.Data.Pin.MessageID)
			unpinResp, err := bot.UnpinMessage(messageID)
			assert.NoError(t, err)
			assert.Equal(t, 0, unpinResp.Code)
		}
	}
}

func TestMessageReactions(t *testing.T) {
	msg := NewMsgBuffer(MsgText)
	om := msg.BindEmail(testUserEmail).Text("hello, world").Build()
	resp, err := bot.PostMessage(om)
	if assert.NoError(t, err) {
		messageID := resp.Data.MessageID
		resp, err := bot.AddReaction(messageID, EmojiTypeOK)
		if assert.NoError(t, err) {
			assert.Equal(t, 0, resp.Code)
			assert.Equal(t, EmojiTypeOK, resp.Data.ReactionType.EmojiType)
			deleteReactionResp, err := bot.DeleteReaction(messageID, resp.Data.ReactionID)
			assert.NoError(t, err)
			assert.Equal(t, 0, deleteReactionResp.Code)
		}
	}
}

func TestForwardMessage(t *testing.T) {
	msg := NewMsgBuffer(MsgText)
	om := msg.BindEmail(testUserEmail).Text("let's forward").Build()
	resp, err := bot.PostMessage(om)
	if assert.NoError(t, err) {
		messageID := resp.Data.MessageID
		resp, err := bot.ForwardMessage(messageID, WithChatID(testGroupChatID))
		if assert.NoError(t, err) {
			assert.Equal(t, 0, resp.Code)
		}
	}
}
