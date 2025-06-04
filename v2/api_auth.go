package lark

import (
	"context"
)

// URLs for auth
const (
	tenantAccessTokenInternalURL = "/open-apis/auth/v3/tenant_access_token/internal"
)

// TenantAccessTokenInternalResponse .
type TenantAccessTokenInternalResponse struct {
	BaseResponse
	TenantAccessToken string `json:"tenant_access_token"`
	Expire            int    `json:"expire"`
}

// GetTenantAccessTokenInternal gets AppAccessToken for internal use
func (bot *Bot) GetTenantAccessTokenInternal(ctx context.Context) (*TenantAccessTokenInternalResponse, error) {
	if !bot.requireType(ChatBot) {
		return nil, ErrBotTypeError
	}

	params := map[string]interface{}{
		"app_id":     bot.appID,
		"app_secret": bot.appSecret,
	}
	var respData TenantAccessTokenInternalResponse
	err := bot.PostAPIRequest(ctx, "GetTenantAccessTokenInternal", tenantAccessTokenInternalURL, false, params, &respData)
	return &respData, err
}
