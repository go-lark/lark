package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttachText(t *testing.T) {
	mb := NewMsgBuffer(MsgText)
	msg := mb.Text("hello").Build()
	assert.Equal(t, "hello", msg.Content.Text.Text)
}

func TestAttachImage(t *testing.T) {
	mb := NewMsgBuffer(MsgImage)
	msg := mb.Image("aaaaa").Build()
	assert.Equal(t, "aaaaa", msg.Content.Image.ImageKey)
}

func TestMsgTextBinding(t *testing.T) {
	mb := NewMsgBuffer(MsgText)
	msg := mb.Text("hello, world").BindEmail(testUserEmail).Build()
	assert.Equal(t, "hello, world", msg.Content.Text.Text)
	assert.Equal(t, testUserEmail, msg.Email)
}

func TestBindingUserIDs(t *testing.T) {
	mb := NewMsgBuffer(MsgText)
	msgEmail := mb.BindEmail(testUserEmail).Build()
	assert.Equal(t, testUserEmail, msgEmail.Email)

	mb.Clear()
	msgOpenChatID := mb.BindOpenChatID(testGroupChatID).Build()
	assert.Equal(t, testGroupChatID, msgOpenChatID.ChatID)

	mb.Clear()
	msgUserID := mb.BindUserID("333444").Build()
	assert.Equal(t, "333444", msgUserID.UserID)

	mb.Clear()
	msgReplyID := mb.BindReply(testMessageID).Build()
	assert.Equal(t, testMessageID, msgReplyID.RootID)
}

func TestMsgShareChat(t *testing.T) {
	mb := NewMsgBuffer(MsgShareCard)
	msg := mb.ShareChat("6559399282837815565").Build()
	assert.Equal(t, MsgShareCard, msg.MsgType)
	assert.Equal(t, "6559399282837815565", msg.Content.ShareChat.ChatID)
}

func TestMsgShareUser(t *testing.T) {
	mb := NewMsgBuffer(MsgShareUser)
	msg := mb.ShareUser("334455").Build()
	assert.Equal(t, MsgShareUser, msg.MsgType)
	assert.Equal(t, "334455", msg.Content.ShareUser.UserID)
}

func TestMsgFile(t *testing.T) {
	mb := NewMsgBuffer(MsgFile)
	msg := mb.File("file_v2_71cafb2c-137f-4bb0-8381-ffd4971dbecg").Build()
	assert.Equal(t, MsgFile, msg.MsgType)
	assert.Equal(t, "file_v2_71cafb2c-137f-4bb0-8381-ffd4971dbecg", msg.Content.File.FileKey)
}

func TestMsgAudio(t *testing.T) {
	mb := NewMsgBuffer(MsgAudio)
	msg := mb.Audio("file_v2_71cafb2c-137f-4bb0-8381-ffd4971dbecg").Build()
	assert.Equal(t, MsgAudio, msg.MsgType)
	assert.Equal(t, "file_v2_71cafb2c-137f-4bb0-8381-ffd4971dbecg", msg.Content.Audio.FileKey)
}

func TestMsgMedia(t *testing.T) {
	mb := NewMsgBuffer(MsgMedia)
	msg := mb.Media("file_v2_b53cd6cc-5327-4968-8bf6-4528deb3068g", "img_v2_b276195a-9ae0-4fec-bbfe-f74b4d5a994g").Build()
	assert.Equal(t, MsgMedia, msg.MsgType)
	assert.Equal(t, "file_v2_b53cd6cc-5327-4968-8bf6-4528deb3068g", msg.Content.Media.FileKey)
	assert.Equal(t, "img_v2_b276195a-9ae0-4fec-bbfe-f74b4d5a994g", msg.Content.Media.ImageKey)
}

func TestMsgSticker(t *testing.T) {
	mb := NewMsgBuffer(MsgSticker)
	msg := mb.Sticker("4ba009df-2453-47b3-a753-444b152217bg").Build()
	assert.Equal(t, MsgSticker, msg.MsgType)
	assert.Equal(t, "4ba009df-2453-47b3-a753-444b152217bg", msg.Content.Sticker.FileKey)
}

func TestMsgWithWrongType(t *testing.T) {
	mb := NewMsgBuffer(MsgText)
	mb.ShareChat("6559399282837815565")
	assert.Equal(t, mb.Error().Error(), "`ShareChat` is only available to `share_chat`")
	mb.ShareUser("334455")
	assert.Equal(t, mb.Error().Error(), "`ShareUser` is only available to `share_user`")
	mb.Image("aaa")
	assert.Equal(t, mb.Error().Error(), "`Image` is only available to `image`")
	mb.File("aaa")
	assert.Equal(t, mb.Error().Error(), "`File` is only available to `file`")
	mb.Audio("aaa")
	assert.Equal(t, mb.Error().Error(), "`Audio` is only available to `audio`")
	mb.Media("aaa", "bbb")
	assert.Equal(t, mb.Error().Error(), "`Media` is only available to `media`")
	mb.Sticker("aaa")
	assert.Equal(t, mb.Error().Error(), "`Sticker` is only available to `sticker`")
	mb.Post(nil)
	assert.Equal(t, mb.Error().Error(), "`Post` is only available to `post`")
	mb.Card("nil")
	assert.Equal(t, mb.Error().Error(), "`Card` is only available to `interactive`")
	mbp := NewMsgBuffer(MsgPost)
	mbp.Text("hello")
	assert.Equal(t, mbp.Error().Error(), "`Text` is only available to `text`")
}

func TestClearMessage(t *testing.T) {
	mb := NewMsgBuffer(MsgText)
	mb.Text("hello, world").Build()
	assert.Equal(t, "hello, world", mb.message.Content.Text.Text)
	mb.Clear()
	assert.Equal(t, MsgText, mb.msgType)
	assert.Empty(t, mb.message.Content)
	mb.Text("attach again").Build()
	assert.Equal(t, "attach again", mb.message.Content.Text.Text)
}

func TestWorkWithTextBuilder(t *testing.T) {
	mb := NewMsgBuffer(MsgText)
	mb.Text(NewTextBuilder().Textln("hello, world").Render()).Build()
	assert.Equal(t, "hello, world\n", mb.message.Content.Text.Text)
}

func TestMsgUpdateMulti(t *testing.T) {
	mb := NewMsgBuffer(MsgInteractive)
	msg := mb.BindOpenChatID("6559399282837815565").UpdateMulti(true).Build()
	assert.Equal(t, MsgInteractive, msg.MsgType)
	assert.True(t, msg.UpdateMulti)
}

func TestWithSign(t *testing.T) {
	mb := NewMsgBuffer(MsgText)
	assert.Empty(t, mb.message.Sign)
	msg := mb.WithSign("xxx", 1661860880).Build()
	assert.NotEmpty(t, mb.message.Sign)
	assert.Equal(t, "QnWVTSBe6FmQDE0bG6X0mURbI+DnvVyu1h+j5dHOjrU=", msg.Sign)
}

func TestWithUUID(t *testing.T) {
	mb := NewMsgBuffer(MsgText)
	assert.Empty(t, mb.message.UUID)
	msg := mb.WithUUID("abc-def-0000").Build()
	assert.NotEmpty(t, mb.message.UUID)
	assert.Equal(t, "abc-def-0000", msg.UUID)
}
