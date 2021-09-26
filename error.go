package lark

import "errors"

// Errors
var (
	ErrBotTypeError           = errors.New("Bot type error")
	ErrParamUserID            = errors.New("Param error: UserID")
	ErrMessageTypeNotSuppored = errors.New("Message type not supported")
	ErrEncryptionNotEnabled   = errors.New("Encryption is not enabled")
	ErrCustomHTTPClientNotSet = errors.New("Custom HTTP client not set")
)
