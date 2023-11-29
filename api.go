package lark

// BaseResponse of an API
type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

// I18NNames .
type I18NNames struct {
	ZhCN string `json:"zh_cn,omitempty"`
	EnUS string `json:"en_us,omitempty"`
	JaJP string `json:"ja_jp,omitempty"`
}

// WithUserIDType .
func (bot *Bot) WithUserIDType(userIDType string) *Bot {
	bot.userIDType = userIDType
	return bot
}
