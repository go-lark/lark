package lark

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

// See https://open.feishu.cn/document/uYjL24iN/uUTNz4SN1MjL1UzM
const (
	// EventGetMessage .
	EventGetMessage = 1
)

// EventMessage .
type EventMessage struct {
	UUID      string `json:"uuid"`
	Timestamp string `json:"ts"`
	// Token is shown by Lark to indicate it is not a fake message, check at your own need
	Token     string    `json:"token"`
	EventType string    `json:"type"`
	Event     EventBody `json:"event"`
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

// EventChallengeReq request of add event hook
type EventChallengeReq struct {
	Token     string `json:"token,omitempty"`
	Challenge string `json:"challenge,omitempty"`
	Type      string `json:"type,omitempty"`
}

// EncryptedReq request of encrypted challagen
type EncryptedReq struct {
	Encrypt string `json:"encrypt,omitempty"`
}

// PostEvent posts event
// 1. help to develop and test ServeEvent callback func much easier
// 2. otherwise, you may use it to forward event
func PostEvent(client *http.Client, hookURL string, message EventMessage) (*http.Response, error) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(message)
	if err != nil {
		log.Printf("Encode json failed: %+v\n", err)
		return nil, err
	}
	resp, err := client.Post(hookURL, "application/json; charset=utf-8", buf)
	return resp, err
}
