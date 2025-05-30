package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUserInfo(t *testing.T) {
	resp, err := bot.GetUserInfo(t.Context(), WithUserID(testUserID))
	if assert.NoError(t, err) {
		assert.Equal(t, resp.Data.User.Name, "David")
	}
	bresp, err := bot.BatchGetUserInfo(t.Context(), UIDUserID, testUserID)
	if assert.NoError(t, err) {
		if assert.NotEmpty(t, bresp.Data.Items) {
			assert.Equal(t, bresp.Data.Items[0].Name, "David")
		}
	}
	_, err = bot.BatchGetUserInfo(t.Context(), UIDUserID)
	assert.ErrorIs(t, err, ErrParamExceedInputLimit)
}
