package lark

import (
	"encoding/json"
	"strconv"
)

// BuildMessage .
func BuildMessage(om OutcomingMessage) (*IMMessageRequest, error) {
	req := IMMessageRequest{
		MsgType:   string(om.MsgType),
		Content:   buildContent(om),
		ReceiveID: buildReceiveID(om),
	}
	if req.ReceiveID == "" {
		return nil, ErrInvalidReceiveID
	}
	if req.Content == "" {
		return nil, ErrMessageNotBuild
	}
	if om.UUID != "" {
		req.UUID = om.UUID
	}
	return &req, nil
}

func buildReplyMessage(om OutcomingMessage) (*IMMessageRequest, error) {
	req := IMMessageRequest{
		MsgType:   string(om.MsgType),
		Content:   buildContent(om),
		ReceiveID: buildReceiveID(om),
	}
	if req.Content == "" {
		return nil, ErrMessageNotBuild
	}
	if om.ReplyInThread {
		req.ReplyInThread = om.ReplyInThread
	}
	if om.UUID != "" {
		req.UUID = om.UUID
	}

	return &req, nil
}

func buildUpdateMessage(om OutcomingMessage) (*IMMessageRequest, error) {
	req := IMMessageRequest{
		Content: buildContent(om),
	}
	if om.MsgType != MsgInteractive {
		req.MsgType = om.MsgType
	}
	if req.Content == "" {
		return nil, ErrMessageNotBuild
	}

	return &req, nil
}

func buildContent(om OutcomingMessage) string {
	var (
		content = ""
		b       []byte
		err     error
	)
	switch om.MsgType {
	case MsgText:
		b, err = json.Marshal(om.Content.Text)
	case MsgImage:
		b, err = json.Marshal(om.Content.Image)
	case MsgFile:
		b, err = json.Marshal(om.Content.File)
	case MsgShareCard:
		b, err = json.Marshal(om.Content.ShareChat)
	case MsgShareUser:
		b, err = json.Marshal(om.Content.ShareUser)
	case MsgPost:
		b, err = json.Marshal(om.Content.Post)
	case MsgInteractive:
		if om.Content.Card != nil {
			b, err = json.Marshal(om.Content.Card)
		} else if om.Content.Template != nil {
			b, err = json.Marshal(om.Content.Template)
		}
	case MsgAudio:
		b, err = json.Marshal(om.Content.Audio)
	case MsgMedia:
		b, err = json.Marshal(om.Content.Media)
	case MsgSticker:
		b, err = json.Marshal(om.Content.Sticker)
	}
	if err != nil {
		return ""
	}
	content = string(b)

	return content
}

func buildReceiveID(om OutcomingMessage) string {
	switch om.UIDType {
	case UIDEmail:
		return om.Email
	case UIDUserID:
		return om.UserID
	case UIDOpenID:
		return om.OpenID
	case UIDChatID:
		return om.ChatID
	case UIDUnionID:
		return om.UnionID
	}
	return ""
}

func buildEphemeralCard(om OutcomingMessage) map[string]interface{} {
	params := map[string]interface{}{
		"msg_type": om.MsgType,
		"chat_id":  om.ChatID, // request must contain chat_id, even if it is empty
	}
	params[om.UIDType] = buildReceiveID(om)
	if len(om.RootID) > 0 {
		params["root_id"] = om.RootID
	}
	content := make(map[string]interface{})
	if om.Content.Text != nil {
		content["text"] = om.Content.Text.Text
	}
	if om.Content.Image != nil {
		content["image_key"] = om.Content.Image.ImageKey
	}
	if om.Content.ShareChat != nil {
		content["share_open_chat_id"] = om.Content.ShareChat.ChatID
	}
	if om.Content.Post != nil {
		content["post"] = *om.Content.Post
	}
	if om.MsgType == MsgInteractive && om.Content.Card != nil {
		params["card"] = *om.Content.Card
	}
	if len(om.Sign) > 0 {
		params["sign"] = om.Sign
		params["timestamp"] = strconv.FormatInt(om.Timestamp, 10)
	}
	params["content"] = content
	return params
}

func buildNotification(om OutcomingMessage) map[string]interface{} {
	params := map[string]interface{}{
		"msg_type": om.MsgType,
	}
	if len(om.RootID) > 0 {
		params["root_id"] = om.RootID
	}
	content := make(map[string]interface{})
	if om.Content.Text != nil {
		content["text"] = om.Content.Text.Text
	}
	if om.Content.Image != nil {
		content["image_key"] = om.Content.Image.ImageKey
	}
	if om.Content.ShareChat != nil {
		content["share_open_chat_id"] = om.Content.ShareChat.ChatID
	}
	if om.Content.Post != nil {
		content["post"] = *om.Content.Post
	}
	if om.MsgType == MsgInteractive && om.Content.Card != nil {
		params["card"] = *om.Content.Card
	}
	if len(om.Sign) > 0 {
		params["sign"] = om.Sign
		params["timestamp"] = strconv.FormatInt(om.Timestamp, 10)
	}
	params["content"] = content
	return params
}
