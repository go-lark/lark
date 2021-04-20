package lark

const (
	getBotInfoURL = "/open-apis/bot/v3/info/"
)

// GetBotInfoResponse .
type GetBotInfoResponse struct {
	Code int `json:"code"`
	Bot  struct {
		ActivateStatus int      `json:"activate_status"`
		AppName        string   `json:"app_name"`
		AvatarURL      string   `json:"avatar_url"`
		IPWhiteList    []string `json:"ip_white_list"`
		OpenID         string   `json:"open_id"`
	} `json:"bot"`
}

// GetBotInfo returns bot info
func (bot *Bot) GetBotInfo() (*GetBotInfoResponse, error) {
	var respData GetBotInfoResponse
	err := bot.PostAPIRequestWithAuth("GetBotInfo", getBotInfoURL, nil, &respData)
	return &respData, err
}
