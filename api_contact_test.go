package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserInfo(t *testing.T) {
	resp, err := bot.GetUserInfo(WithUserID(testUserID))
	if assert.NoError(t, err) {
		assert.Equal(t, resp.Data.User.Name, "David")
	}
	bresp, err := bot.BatchGetUserInfo(UIDUserID, testUserID)
	if assert.NoError(t, err) {
		if assert.NotEmpty(t, bresp.Data.Items) {
			assert.Equal(t, bresp.Data.Items[0].Name, "David")
		}
	}
}
