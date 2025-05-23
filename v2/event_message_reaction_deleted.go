package lark

// EventMessageReactionDeleted .
type EventMessageReactionDeleted struct {
	MessageID    string      `json:"message_id,omitempty"`
	OperatorType string      `json:"operator_type,omitempty"`
	UserID       EventUserID `json:"user_id,omitempty"`
	AppID        string      `json:"app_id,omitempty"`
	ActionTime   string      `json:"action_time,omitempty"`
	ReactionType struct {
		EmojiType string `json:"emoji_type,omitempty"`
	} `json:"reaction_type,omitempty"`
}

// GetMessageReactionDeleted .
func (e Event) GetMessageReactionDeleted() (*EventMessageReactionDeleted, error) {
	var body EventMessageReactionDeleted
	err := e.GetEvent(EventTypeMessageReactionDeleted, &body)
	return &body, err
}
