package lark

import "fmt"

const (
	imMessageURL = "/open-apis/im/v1/messages?receive_id_type=%s"
)

// IMMessageRequest .
type IMMessageRequest struct {
	ReceiveID string `json:"receive_id"`
	Content   string `json:"content"`
	MsgType   string `json:"msg_type"`
}

// IMMessageResponse .
type IMMessageResponse struct {
	BaseResponse

	Data struct {
		MessageID      string `json:"message_id"`
		UpperMessageID string `json:"upper_message_id"`
		RootID         string `json:"root_id"`
		ParentID       string `json:"parent_id"`
		ChatID         string `json:"chat_id"`
		MsgType        string `json:"msg_type"`
		CreateTime     string `json:"create_time"`
		UpdateTime     string `json:"update_time"`
		Deleted        bool   `json:"deleted"`
		Updated        bool   `json:"updated"`
		Body           struct {
			Content string `json:"content"`
		}
	} `json:"data"`
}

// PostIMMessage posts message with im/v1
func (bot Bot) PostIMMessage(om OutcomingMessage) (*PostMessageResponse, error) {
	req := BuildIMMessage(om)
	if req == nil {
		return nil, ErrMessageNotBuild
	}
	var respData PostMessageResponse
	err := bot.PostAPIRequest("PostIMMessage", fmt.Sprintf(imMessageURL, om.UIDType), true, req, &respData)
	return &respData, err
}
