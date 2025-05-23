package lark

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// EventType definitions
const (
	EventTypeMessageReceived        = "im.message.receive_v1"
	EventTypeMessageRead            = "im.message.message_read_v1"
	EventTypeMessageRecalled        = "im.message.recalled_v1"
	EventTypeMessageReactionCreated = "im.message.reaction.created_v1"
	EventTypeMessageReactionDeleted = "im.message.reaction.deleted_v1"
	EventTypeChatDisbanded          = "im.chat.disbanded_v1"
	EventTypeUserAdded              = "im.chat.member.user.added_v1"
	EventTypeUserDeleted            = "im.chat.member.user.deleted_v1"
	EventTypeBotAdded               = "im.chat.member.bot.added_v1"
	EventTypeBotDeleted             = "im.chat.member.bot.deleted_v1"
	// not supported yet
	EventTypeChatUpdated   = "im.chat.updated_v1"
	EventTypeUserWithdrawn = "im.chat.member.user.withdrawn_v1"
)

// Event handles events with v2 schema
type Event struct {
	Schema string      `json:"schema,omitempty"`
	Header EventHeader `json:"header,omitempty"`

	EventRaw json.RawMessage `json:"event,omitempty"`
	Event    interface{}     `json:"-"`
}

// EventHeader .
type EventHeader struct {
	EventID    string `json:"event_id,omitempty"`
	EventType  string `json:"event_type,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	Token      string `json:"token,omitempty"`
	AppID      string `json:"app_id,omitempty"`
	TenantKey  string `json:"tenant_key,omitempty"`
}

// EventBody .
type EventBody struct {
	Type          string `json:"type"`
	AppID         string `json:"app_id"`
	TenantKey     string `json:"tenant_key"`
	ChatType      string `json:"chat_type"`
	MsgType       string `json:"msg_type"`
	RootID        string `json:"root_id,omitempty"`
	ParentID      string `json:"parent_id,omitempty"`
	OpenID        string `json:"open_id,omitempty"`
	OpenChatID    string `json:"open_chat_id,omitempty"`
	OpenMessageID string `json:"open_message_id,omitempty"`
	IsMention     bool   `json:"is_mention,omitempty"`
	Title         string `json:"title,omitempty"`
	Text          string `json:"text,omitempty"`
	RealText      string `json:"text_without_at_bot,omitempty"`
	ImageKey      string `json:"image_key,omitempty"`
	ImageURL      string `json:"image_url,omitempty"`
	FileKey       string `json:"file_key,omitempty"`
}

// EventUserID .
type EventUserID struct {
	UnionID string `json:"union_id,omitempty"`
	UserID  string `json:"user_id,omitempty"`
	OpenID  string `json:"open_id,omitempty"`
}

// PostEvent posts events
func (e Event) PostEvent(client *http.Client, hookURL string) (*http.Response, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(e)
	if err != nil {
		log.Printf("Encode json failed: %+v\n", err)
		return nil, err
	}
	resp, err := client.Post(hookURL, "application/json; charset=utf-8", buf)
	return resp, err
}

// GetEvent .
func (e Event) GetEvent(eventType string, body interface{}) error {
	if e.Header.EventType != eventType {
		return ErrEventTypeNotMatch
	}
	err := json.Unmarshal(e.EventRaw, &body)
	return err
}
