package lark

import "context"

// PostNotificationResp .
type PostNotificationResp struct {
	Code          int    `json:"code"`
	Msg           string `json:"msg"`
	StatusCode    int    `json:"StatusCode"`
	StatusMessage string `json:"StatusMessage"`
}

// PostNotification posts nofication to a given webhook
func (bot *Bot) PostNotification(ctx context.Context, om OutcomingMessage) (*PostNotificationResp, error) {
	if !bot.requireType(NotificationBot) {
		return nil, ErrBotTypeError
	}

	params := buildNotification(om)
	var respData PostNotificationResp
	err := bot.PostAPIRequest(ctx, "PostNotification", bot.webhook, false, params, &respData)
	return &respData, err
}
