package lark

// OutcomingMessage struct of an outcoming message
type OutcomingMessage struct {
	MsgType MessageType `json:"msg_type"`
	// ID for user
	OpenID *string `json:"open_id,omitempty"`
	Email  *string `json:"email,omitempty"`
	UserID *string `json:"user_id,omitempty"`
	ChatID *string `json:"chat_id,omitempty"`
	// For reply
	RootID *string `json:"root_id,omitempty"`
	// Content
	Content MessageContent `json:"content"`
	Card    CardContent    `json:"card"`
	// UpdateMulti card
	UpdateMulti bool `json:"update_multi"`
}

// CardContent struct of card content
type CardContent *map[string]interface{}

// MessageContent struct of message content
type MessageContent struct {
	Text      *string      `json:"text,omitempty"`
	ImageKey  *string      `json:"image_key,omitempty"`
	ShareChat *string      `json:"share_open_chat_id,omitempty"`
	Post      *PostContent `json:"post,omitempty"`
}

// MessageType message type
type MessageType string

const (
	// MsgText simple text message
	MsgText MessageType = "text"
	// MsgPost rich text message
	MsgPost MessageType = "post"
	// MsgImage simple image message
	MsgImage MessageType = "image"
	// MsgShareCard share chat group card
	MsgShareCard MessageType = "share_chat"
	// MsgInteractive interactive widget
	MsgInteractive MessageType = "interactive"
)

// BuildOutcomingMessageReq for msg builder
func BuildOutcomingMessageReq(om OutcomingMessage) map[string]interface{} {
	params := map[string]interface{}{
		"msg_type":     om.MsgType,
		"root_id":      om.RootID,
		"chat_id":      om.ChatID,
		"open_id":      om.OpenID,
		"email":        om.Email,
		"user_id":      om.UserID,
		"update_multi": om.UpdateMulti,
	}
	content := make(map[string]interface{})
	if om.Content.Text != nil {
		content["text"] = *om.Content.Text
	}
	if om.Content.ImageKey != nil {
		content["image_key"] = *om.Content.ImageKey
	}
	if om.Content.ShareChat != nil {
		content["share_open_chat_id"] = *om.Content.ShareChat
	}
	if om.Content.Post != nil {
		content["post"] = *om.Content.Post
	}
	if om.MsgType == MsgInteractive && om.Card != nil {
		params["card"] = om.Card
	}
	params["content"] = content
	return params
}
