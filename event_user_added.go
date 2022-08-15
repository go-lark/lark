package lark

import "encoding/json"

// EventV2UserAdded .
type EventV2UserAdded struct {
	ChatID            string        `json:"chat_id,omitempty"`
	OperatorID        EventV2UserID `json:"operator_id,omitempty"`
	External          bool          `json:"external,omitempty"`
	OperatorTenantKey string        `json:"operator_tenant_key,omitempty"`
	Users             []struct {
		Name      string        `json:"name,omitempty"`
		TenantKey string        `json:"tenant_key,omitempty"`
		UserID    EventV2UserID `json:"user_id,omitempty"`
	} `json:"users,omitempty"`
}

// GetUserAdded .
func (e EventV2) GetUserAdded() (*EventV2UserAdded, error) {
	if e.Header.EventType != EventTypeUserAdded {
		return nil, ErrEventTypeNotMatch
	}
	var body EventV2UserAdded
	err := json.Unmarshal(e.EventRaw, &body)
	e.Event = body
	return &body, err
}
