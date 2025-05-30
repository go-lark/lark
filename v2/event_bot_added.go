package lark

// EventBotAdded .
type EventBotAdded struct {
	ChatID            string      `json:"chat_id,omitempty"`
	OperatorID        EventUserID `json:"operator_id,omitempty"`
	External          bool        `json:"external,omitempty"`
	OperatorTenantKey string      `json:"operator_tenant_key,omitempty"`
}

// GetBotAdded .
func (e Event) GetBotAdded() (*EventBotAdded, error) {
	var body EventBotAdded
	err := e.GetEvent(EventTypeBotAdded, &body)
	return &body, err
}
