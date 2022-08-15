package lark

import "encoding/json"

// EventV2BotDeleted .
type EventV2BotDeleted = EventV2BotAdded

// GetBotDeleted .
func (e EventV2) GetBotDeleted() (*EventV2BotDeleted, error) {
	if e.Header.EventType != EventTypeBotDeleted {
		return nil, ErrEventTypeNotMatch
	}
	var body EventV2BotDeleted
	err := json.Unmarshal(e.EventRaw, &body)
	e.Event = body
	return &body, err
}
