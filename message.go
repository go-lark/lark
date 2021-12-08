package lark

// Msg Types
const (
	MsgText        = "text"
	MsgPost        = "post"
	MsgInteractive = "interactive"
	MsgImage       = "image"
	MsgShareCard   = "share_chat"
	MsgShareUser   = "share_user"
	MsgAudio       = "audio"
	MsgMedia       = "media"
	MsgFile        = "file"
	MsgSticker     = "sticker"
)

// OutcomingMessage struct of an outcoming message
type OutcomingMessage struct {
	MsgType string
	Content MessageContent
	// ID for user
	UIDType string
	OpenID  string
	Email   string
	UserID  string
	ChatID  string
	UnionID string
	// For reply
	RootID string
	// UpdateMulti card
	UpdateMulti bool
}

// CardContent struct of card content
type CardContent map[string]interface{}

// MessageContent struct of message content
type MessageContent struct {
	Text      *TextContent
	Image     *ImageContent
	Post      *PostContent
	Card      *CardContent
	ShareChat *ShareChatContent
	ShareUser *ShareUserContent
	Audio     *AudioContent
	Media     *MediaContent
	File      *FileContent
	Sticker   *StickerContent
}

// TextContent .
type TextContent struct {
	Text string `json:"text"`
}

// ImageContent .
type ImageContent struct {
	ImageKey string `json:"image_key"`
}

// ShareChatContent .
type ShareChatContent struct {
	ChatID string `json:"chat_id"`
}

// ShareUserContent .
type ShareUserContent struct {
	UserID string `json:"user_id"`
}

// AudioContent .
type AudioContent struct {
	FileKey string `json:"file_key"`
}

// MediaContent .
type MediaContent struct {
	FileKey  string `json:"file_key"`
	ImageKey string `json:"image_key"`
}

// FileContent .
type FileContent struct {
	FileKey string `json:"file_key"`
}

// StickerContent .
type StickerContent struct {
	FileKey string `json:"file_key"`
}
