package lark

import "fmt"

const (
	postIMMessageURL = "/open-apis/im/v1/messages?receive_id_type=%s"
	getIMMessageURL  = "/open-apis/im/v1/messages/%s"
)

// IMMessageRequest .
type IMMessageRequest struct {
	ReceiveID string `json:"receive_id"`
	Content   string `json:"content"`
	MsgType   string `json:"msg_type"`
}

// IMSendor .
type IMSendor struct {
	ID         string `json:"id"`
	IDType     string `json:"id_type"`
	SendorType string `json:"sendor_type"`
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
	Sendor         IMSendor
	Mentions       []IMMention
	Body           IMBody
}

// PostIMMessageResponse .
type PostIMMessageResponse struct {
	BaseResponse

	Data IMMessage `json:"data"`
}

// GetIMMessageResponse .
type GetIMMessageResponse struct {
	BaseResponse

	Data struct {
		Items []IMMessage `json:"items"`
	} `json:"data"`
}

// PostIMMessage posts message with im/v1
func (bot Bot) PostIMMessage(om OutcomingMessage) (*PostIMMessageResponse, error) {
	req, err := BuildIMMessage(om)
	if err != nil {
		return nil, err
	}
	var respData PostIMMessageResponse
	err = bot.PostAPIRequest("PostIMMessage", fmt.Sprintf(postIMMessageURL, om.UIDType), true, req, &respData)
	return &respData, err
}

// GetIMMessage posts message with im/v1
func (bot Bot) GetIMMessage(messageID string) (*GetIMMessageResponse, error) {
	var respData GetIMMessageResponse
	err := bot.GetAPIRequest("GetIMMessage", fmt.Sprintf(getIMMessageURL, messageID), true, nil, &respData)
	return &respData, err
}
