package lark

// EventChatDisbanded .
type EventChatDisbanded struct {
	ChatID            string      `json:"chat_id,omitempty"`
	OperatorID        EventUserID `json:"operator_id,omitempty"`
	External          bool        `json:"external,omitempty"`
	OperatorTenantKey string      `json:"operator_tenant_key,omitempty"`
}

// GetChatDisbanded .
func (e Event) GetChatDisbanded() (*EventChatDisbanded, error) {
	var body EventChatDisbanded
	err := e.GetEvent(EventTypeChatDisbanded, &body)
	return &body, err
}
