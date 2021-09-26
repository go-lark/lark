package lark

import (
	"encoding/json"
	"fmt"
)

// MsgBuffer stores all the messages attached
// You can call every function, but some of which is only available for specific condition
type MsgBuffer struct {
	// Message type
	msgType MessageType
	// Output
	message OutcomingMessage

	err error
}

// NewMsgBuffer create a message buffer
func NewMsgBuffer(newMsgType MessageType) *MsgBuffer {
	msgBuffer := MsgBuffer{
		message: OutcomingMessage{
			MsgType: newMsgType,
		},
		msgType: newMsgType,
	}
	return &msgBuffer
}

// BindOpenID binds open_id
func (m *MsgBuffer) BindOpenID(openID string) *MsgBuffer {
	m.message.OpenID = &openID
	return m
}

// BindEmail binds email
func (m *MsgBuffer) BindEmail(email string) *MsgBuffer {
	m.message.Email = &email
	return m
}

// BindChatID binds chat_id
func (m *MsgBuffer) BindChatID(chatID string) *MsgBuffer {
	m.message.ChatID = &chatID
	return m
}

// BindOpenChatID binds open_chat_id
func (m *MsgBuffer) BindOpenChatID(openChatID string) *MsgBuffer {
	m.BindChatID(openChatID)
	return m
}

// BindUserID binds open_id
func (m *MsgBuffer) BindUserID(userID string) *MsgBuffer {
	m.message.UserID = &userID
	return m
}

// BindReply binds root id for reply
// rootID is OpenMessageID of the message you reply
func (m *MsgBuffer) BindReply(rootID string) *MsgBuffer {
	m.message.RootID = &rootID
	return m
}

// UpdateMulti set multi, shared card
// default false, not share
func (m *MsgBuffer) UpdateMulti(flag bool) *MsgBuffer {
	m.message.UpdateMulti = flag
	return m
}

func (m MsgBuffer) typeError(funcName string, msgType MessageType) error {
	return fmt.Errorf("`%s` is only available to `%s`", funcName, msgType)
}

// Text attaches text
func (m *MsgBuffer) Text(text string) *MsgBuffer {
	if m.msgType != MsgText {
		m.err = m.typeError("Text", MsgText)
		return m
	}
	m.message.Content.Text = &text
	return m
}

// Image attaches image key
// for MsgImage only
func (m *MsgBuffer) Image(imageKey string) *MsgBuffer {
	if m.msgType != MsgImage {
		m.err = m.typeError("Image", MsgImage)
		return m
	}
	m.message.Content.ImageKey = &imageKey
	return m
}

// ShareChat attaches chat id
// for MsgShareChat only
func (m *MsgBuffer) ShareChat(chatID string) *MsgBuffer {
	if m.msgType != MsgShareCard {
		m.err = m.typeError("ShareChat", MsgShareCard)
		return m
	}
	m.message.Content.ShareChat = &chatID
	return m
}

// Post sets raw post content
func (m *MsgBuffer) Post(postContent *PostContent) *MsgBuffer {
	if m.msgType != MsgPost {
		m.err = m.typeError("Post", MsgPost)
		return m
	}
	m.message.Content.Post = postContent
	return m
}

// Card binds card content with V4 format
func (m *MsgBuffer) Card(cardContent string) *MsgBuffer {
	if m.msgType != MsgInteractive {
		m.err = m.typeError("Card", MsgInteractive)
		return m
	}
	m.message.Card = new(map[string]interface{})
	_ = json.Unmarshal([]byte(cardContent), m.message.Card)
	return m
}

// Build message and return message body
func (m *MsgBuffer) Build() OutcomingMessage {
	return m.message
}

// Error returns last error
func (m *MsgBuffer) Error() error {
	return m.err
}

// Clear message in buffer
func (m *MsgBuffer) Clear() *MsgBuffer {
	m.message = OutcomingMessage{
		MsgType: m.msgType,
	}
	m.err = nil
	return m
}
