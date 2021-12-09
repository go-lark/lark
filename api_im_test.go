package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostIMMessage(t *testing.T) {
	// text message
	msg := NewMsgBuffer(MsgText)
	om := msg.BindEmail(testUserEmail).Text("hello, world").Build()
	resp, err := bot.PostIMMessage(om)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
		t.Log(resp.Data)

		// get message
		getResp, err := bot.GetIMMessage(resp.Data.MessageID)
		t.Log(getResp)
		if assert.NoError(t, err) {
			assert.Equal(t, 0, getResp.Code)
			if assert.NotEmpty(t, getResp.Data.Items) {
				assert.Equal(t, getResp.Data.Items[0].MessageID, resp.Data.MessageID)
			}
		}
	}
}

func TestPostShareUser(t *testing.T) {
	// share user
	msg := NewMsgBuffer(MsgShareUser)
	om := msg.BindChatID(testGroupChatID).ShareUser(testUserOpenID).Build()
	resp, err := bot.PostIMMessage(om)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.Data.MessageID)
	}
}
