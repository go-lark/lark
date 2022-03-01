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
	bot.debugHeartbeat.Store(1)
	assert.Nil(t, bot.heartbeat)
	bot.StartHeartbeat()
	time.Sleep(1 * time.Second)
	assert.NotEmpty(t, bot.tenantAccessToken)
	assert.Equal(t, 1, bot.debugHeartbeat.Load().(int))
	time.Sleep(1 * time.Second)
	assert.Equal(t, 2, bot.debugHeartbeat.Load().(int))
	bot.StopHeartbeat()
	time.Sleep(2 * time.Second)
	assert.Equal(t, 2, bot.debugHeartbeat.Load().(int))
	// restart heartbeat
	bot.StartHeartbeat()
	time.Sleep(2 * time.Second)
	assert.Equal(t, 3, bot.debugHeartbeat.Load().(int))
}

func TestInvalidHeartbeat(t *testing.T) {
	bot := NewNotificationBot("", "")
	output := bot.captureOutput(func() { bot.StartHeartbeat() })
	assert.Contains(t, output, "only support")
}
