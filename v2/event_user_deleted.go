package lark

// EventUserDeleted .
type EventUserDeleted = EventUserAdded

// GetUserDeleted .
func (e Event) GetUserDeleted() (*EventUserDeleted, error) {
	var body EventUserDeleted
	err := e.GetEvent(EventTypeUserDeleted, &body)
	return &body, err
}
