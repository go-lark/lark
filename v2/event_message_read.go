package lark

// EventMessageRead .
type EventMessageRead struct {
	Reader struct {
		ReaderID  EventUserID `json:"reader_id,omitempty"`
		ReadTime  string      `json:"read_time,omitempty"`
		TenantKey string      `json:"tenant_key,omitempty"`
	} `json:"reader,omitempty"`
	MessageIDList []string `json:"message_id_list,omitempty"`
}

// GetMessageRead .
func (e Event) GetMessageRead() (*EventMessageRead, error) {
	var body EventMessageRead
	err := e.GetEvent(EventTypeMessageRead, &body)
	return &body, err
}
