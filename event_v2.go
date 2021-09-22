package lark

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// EventV2 handles events with v2 schema
type EventV2 struct {
	Schema string        `json:"schema,omitempty"`
	Header EventV2Header `json:"header,omitempty"`
	Event  interface{}   `json:"event,omitempty"`
}

// GetMessageReceived .
func (e EventV2) GetMessageReceived() (EventV2MessageReceived, bool) {
	body, ok := e.Event.(EventV2MessageReceived)
	return body, ok
}

// GetChatDisbanded .
func (e EventV2) GetChatDisbanded() (EventV2ChatDisbanded, bool) {
	body, ok := e.Event.(EventV2ChatDisbanded)
	return body, ok
}

// EventType definitions
const (
	EventTypeMessageReceived = "im.message.receive_v1"
	EventTypeChatDisbanded   = "im.chat.disbanded_v1"
	// not supported yet
	EventTypeMessageRead   = "im.message.message_read_v1"
	EventTypeChatUpdated   = "im.chat.updated_v1"
	EventTypeBotAdded      = "im.chat.member.bot.added_v1"
	EventTypeBotDeleted    = "im.chat.member.bot.deleted_v1"
	EventTypeUserAdded     = "im.chat.member.user.added_v1"
	EventTypeUserWithdrawn = "im.chat.member.user.withdrawn_v1"
	EventTypeUserDeleted   = "im.chat.member.user.deleted_v1"
)

// PostEvent with event v2 format
// and it's part of EventV2 instead of package method
func (e EventV2) PostEvent(client *http.Client, hookURL string, event EventV2) (*http.Response, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(event)
	if err != nil {
		log.Printf("Encode json failed: %+v\n", err)
		return nil, err
	}
	resp, err := client.Post(hookURL, "application/json; charset=utf-8", buf)
	return resp, err
}

// EventV2Header .
type EventV2Header struct {
	EventID    string `json:"event_id,omitempty"`
	EventType  string `json:"event_type,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
	Token      string `json:"token,omitempty"`
	AppID      string `json:"app_id,omitempty"`
	TenantKey  string `json:"tenant_key,omitempty"`
}

// EventV2UserID .
type EventV2UserID struct {
	UnionID string `json:"union_id,omitempty"`
	UserID  string `json:"user_id,omitempty"`
	OpenID  string `json:"open_id,omitempty"`
}

// EventV2MessageReceived .
type EventV2MessageReceived struct {
	Sender struct {
		SenderID   EventV2UserID `json:"sender_id,omitempty"`
		SenderType string        `json:"sender_type,omitempty"`
		TenantKey  string        `json:"tenant_key,omitempty"`
	} `json:"sender,omitempty"`
	Message struct {
		MessageID   string `json:"message_id,omitempty"`
		RootID      string `json:"root_id,omitempty"`
		ParentID    string `json:"parent_id,omitempty"`
		CreateTime  string `json:"create_time,omitempty"`
		ChatID      string `json:"chat_id,omitempty"`
		ChatType    string `json:"chat_type,omitempty"`
		MessageType string `json:"message_type,omitempty"`
		Content     string `json:"content,omitempty"`
		Mentions    []struct {
			Key       string        `json:"key,omitempty"`
			ID        EventV2UserID `json:"id,omitempty"`
			Name      string        `json:"name,omitempty"`
			TenantKey string        `json:"tenant_key,omitempty"`
		} `json:"mentions,omitempty"`
	} `json:"message,omitempty"`
}

// EventV2ChatDisbanded .
type EventV2ChatDisbanded struct {
	ChatID            string        `json:"chat_id,omitempty"`
	OperatorID        EventV2UserID `json:"operator_id,omitempty"`
	External          bool          `json:"external,omitempty"`
	OperatorTenantKey string        `json:"operator_tenant_key,omitempty"`
}
