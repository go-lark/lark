package lark

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAuthAccessTokenInternal(t *testing.T) {
	bot := newTestBot()
	resp, err := bot.GetAccessTokenInternal(true)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.AppAccessToken)
		t.Log(resp.AppAccessToken)
		assert.NotEmpty(t, resp.Expire)
	}
}

func TestAuthTenantAccessTokenInternal(t *testing.T) {
	bot := newTestBot()
	resp, err := bot.GetTenantAccessTokenInternal(true)
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.TenantAppAccessToken)
		t.Log(resp.TenantAppAccessToken)
		assert.NotEmpty(t, resp.Expire)
	}
}

func TestHeartbeat(t *testing.T) {
	bot := newTestBot()
	assert.Nil(t, bot.heartbeat)
	assert.Nil(t, bot.startHeartbeat(time.Second*1))
	assert.NotEmpty(t, bot.tenantAccessToken)
	assert.Equal(t, int64(1), bot.heartbeatCounter)
	time.Sleep(2 * time.Second)
	assert.Equal(t, int64(2), bot.heartbeatCounter)
	bot.StopHeartbeat()
	time.Sleep(2 * time.Second)
	assert.Equal(t, int64(2), bot.heartbeatCounter)
	// restart heartbeat
	assert.Nil(t, bot.startHeartbeat(time.Second*1))
	time.Sleep(2 * time.Second)
	assert.Equal(t, int64(4), bot.heartbeatCounter)
}

func TestInvalidHeartbeat(t *testing.T) {
	bot := NewNotificationBot("")
	err := bot.StartHeartbeat()
	assert.Error(t, err, ErrBotTypeError)
}
