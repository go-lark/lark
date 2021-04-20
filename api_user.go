package lark

const (
	getUserInfoURL      = "/open-apis/user/v4/info"
	getUserIDByEmailURL = "/open-apis/user/v4/email2id"
)

// GetUserInfoResponse .
type GetUserInfoResponse struct {
	Code int `json:"code"`
	Data struct {
		OpenID string `json:"open_id"`
		UserID string `json:"user_id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
		Avatar string `json:"avatar"`
	} `json:"data"`
}

// GetUserIDByEmailResponse .
type GetUserIDByEmailResponse struct {
	Code int `json:"code"`
	Data struct {
		OpenID string `json:"open_id"`
		UserID string `json:"user_id"`
	} `json:"data"`
}

// GetUserInfo returns user info
func (bot *Bot) GetUserInfo(openID string) (*GetUserInfoResponse, error) {
	params := map[string]interface{}{
		"open_id": openID,
	}
	var respData GetUserInfoResponse
	err := bot.PostAPIRequestWithAuth("GetUserInfo", getUserInfoURL, params, &respData)
	return &respData, err
}

// GetUserIDByEmail converts email to ID
func (bot *Bot) GetUserIDByEmail(email string) (*GetUserIDByEmailResponse, error) {
	params := map[string]interface{}{
		"email": email,
	}
	var respData GetUserIDByEmailResponse
	err := bot.PostAPIRequestWithAuth("GetUserIDByEmail", getUserIDByEmailURL, params, &respData)
	return &respData, err
}
