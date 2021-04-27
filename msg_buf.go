package lark

import (
	"encoding/json"
	"log"
)

// MsgBuffer stores all the messages attached
// You can call every function, but some of which is only available for specific condition
type MsgBuffer struct {
	// Message type
	msgType MessageType
	// Output
	message OutcomingMessage
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

// Text attaches text
func (m *MsgBuffer) Text(text string) *MsgBuffer {
	if m.msgType != MsgText {
		log.Println("`Text` is only available to MsgText")
	}
	m.message.Content.Text = &text
	return m
}

// Image attaches image key
// for MsgImage only
func (m *MsgBuffer) Image(imageKey string) *MsgBuffer {
	if m.msgType != MsgImage {
		log.Println("`Image` is only available to MsgImage")
	}
	m.message.Content.ImageKey = &imageKey
	return m
}

// ShareChat attaches chat id
// for MsgShareChat only
func (m *MsgBuffer) ShareChat(chatID string) *MsgBuffer {
	if m.msgType != MsgShareCard {
		log.Println("`ShareChat` is only available to MsgShareChat")
	}
	m.message.Content.ShareChat = &chatID
	return m
}

// Post sets raw post content
func (m *MsgBuffer) Post(postContent *PostContent) *MsgBuffer {
	if m.msgType != MsgPost {
		log.Println("`Post` is only available to MsgPost")
	}
	m.message.Content.Post = postContent
	return m
}

// Card binds card content with V4 format
func (m *MsgBuffer) Card(cardContent string) *MsgBuffer {
	if m.msgType != MsgInteractive {
		log.Println("`Card` is only available to MsgInteractive")
	}
	m.message.Card = new(map[string]interface{})
	_ = json.Unmarshal([]byte(cardContent), m.message.Card)
	return m
}

// Build message and return message body
func (m *MsgBuffer) Build() OutcomingMessage {
	return m.message
}

// Clear message in buffer
func (m *MsgBuffer) Clear() *MsgBuffer {
	m.message = OutcomingMessage{
		MsgType: m.msgType,
	}
	return m
}
