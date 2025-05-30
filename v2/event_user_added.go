package lark

// EventUserAdded .
type EventUserAdded struct {
	ChatID            string      `json:"chat_id,omitempty"`
	OperatorID        EventUserID `json:"operator_id,omitempty"`
	External          bool        `json:"external,omitempty"`
	OperatorTenantKey string      `json:"operator_tenant_key,omitempty"`
	Users             []struct {
		Name      string      `json:"name,omitempty"`
		TenantKey string      `json:"tenant_key,omitempty"`
		UserID    EventUserID `json:"user_id,omitempty"`
	} `json:"users,omitempty"`
}

// GetUserAdded .
func (e Event) GetUserAdded() (*EventUserAdded, error) {
	var body EventUserAdded
	err := e.GetEvent(EventTypeUserAdded, &body)
	return &body, err
}
