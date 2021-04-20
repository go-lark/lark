package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAttachText(t *testing.T) {
	mb := NewMsgBuffer(MsgText)
	msg := mb.Text("hello").Build()
	assert.Equal(t, "hello", *msg.Content.Text)
}

func TestAttachImage(t *testing.T) {
	mb := NewMsgBuffer(MsgImage)
	msg := mb.Image("aaaaa").Build()
	assert.Equal(t, "aaaaa", *msg.Content.ImageKey)
}

func TestMsgTextBinding(t *testing.T) {
	mb := NewMsgBuffer(MsgText)
	msg := mb.Text("hello, world").BindEmail(testUserEmail).Build()
	assert.Equal(t, "hello, world", *msg.Content.Text)
	assert.Equal(t, testUserEmail, *msg.Email)
}

func TestBindingUserIDs(t *testing.T) {
	mb := NewMsgBuffer(MsgText)
	msgEmail := mb.BindEmail(testUserEmail).Build()
	assert.Equal(t, testUserEmail, *msgEmail.Email)

	mb.Clear()
	msgOpenChatID := mb.BindOpenChatID(testGroupChatID).Build()
	assert.Equal(t, testGroupChatID, *msgOpenChatID.ChatID)

	mb.Clear()
	msgUserID := mb.BindUserID("333444").Build()
	assert.Equal(t, "333444", *msgUserID.UserID)

	mb.Clear()
	msgReplyID := mb.BindReply(testMessageID).Build()
	assert.Equal(t, testMessageID, *msgReplyID.RootID)
}

func TestMsgShareChat(t *testing.T) {
	mb := NewMsgBuffer(MsgShareCard)
	msg := mb.ShareChat("6559399282837815565").Build()
	assert.Equal(t, MsgShareCard, msg.MsgType)
	assert.Equal(t, "6559399282837815565", *msg.Content.ShareChat)
}

func TestMsgWithWrongType(t *testing.T) {
	// with wrong type
	mb := NewMsgBuffer(MsgText)
	output := captureOutput(func() {
		mb.ShareChat("6559399282837815565")
	})
	assert.Contains(t, output, "`ShareChat` is only available to MsgShareChat")
	output = captureOutput(func() {
		mb.Image("aaa")
	})
	assert.Contains(t, output, "`Image` is only available to MsgImage")
	output = captureOutput(func() {
		mb.Post(nil)
	})
	assert.Contains(t, output, "`Post` is only available to MsgPost")
	output = captureOutput(func() {
		mb.Card("nil")
	})
	assert.Contains(t, output, "`Card` is only available to MsgInteractive")

	mbp := NewMsgBuffer(MsgPost)
	output = captureOutput(func() {
		mbp.Text("hello")
	})
	assert.Contains(t, output, "`Text` is only available to MsgText")
}

func TestClearMessage(t *testing.T) {
	mb := NewMsgBuffer(MsgText)
	mb.Text("hello, world").Build()
	assert.Equal(t, "hello, world", *mb.message.Content.Text)
	mb.Clear()
	assert.Equal(t, MsgText, mb.msgType)
	assert.Empty(t, mb.message.Content)
	mb.Text("attach again").Build()
	assert.Equal(t, "attach again", *mb.message.Content.Text)
}

func TestWorkWithTextBuilder(t *testing.T) {
	mb := NewMsgBuffer(MsgText)
	mb.Text(NewTextBuilder().Textln("hello, world").Render()).Build()
	assert.Equal(t, "hello, world\n", *mb.message.Content.Text)
}

func TestMsgUpdateMulti(t *testing.T) {
	mb := NewMsgBuffer(MsgInteractive)
	msg := mb.BindOpenChatID("6559399282837815565").UpdateMulti(true).Build()
	assert.Equal(t, MsgInteractive, msg.MsgType)
	assert.True(t, msg.UpdateMulti)
}
