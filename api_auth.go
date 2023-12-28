package lark

import (
	"errors"
	"sync/atomic"
	"time"
)

// URLs for auth
const (
	appAccessTokenInternalURL       = "/open-apis/auth/v3/app_access_token/internal"
	tenantAppAccessTokenInternalURL = "/open-apis/auth/v3/tenant_access_token/internal/"
)

// AuthTokenInternalResponse .
type AuthTokenInternalResponse struct {
	BaseResponse
	AppAccessToken string `json:"app_access_token"`
	Expire         int    `json:"expire"`
}

// TenantAuthTokenInternalResponse .
type TenantAuthTokenInternalResponse struct {
	BaseResponse
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
		bot.accessToken.Store(respData.AppAccessToken)
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
		bot.tenantAccessToken.Store(respData.TenantAppAccessToken)
	}
	return &respData, err
}

// StopHeartbeat stop auto-renew
func (bot *Bot) StopHeartbeat() {
	bot.heartbeat <- true
}

// StartHeartbeat renew auth token periodically
func (bot *Bot) StartHeartbeat() error {
	return bot.startHeartbeat(10 * time.Second)
}

func (bot *Bot) startHeartbeat(defaultInterval time.Duration) error {
	if !bot.requireType(ChatBot) {
		return ErrBotTypeError
	}

	// First initialize the token in blocking mode
	resp, err := bot.GetTenantAccessTokenInternal(true)
	if err != nil {
		return err
	}
	if resp == nil {
		return ErrNilResponse
	}
	if resp.Code != 0 {
		return errors.New(resp.Msg)
	}
	atomic.AddInt64(&bot.heartbeatCounter, 1)

	interval := defaultInterval
	bot.heartbeat = make(chan bool)
	go func() {
		for {
			t := time.NewTimer(interval)
			select {
			case <-bot.heartbeat:
				return
			case <-t.C:
				interval = defaultInterval
				resp, err := bot.GetTenantAccessTokenInternal(true)
				// if GetTenantAccessTokenInternal failed, then we would have retried 100ms later
				if err != nil {
					bot.httpErrorLog("Heartbeat", "failed to get tenant access token", err)
					interval = time.Duration(100 * time.Millisecond)
				} else if resp == nil {
					bot.httpErrorLog("Heartbeat", "response is nil", ErrNilResponse)
					interval = time.Duration(100 * time.Millisecond)
				} else if resp.Code != 0 {
					bot.httpErrorLog("Heartbeat", "failed to get tenant access token", errors.New(resp.Msg))
					interval = time.Duration(100 * time.Millisecond)
				} else {
					atomic.AddInt64(&bot.heartbeatCounter, 1)
					if resp.Expire-20 > 0 {
						interval = time.Duration(resp.Expire-20) * time.Second
					}
				}
			}
		}
	}()
	return nil
}
