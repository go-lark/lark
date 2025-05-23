package lark

// EventBotDeleted .
type EventBotDeleted = EventBotAdded

// GetBotDeleted .
func (e Event) GetBotDeleted() (*EventBotDeleted, error) {
	var body EventBotDeleted
	err := e.GetEvent(EventTypeBotDeleted, &body)
	return &body, err
}
