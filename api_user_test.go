package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserInfo(t *testing.T) {
	bot := newTestBot()
	_, _ = bot.GetTenantAccessTokenInternal(true)
	resp, err := bot.GetUserInfo(testUserOpenID)
	if assert.NoError(t, err) {
		assert.Zero(t, resp.Code)
		assert.NotEmpty(t, resp.Data.OpenID)
		assert.NotEmpty(t, resp.Data.Name)
		assert.NotEmpty(t, resp.Data.Avatar)
		t.Log(resp)
	}
}

func TestGetUserIDByEmail(t *testing.T) {
	bot := newTestBot()
	_, _ = bot.GetTenantAccessTokenInternal(true)
	resp, err := bot.GetUserIDByEmail(testUserEmail)
	if assert.NoError(t, err) {
		assert.Zero(t, resp.Code)
		t.Log(resp.Data.OpenID)
		assert.NotEmpty(t, resp.Data.OpenID)
	}
}
