package lark

import (
	"time"
)

// URLs for auth
const (
	appAccessTokenInternalURL       = "/open-apis/auth/v3/app_access_token/internal"
	tenantAppAccessTokenInternalURL = "/open-apis/auth/v3/tenant_access_token/internal/"
)

// AuthTokenInternalResponse .
type AuthTokenInternalResponse struct {
	Code           int    `json:"code"`
	AppAccessToken string `json:"app_access_token"`
	Expire         int    `json:"expire"`
}

// TenantAuthTokenInternalResponse .
type TenantAuthTokenInternalResponse struct {
	Code                 int    `json:"code"`
	TenantAppAccessToken string `json:"tenant_access_token"`
	Expire               int    `json:"expire"`
}

// GetAccessTokenInternal gets AppAccessToken for internal use
func (bot *Bot) GetAccessTokenInternal(updateToken bool) (*AuthTokenInternalResponse, error) {
	if !bot.requireType(ChatBot) {
		return nil, ErrBotTypeError
	}

	params := map[string]interface{}{
		"app_id":     bot.appID,
		"app_secret": bot.appSecret,
	}
	var respData AuthTokenInternalResponse
	err := bot.PostAPIRequest("GetAccessTokenInternal", appAccessTokenInternalURL, false, params, &respData)
	if err == nil && updateToken {
		bot.accessToken = respData.AppAccessToken
	}
	return &respData, err
}

// GetTenantAccessTokenInternal gets AppAccessToken for internal use
func (bot *Bot) GetTenantAccessTokenInternal(updateToken bool) (*TenantAuthTokenInternalResponse, error) {
	if !bot.requireType(ChatBot) {
		return nil, ErrBotTypeError
	}

	params := map[string]interface{}{
		"app_id":     bot.appID,
		"app_secret": bot.appSecret,
	}
	var respData TenantAuthTokenInternalResponse
	err := bot.PostAPIRequest("GetTenantAccessTokenInternal", tenantAppAccessTokenInternalURL, false, params, &respData)
	if err == nil && updateToken {
		bot.tenantAccessToken = respData.TenantAppAccessToken
	}
	return &respData, err
}

// StopHeartbeat stop auto-renew
func (bot *Bot) StopHeartbeat() {
	bot.heartbeat <- true
}

// StartHeartbeat renew auth token periodically
func (bot *Bot) StartHeartbeat() {
	if !bot.requireType(ChatBot) {
		bot.logger.Log(bot.ctx, LogLevelError, "Heartbeat only support Chat Bot")
		return
	}

	bot.heartbeat = make(chan bool)
	go func() {
		for {
			next := time.Second * 10
			resp, err := bot.GetTenantAccessTokenInternal(true)
			if err != nil {
				bot.httpErrorLog("Heartbeat", "failed to get tenant access token", err)
			}
			if resp != nil && resp.Expire-20 > 0 {
				next = time.Duration(resp.Expire-20) * time.Second
			}
			if bot.debugHeartbeat > 0 {
				next = time.Second * 1
			}
			t := time.NewTimer(next)
			select {
			case <-bot.heartbeat:
				return
			case <-t.C:
				if bot.debugHeartbeat > 0 {
					bot.debugHeartbeat++
				}
			}
		}
	}()
}
