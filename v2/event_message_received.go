package lark

// EventMessageReceived .
type EventMessageReceived struct {
	Sender struct {
		SenderID   EventUserID `json:"sender_id,omitempty"`
		SenderType string      `json:"sender_type,omitempty"`
		TenantKey  string      `json:"tenant_key,omitempty"`
	} `json:"sender,omitempty"`
	Message struct {
		MessageID   string `json:"message_id,omitempty"`
		RootID      string `json:"root_id,omitempty"`
		ParentID    string `json:"parent_id,omitempty"`
		CreateTime  string `json:"create_time,omitempty"`
		UpdateTime  string `json:"update_time,omitempty"`
		ChatID      string `json:"chat_id,omitempty"`
		ChatType    string `json:"chat_type,omitempty"`
		ThreadID    string `json:"thread_id,omitempty"`
		MessageType string `json:"message_type,omitempty"`
		Content     string `json:"content,omitempty"`
		Mentions    []struct {
			Key       string      `json:"key,omitempty"`
			ID        EventUserID `json:"id,omitempty"`
			Name      string      `json:"name,omitempty"`
			TenantKey string      `json:"tenant_key,omitempty"`
		} `json:"mentions,omitempty"`
		UserAgent string `json:"user_agent,omitempty"`
	} `json:"message,omitempty"`
}

// GetMessageReceived .
func (e Event) GetMessageReceived() (*EventMessageReceived, error) {
	var body EventMessageReceived
	err := e.GetEvent(EventTypeMessageReceived, &body)
	return &body, err
}
