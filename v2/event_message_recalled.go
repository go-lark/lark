package lark

// EventMessageRecalled .
type EventMessageRecalled struct {
	MessageID  string `json:"message_id,omitempty"`
	ChatID     string `json:"chat_id,omitempty"`
	RecallTime string `json:"recall_time,omitempty"`
	RecallType string `json:"recall_type,omitempty"`
}

// GetMessageRecalled .
func (e Event) GetMessageRecalled() (*EventMessageRecalled, error) {
	var body EventMessageRecalled
	err := e.GetEvent(EventTypeMessageRecalled, &body)
	return &body, err
}
