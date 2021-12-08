package lark

import (
	"encoding/json"
)

// BuildIMMessage .
func BuildIMMessage(om OutcomingMessage) *IMMessageRequest {
	req := IMMessageRequest{
		MsgType:   string(om.MsgType),
		Content:   buildIMContent(om),
		ReceiveID: buildReceiveID(om),
	}
	if req.ReceiveID == "" || req.Content == "" {
		return nil
	}
	return &req
}

func buildIMContent(om OutcomingMessage) string {
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
	case MsgShareCard:
		b, err = json.Marshal(om.Content.ShareChat)
	case MsgShareUser:
		b, err = json.Marshal(om.Content.ShareUser)
	case MsgPost:
		b, err = json.Marshal(om.Content.Post)
	case MsgInteractive:
		b, err = json.Marshal(om.Content.Card)
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
