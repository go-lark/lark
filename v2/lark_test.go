package lark

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// IDs for test use
var (
	testAppID         string
	testAppSecret     string
	testUserEmail     string
	testUserOpenID    string
	testUserID        string
	testUserUnionID   string
	testGroupChatID   string
	testWebhookV1     string
	testWebhook       string
	testWebhookSigned string
)

func newTestBot() *Bot {
	testMode := os.Getenv("GO_LARK_TEST_MODE")
	if testMode == "" {
		testMode = "testing"
	}
	if testMode == "local" {
		err := godotenv.Load(".env")
		if err != nil {
			panic(err)
		}
	}
	testAppID = os.Getenv("LARK_APP_ID")
	testAppSecret = os.Getenv("LARK_APP_SECRET")
	testUserEmail = os.Getenv("LARK_USER_EMAIL")
	testUserID = os.Getenv("LARK_USER_ID")
	testUserUnionID = os.Getenv("LARK_UNION_ID")
	testUserOpenID = os.Getenv("LARK_OPEN_ID")
	testGroupChatID = os.Getenv("LARK_CHAT_ID")
	testWebhook = os.Getenv("LARK_WEBHOOK")
	testWebhookSigned = os.Getenv("LARK_WEBHOOK_SIGNED")
	if len(testAppID) == 0 ||
		len(testAppSecret) == 0 ||
		len(testUserEmail) == 0 ||
		len(testUserID) == 0 ||
		len(testUserUnionID) == 0 ||
		len(testUserOpenID) == 0 ||
		len(testGroupChatID) == 0 {
		panic("insufficient test environment")
	}
	return NewChatBot(testAppID, testAppSecret)
}

func captureOutput(f func()) string {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	f()
	log.SetOutput(os.Stderr)
	return buf.String()
}

func performRequest(r http.HandlerFunc, method, path string, body interface{}) *httptest.ResponseRecorder {
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(body)
	req := httptest.NewRequest(method, path, buf)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

// for general API test suites
var bot *Bot

func init() {
	bot = newTestBot()
}

func TestBotProperties(t *testing.T) {
	chatBot := newTestBot()
	assert.NotEmpty(t, chatBot.appID)
	assert.NotEmpty(t, chatBot.appSecret)
	assert.Empty(t, chatBot.webhook)
	assert.Equal(t, DomainFeishu, chatBot.domain)
	assert.Equal(t, ChatBot, chatBot.botType)
	assert.NotNil(t, chatBot.client)
	assert.NotNil(t, chatBot.logger)

	notifyBot := NewNotificationBot(testWebhook)
	assert.Empty(t, notifyBot.appID)
	assert.Empty(t, notifyBot.appSecret)
	assert.NotEmpty(t, notifyBot.webhook)
	assert.Empty(t, notifyBot.domain)
	assert.Equal(t, NotificationBot, notifyBot.botType)
	assert.NotNil(t, notifyBot.client)
	assert.NotNil(t, notifyBot.logger)
}

func TestRequiredType(t *testing.T) {
	bot := newTestBot()
	assert.True(t, bot.requireType(ChatBot))
	assert.False(t, bot.requireType(NotificationBot))
}

func TestSetDomain(t *testing.T) {
	bot := newTestBot()
	assert.Equal(t, DomainFeishu, bot.domain)
	assert.Equal(t, DomainFeishu, bot.Domain())
	bot.SetDomain("https://test.test")
	assert.Equal(t, "https://test.test", bot.domain)
	assert.Equal(t, "https://test.test", bot.Domain())
}

func TestBotGetters(t *testing.T) {
	bot := newTestBot()
	assert.Equal(t, testAppID, bot.AppID())
	assert.Equal(t, ChatBot, bot.BotType())
}

func TestSetClient(t *testing.T) {
	bot := &Bot{}
	assert.Nil(t, bot.client)
	bot.SetClient(newDefaultClient())
	assert.NotNil(t, bot.client)
}

func TestSetWebhook(t *testing.T) {
	bot := NewNotificationBot("abc")
	assert.Equal(t, "abc", bot.webhook)
	bot.SetWebhook("def")
	assert.Equal(t, "def", bot.webhook)
}

func TestSetAutoRenew(t *testing.T) {
	bot := newTestBot()
	assert.True(t, bot.autoRenew)
	bot.SetAutoRenew(false)
	assert.False(t, bot.autoRenew)
	bot.SetAutoRenew(true)
	assert.True(t, bot.autoRenew)
}

func TestTenantAccessToken(t *testing.T) {
	bot := newTestBot()
	assert.Equal(t, "", bot.TenantAccessToken())
	bot.SetTenantAccessToken(TenantAccessToken{
		TenantAccessToken: "test",
	})
	assert.Equal(t, "test", bot.TenantAccessToken())
}
