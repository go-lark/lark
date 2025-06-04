package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTenantAccessTokenInternal(t *testing.T) {
	bot := newTestBot()
	resp, err := bot.GetTenantAccessTokenInternal(t.Context())
	if assert.NoError(t, err) {
		assert.Equal(t, 0, resp.Code)
		assert.NotEmpty(t, resp.TenantAccessToken)
		assert.NotEmpty(t, resp.Expire)
		t.Log(resp)
	}
}
