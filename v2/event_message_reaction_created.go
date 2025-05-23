package lark

// EventMessageReactionCreated .
type EventMessageReactionCreated struct {
	MessageID    string      `json:"message_id,omitempty"`
	OperatorType string      `json:"operator_type,omitempty"`
	UserID       EventUserID `json:"user_id,omitempty"`
	AppID        string      `json:"app_id,omitempty"`
	ActionTime   string      `json:"action_time,omitempty"`
	ReactionType struct {
		EmojiType string `json:"emoji_type,omitempty"`
	} `json:"reaction_type,omitempty"`
}

// GetMessageReactionCreated .
func (e Event) GetMessageReactionCreated() (*EventMessageReactionCreated, error) {
	var body EventMessageReactionCreated
	err := e.GetEvent(EventTypeMessageReactionCreated, &body)
	return &body, err
}
