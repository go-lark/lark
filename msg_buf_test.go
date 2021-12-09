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

func TestMsgWithWrongType(t *testing.T) {
	mb := NewMsgBuffer(MsgText)
	mb.ShareChat("6559399282837815565")
	assert.Equal(t, mb.Error().Error(), "`ShareChat` is only available to `share_chat`")
	mb.ShareUser("334455")
	assert.Equal(t, mb.Error().Error(), "`ShareUser` is only available to `share_user`")
	mb.Image("aaa")
	assert.Equal(t, mb.Error().Error(), "`Image` is only available to `image`")
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
