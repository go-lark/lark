package lark

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithFunctions(t *testing.T) {
	emailUID := WithEmail(testUserEmail)
	assert.Equal(t, "email", emailUID.IDType)
	assert.Equal(t, testUserEmail, emailUID.RealID)

	openIDUID := WithOpenID(testUserOpenID)
	assert.Equal(t, "open_id", openIDUID.IDType)
	assert.Equal(t, testUserOpenID, openIDUID.RealID)

	chatIDUID := WithChatID(testGroupChatID)
	assert.Equal(t, "chat_id", chatIDUID.IDType)
	assert.Equal(t, testGroupChatID, chatIDUID.RealID)

	fakeUID := "6893390418998738946"
	userIDUID := WithUserID(fakeUID)
	assert.Equal(t, "user_id", userIDUID.IDType)
	assert.Equal(t, fakeUID, userIDUID.RealID)
}
