package lark

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"
)

// PostNotificationResp response of PostNotification
type PostNotificationResp struct {
	Ok bool `json:"ok,omitempty"`
}

// PostNotificationV2Resp response of PostNotificationV2
type PostNotificationV2Resp struct {
	BaseResponse
	StatusCode    int    `json:"StatusCode"`
	StatusMessage string `json:"StatusMessage"`
}

// PostNotification send message to a webhook
func (bot *Bot) PostNotification(title, text string) (*PostNotificationResp, error) {
	if !bot.requireType(NotificationBot) {
		return nil, ErrBotTypeError
	}

	params := map[string]interface{}{
		"title": title,
		"text":  text,
	}
	var respData PostNotificationResp
	err := bot.PostAPIRequest("PostNotification", bot.webhook, false, params, &respData)
	return &respData, err
}

// PostNotificationV2 support v2 format
func (bot *Bot) PostNotificationV2(om OutcomingMessage) (*PostNotificationV2Resp, error) {
	if !bot.requireType(NotificationBot) {
		return nil, ErrBotTypeError
	}

	params := BuildOutcomingMessageReq(om)

	if bot.webhookSignSecret != "" {
		timestamp := time.Now().Unix()
		params["timestamp"] = timestamp
		params["sign"] = bot.webhookSignSecret
	}

	var respData PostNotificationV2Resp
	err := bot.PostAPIRequest("PostNotificationV2", bot.webhook, false, params, &respData)
	return &respData, err
}

func GenSign(secret string, timestamp int64) (string, error) {
	//timestamp + key 做sha256, 再进行base64 encode
	stringToSign := fmt.Sprintf("%v", timestamp) + "\n" + secret

	var data []byte
	h := hmac.New(sha256.New, []byte(stringToSign))
	_, err := h.Write(data)
	if err != nil {
		return "", err
	}

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}
