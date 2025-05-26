package lark

// BaseResponse of an API
type BaseResponse struct {
	Code  int       `json:"code"`
	Msg   string    `json:"msg"`
	Error BaseError `json:"error"`
}

// DummyResponse is used to unmarshal from a complete JSON response but only to retrieve error
type DummyResponse struct {
	BaseResponse
}

// BaseError is returned by the platform
type BaseError struct {
	LogID string `json:"log_id,omitempty"`
}

// I18NNames structure of names in multiple locales
type I18NNames struct {
	ZhCN string `json:"zh_cn,omitempty"`
	EnUS string `json:"en_us,omitempty"`
	JaJP string `json:"ja_jp,omitempty"`
}

// WithUserIDType assigns user ID type
func (bot *Bot) WithUserIDType(userIDType string) *Bot {
	bot.userIDType = userIDType
	return bot
}
