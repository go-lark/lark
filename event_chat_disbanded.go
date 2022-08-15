package lark

import "encoding/json"

// EventV2ChatDisbanded .
type EventV2ChatDisbanded struct {
	ChatID            string        `json:"chat_id,omitempty"`
	OperatorID        EventV2UserID `json:"operator_id,omitempty"`
	External          bool          `json:"external,omitempty"`
	OperatorTenantKey string        `json:"operator_tenant_key,omitempty"`
}

// GetChatDisbanded .
func (e EventV2) GetChatDisbanded() (*EventV2ChatDisbanded, error) {
	if e.Header.EventType != EventTypeChatDisbanded {
		return nil, ErrEventTypeNotMatch
	}
	var body EventV2ChatDisbanded
	err := json.Unmarshal(e.EventRaw, &body)
	e.Event = body
	return &body, err
}
