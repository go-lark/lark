package lark

import "encoding/json"

// EventV2UserDeleted .
type EventV2UserDeleted = EventV2UserAdded

// GetUserDeleted .
func (e EventV2) GetUserDeleted() (*EventV2UserDeleted, error) {
	if e.Header.EventType != EventTypeUserDeleted {
		return nil, ErrEventTypeNotMatch
	}
	var body EventV2UserDeleted
	err := json.Unmarshal(e.EventRaw, &body)
	e.Event = body
	return &body, err
}
