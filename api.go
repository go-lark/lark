package lark

// BaseResponse of an API
type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
