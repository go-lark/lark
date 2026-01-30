package lark

import (
	"context"
	"fmt"
)

const (
	messageURL                = "/open-apis/im/v1/messages?receive_id_type=%s"
	replyMessageURL           = "/open-apis/im/v1/messages/%s/reply"
	reactionsMessageURL       = "/open-apis/im/v1/messages/%s/reactions"
	deleteReactionsMessageURL = "/open-apis/im/v1/messages/%s/reactions/%s"
	getMessageURL             = "/open-apis/im/v1/messages/%s"
	updateMessageURL          = "/open-apis/im/v1/messages/%s"
	recallMessageURL          = "/open-apis/im/v1/messages/%s"
	messageReceiptURL         = "/open-apis/im/v1/messages/%s/read_users?user_id_type=%s&page_size=%d&page_token=%s"
	ephemeralMessageURL       = "/open-apis/ephemeral/v1/send"
	deleteEphemeralMessageURL = "/open-apis/ephemeral/v1/delete"
	pinMessageURL             = "/open-apis/im/v1/pins"
	unpinMessageURL           = "/open-apis/im/v1/pins/%s"
	forwardMessageURL         = "/open-apis/im/v1/messages/%s/forward?receive_id_type=%s"
)

// PostMessageResponse .
type PostMessageResponse struct {
	BaseResponse

	Data IMMessage `json:"data"`
}

// IMMessageRequest .
type IMMessageRequest struct {
	Content       string `json:"content"`
	MsgType       string `json:"msg_type,omitempty"`
	ReceiveID     string `json:"receive_id,omitempty"`
	UUID          string `json:"uuid,omitempty"`
	ReplyInThread bool   `json:"reply_in_thread,omitempty"`
}

// IMSender .
type IMSender struct {
	ID         string `json:"id"`
	IDType     string `json:"id_type"`
	SenderType string `json:"sender_type"`
	TenantKey  string `json:"tenant_key"`
}

// IMMention .
type IMMention struct {
	ID     string `json:"id"`
	IDType string `json:"id_type"`
	Key    string `json:"key"`
	Name   string `json:"name"`
}

// IMBody .
type IMBody struct {
	Content string `json:"content"`
}

// IMMessage .
type IMMessage struct {
	MessageID      string      `json:"message_id"`
	UpperMessageID string      `json:"upper_message_id"`
	RootID         string      `json:"root_id"`
	ParentID       string      `json:"parent_id"`
	ThreadID       string      `json:"thread_id"`
	ChatID         string      `json:"chat_id"`
	MsgType        string      `json:"msg_type"`
	CreateTime     string      `json:"create_time"`
	UpdateTime     string      `json:"update_time"`
	Deleted        bool        `json:"deleted"`
	Updated        bool        `json:"updated"`
	Sender         IMSender    `json:"sender"`
	Mentions       []IMMention `json:"mentions"`
	Body           IMBody      `json:"body"`
}

// ReactionResponse .
type ReactionResponse struct {
	BaseResponse
	Data struct {
		ReactionID string `json:"reaction_id"`
		Operator   struct {
			OperatorID   string `json:"operator_id"`
			OperatorType string `json:"operator_type"`
			ActionTime   string `json:"action_time"`
		} `json:"operator"`
		ReactionType struct {
			EmojiType EmojiType `json:"emoji_type"`
		} `json:"reaction_type"`
	} `json:"data"`
}

// GetMessageResponse .
type GetMessageResponse struct {
	BaseResponse

	Data struct {
		Items []IMMessage `json:"items"`
	} `json:"data"`
}

// PostEphemeralMessageResponse .
type PostEphemeralMessageResponse struct {
	BaseResponse
	Data struct {
		MessageID string `json:"message_id"`
	} `json:"data"`
}

// DeleteEphemeralMessageResponse .
type DeleteEphemeralMessageResponse = BaseResponse

// RecallMessageResponse .
type RecallMessageResponse = BaseResponse

// UpdateMessageResponse .
type UpdateMessageResponse = BaseResponse

// ForwardMessageResponse .
type ForwardMessageResponse = PostMessageResponse

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

// PinMessageResponse .
type PinMessageResponse struct {
	BaseResponse
	Data struct {
		Pin struct {
			MessageID      string `json:"message_id"`
			ChatID         string `json:"chat_id"`
			OperatorID     string `json:"operator_id"`
			OperatorIDType string `json:"operator_id_type"`
			CreateTime     string `json:"create_time"`
		} `json:"pin"`
	} `json:"data"`
}

// UnpinMessageResponse .
type UnpinMessageResponse = BaseResponse

func newMsgBufWithOptionalUserID(msgType string, userID *OptionalUserID) *MsgBuffer {
	mb := NewMsgBuffer(msgType)
	realID := userID.RealID
	switch userID.UIDType {
	case "email":
		mb.BindEmail(realID)
	case "open_id":
		mb.BindOpenID(realID)
	case "chat_id":
		mb.BindChatID(realID)
	case "user_id":
		mb.BindUserID(realID)
	case "union_id":
		mb.BindUnionID(realID)
	default:
		return nil
	}
	return mb
}

// PostText is a simple way to send text messages
func (bot Bot) PostText(ctx context.Context, text string, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgText, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	om := mb.Text(text).Build()
	return bot.PostMessage(ctx, om)
}

// PostRichText is a simple way to send rich text messages
func (bot Bot) PostRichText(ctx context.Context, postContent *PostContent, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgPost, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	om := mb.Post(postContent).Build()
	return bot.PostMessage(ctx, om)
}

// PostTextMention is a simple way to send text messages with @user
func (bot Bot) PostTextMention(ctx context.Context, text string, atUserID string, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgText, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	tb := NewTextBuilder()
	om := mb.Text(tb.Text(text).Mention(atUserID).Render()).Build()
	return bot.PostMessage(ctx, om)
}

// PostTextMentionAll is a simple way to send text messages with @all
func (bot Bot) PostTextMentionAll(ctx context.Context, text string, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgText, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	tb := NewTextBuilder()
	om := mb.Text(tb.Text(text).MentionAll().Render()).Build()
	return bot.PostMessage(ctx, om)
}

// PostTextMentionAndReply is a simple way to send text messages with @user and reply a message
func (bot Bot) PostTextMentionAndReply(ctx context.Context, text string, atUserID string, userID *OptionalUserID, replyID string) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgText, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	tb := NewTextBuilder()
	om := mb.Text(tb.Text(text).Mention(atUserID).Render()).BindReply(replyID).Build()
	return bot.PostMessage(ctx, om)
}

// PostImage is a simple way to send image
func (bot Bot) PostImage(ctx context.Context, imageKey string, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgImage, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	om := mb.Image(imageKey).Build()
	return bot.PostMessage(ctx, om)
}

// PostShareChat is a simple way to share chat
func (bot Bot) PostShareChat(ctx context.Context, chatID string, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgShareCard, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	om := mb.ShareChat(chatID).Build()
	return bot.PostMessage(ctx, om)
}

// PostShareUser is a simple way to share user
func (bot Bot) PostShareUser(ctx context.Context, openID string, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgShareUser, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	om := mb.ShareUser(openID).Build()
	return bot.PostMessage(ctx, om)
}

// PostMessage posts a message
func (bot Bot) PostMessage(ctx context.Context, om OutcomingMessage) (*PostMessageResponse, error) {
	req, err := BuildMessage(om)
	if err != nil {
		return nil, err
	}
	var respData PostMessageResponse
	if om.RootID == "" {
		err = bot.PostAPIRequest(ctx, "PostMessage", fmt.Sprintf(messageURL, om.UIDType), true, req, &respData)
	} else {
		resp, err := bot.ReplyMessage(ctx, om)
		return resp, err
	}
	return &respData, err
}

// ReplyMessage replies a message
func (bot Bot) ReplyMessage(ctx context.Context, om OutcomingMessage) (*PostMessageResponse, error) {
	req, err := buildReplyMessage(om)
	if err != nil {
		return nil, err
	}
	if om.RootID == "" {
		return nil, ErrParamMessageID
	}
	var respData PostMessageResponse
	err = bot.PostAPIRequest(ctx, "ReplyMessage", fmt.Sprintf(replyMessageURL, om.RootID), true, req, &respData)
	return &respData, err
}

// AddReaction adds reaction to a message
func (bot Bot) AddReaction(ctx context.Context, messageID string, emojiType EmojiType) (*ReactionResponse, error) {
	req := map[string]interface{}{
		"reaction_type": map[string]interface{}{
			"emoji_type": emojiType,
		},
	}
	var respData ReactionResponse
	err := bot.PostAPIRequest(ctx, "AddReaction", fmt.Sprintf(reactionsMessageURL, messageID), true, req, &respData)
	return &respData, err
}

// DeleteReaction deletes reaction of a message
func (bot Bot) DeleteReaction(ctx context.Context, messageID string, reactionID string) (*ReactionResponse, error) {
	var respData ReactionResponse
	err := bot.DeleteAPIRequest(ctx, "DeleteReaction", fmt.Sprintf(deleteReactionsMessageURL, messageID, reactionID), true, nil, &respData)
	return &respData, err
}

// UpdateMessage updates a message
func (bot Bot) UpdateMessage(ctx context.Context, messageID string, om OutcomingMessage) (*UpdateMessageResponse, error) {
	if om.MsgType != MsgInteractive &&
		om.MsgType != MsgText &&
		om.MsgType != MsgPost {
		return nil, ErrMessageType
	}
	req, err := buildUpdateMessage(om)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(updateMessageURL, messageID)
	var respData UpdateMessageResponse
	if om.MsgType == MsgInteractive {
		err = bot.PatchAPIRequest(ctx, "UpdateMessage", url, true, req, &respData)
	} else {
		err = bot.PutAPIRequest(ctx, "UpdateMessage", url, true, req, &respData)
	}
	return &respData, err
}

// GetMessage gets a message with im/v1
func (bot Bot) GetMessage(ctx context.Context, messageID string) (*GetMessageResponse, error) {
	var respData GetMessageResponse
	err := bot.GetAPIRequest(ctx, "GetMessage", fmt.Sprintf(getMessageURL, messageID), true, nil, &respData)
	return &respData, err
}

// RecallMessage recalls a message with ID
func (bot Bot) RecallMessage(ctx context.Context, messageID string) (*RecallMessageResponse, error) {
	url := fmt.Sprintf(recallMessageURL, messageID)
	var respData RecallMessageResponse
	err := bot.DeleteAPIRequest(ctx, "RecallMessage", url, true, nil, &respData)
	return &respData, err
}

// MessageReadReceipt queries message read receipt
func (bot Bot) MessageReadReceipt(ctx context.Context, messageID string, pageToken string, pageSize int) (*MessageReceiptResponse, error) {
	url := fmt.Sprintf(messageReceiptURL, messageID, bot.userIDType, pageSize, pageToken)
	var respData MessageReceiptResponse
	err := bot.GetAPIRequest(ctx, "MessageReadReceipt", url, true, nil, &respData)
	return &respData, err
}

// PostEphemeralMessage posts an ephemeral message
func (bot Bot) PostEphemeralMessage(ctx context.Context, om OutcomingMessage) (*PostEphemeralMessageResponse, error) {
	if om.UIDType == UIDUnionID {
		return nil, ErrUnsupportedUIDType
	}
	params := buildEphemeralCard(om)
	var respData PostEphemeralMessageResponse
	err := bot.PostAPIRequest(ctx, "PostEphemeralMessage", ephemeralMessageURL, true, params, &respData)
	return &respData, err
}

// DeleteEphemeralMessage deletes an ephemeral message
func (bot Bot) DeleteEphemeralMessage(ctx context.Context, messageID string) (*DeleteEphemeralMessageResponse, error) {
	params := map[string]interface{}{
		"message_id": messageID,
	}
	var respData DeleteEphemeralMessageResponse
	err := bot.PostAPIRequest(ctx, "DeleteEphemeralMessage", deleteEphemeralMessageURL, true, params, &respData)
	return &respData, err
}

// PinMessage pins a message
func (bot Bot) PinMessage(ctx context.Context, messageID string) (*PinMessageResponse, error) {
	params := map[string]interface{}{
		"message_id": messageID,
	}
	var respData PinMessageResponse
	err := bot.PostAPIRequest(ctx, "PinMessage", pinMessageURL, true, params, &respData)
	return &respData, err
}

// UnpinMessage unpins a message
func (bot Bot) UnpinMessage(ctx context.Context, messageID string) (*UnpinMessageResponse, error) {
	url := fmt.Sprintf(unpinMessageURL, messageID)
	var respData UnpinMessageResponse
	err := bot.DeleteAPIRequest(ctx, "PinMessage", url, true, nil, &respData)
	return &respData, err
}

// ForwardMessage forwards a message
func (bot Bot) ForwardMessage(ctx context.Context, messageID string, receiveID *OptionalUserID) (*ForwardMessageResponse, error) {
	url := fmt.Sprintf(forwardMessageURL, messageID, receiveID.UIDType)
	params := map[string]interface{}{
		"receive_id": receiveID.RealID,
	}
	var respData ForwardMessageResponse
	err := bot.PostAPIRequest(ctx, "ForwardMessage", url, true, params, &respData)
	return &respData, err
}
