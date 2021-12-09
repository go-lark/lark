package lark

const (
	messageURL                = "/open-apis/message/v4/send/"
	ephemeralMessageURL       = "/open-apis/ephemeral/v1/send"
	deleteEphemeralMessageURL = "/open-apis/ephemeral/v1/delete"
	recallMessageURL          = "/open-apis/message/v4/recall/"
	messageReceiptURL         = "/open-apis/message/v4/read_info/"
)

// PostMessageResponse .
type PostMessageResponse struct {
	BaseResponse
	Data struct {
		MessageID string `json:"message_id"`
	} `json:"data"`
}

// PostEphemeralMessageResponse .
type PostEphemeralMessageResponse = PostMessageResponse

// DeleteEphemeralMessageResponse .
type DeleteEphemeralMessageResponse = BaseResponse

// RecallMessageResponse .
type RecallMessageResponse = BaseResponse

// MessageReceiptResponse .
type MessageReceiptResponse struct {
	BaseResponse
	Data struct {
		ReadUsers []struct {
			OpenID    string `json:"open_id"`
			UserID    string `json:"user_id"`
			Timestamp string `json:"timestamp"`
		} `json:"read_users"`
	} `json:"data"`
}

func newMsgBufWithOptionalUserID(msgType string, userID *OptionalUserID) *MsgBuffer {
	mb := NewMsgBuffer(msgType)
	realID := userID.RealID
	switch userID.IDType {
	case "email":
		mb.BindEmail(realID)
	case "open_id":
		mb.BindOpenID(realID)
	case "chat_id":
		mb.BindChatID(realID)
	case "user_id":
		mb.BindUserID(realID)
	default:
		return nil
	}
	return mb
}

// PostText is a simple way to send text messages
func (bot *Bot) PostText(text string, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgText, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	om := mb.Text(text).Build()
	return bot.PostMessage(om)
}

// PostRichText is a simple way to send rich text messages
func (bot *Bot) PostRichText(postContent *PostContent, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgPost, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	om := mb.Post(postContent).Build()
	return bot.PostMessage(om)
}

// PostTextMention is a simple way to send text messages with @user
func (bot *Bot) PostTextMention(text string, atUserID string, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgText, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	tb := NewTextBuilder()
	om := mb.Text(tb.Text(text).Mention(atUserID).Render()).Build()
	return bot.PostMessage(om)
}

// PostTextMentionAndReply is a simple way to send text messages with @user and reply a message
func (bot *Bot) PostTextMentionAndReply(text string, atUserID string, userID *OptionalUserID, replyID string) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgText, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	tb := NewTextBuilder()
	om := mb.Text(tb.Text(text).Mention(atUserID).Render()).BindReply(replyID).Build()
	return bot.PostMessage(om)
}

// PostTextMentionAll is a simple way to send text messages with @all
func (bot *Bot) PostTextMentionAll(text string, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgText, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	tb := NewTextBuilder()
	om := mb.Text(tb.Text(text).MentionAll().Render()).Build()
	return bot.PostMessage(om)
}

// PostImage is a simple way to send image
func (bot *Bot) PostImage(imageKey string, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgImage, userID)
	om := mb.Image(imageKey).Build()
	return bot.PostMessage(om)
}

// PostShareChat is a simple way to share chat
func (bot *Bot) PostShareChat(openChatID string, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgShareCard, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	om := mb.ShareChat(openChatID).Build()
	return bot.PostMessage(om)
}

// PostMessage posts message
func (bot *Bot) PostMessage(om OutcomingMessage) (*PostMessageResponse, error) {
	if om.UIDType == UIDUnionID {
		return nil, ErrUnsupportedUIDType
	}
	params := BuildOutcomingMessageReq(om)
	var respData PostMessageResponse
	err := bot.PostAPIRequest("PostMessage", messageURL, true, params, &respData)
	return &respData, err
}

// PostEphemeralMessage posts an ephemeral message
func (bot *Bot) PostEphemeralMessage(om OutcomingMessage) (*PostEphemeralMessageResponse, error) {
	if om.UIDType == UIDUnionID {
		return nil, ErrUnsupportedUIDType
	}
	params := BuildOutcomingMessageReq(om)
	var respData PostEphemeralMessageResponse
	err := bot.PostAPIRequest("PostEphemeralMessage", ephemeralMessageURL, true, params, &respData)
	return &respData, err
}

// DeleteEphemeralMessage deletes an ephemeral message
func (bot *Bot) DeleteEphemeralMessage(messageID string) (*DeleteEphemeralMessageResponse, error) {
	params := map[string]interface{}{
		"message_id": messageID,
	}
	var respData DeleteEphemeralMessageResponse
	err := bot.PostAPIRequest("DeleteEphemeralMessage", deleteEphemeralMessageURL, true, params, &respData)
	return &respData, err
}

// RecallMessage recalls a message with ID
func (bot *Bot) RecallMessage(messageID string) (*RecallMessageResponse, error) {
	params := map[string]interface{}{
		"message_id": messageID,
	}
	var respData RecallMessageResponse
	err := bot.PostAPIRequest("RecallMessage", recallMessageURL, true, params, &respData)
	return &respData, err
}

// MessageReadReceipt queries message read receipt
func (bot *Bot) MessageReadReceipt(messageID string) (*MessageReceiptResponse, error) {
	params := map[string]interface{}{
		"message_id": messageID,
	}
	var respData MessageReceiptResponse
	err := bot.PostAPIRequest("MessageReadReceipt", messageReceiptURL, true, params, &respData)
	return &respData, err
}
