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

// This appID & appSecret is for test use
var (
	testAppID       string
	testAppSecret   string
	testUserEmail   string
	testUserOpenID  string
	testGroupChatID string
	testMessageID   string
)

func newTestBot() *Bot {
	testMode := os.Getenv("GO_LARK_TEST_MODE")
	if testMode == "" {
		testMode = "testing"
	}
	if testMode == "local" {
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	testAppID = os.Getenv("lark_app_id")
	testAppSecret = os.Getenv("lark_app_secret")

	testUserEmail = os.Getenv("lark_user_email")
	testUserOpenID = os.Getenv("lark_open_id")
	testGroupChatID = os.Getenv("lark_chat_id")
	testMessageID = os.Getenv("lark_message_id")
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

func TestBotProperties(t *testing.T) {
	chatBot := newTestBot()
	assert.NotEmpty(t, chatBot.appID)
	assert.NotEmpty(t, chatBot.appSecret)
	assert.Empty(t, chatBot.webhook)
	assert.Equal(t, defaultDomain, chatBot.domain)
	assert.Equal(t, ChatBot, chatBot.botType)
	assert.NotNil(t, chatBot.client)
	assert.NotNil(t, chatBot.logger)

	notifyBot := NewNotificationBot(ciHookURLV1)
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
	assert.Equal(t, "https://open.feishu.cn", bot.domain)
	bot.SetDomain("https://test.test")
	assert.Equal(t, "https://test.test", bot.domain)
}

func TestBotGetters(t *testing.T) {
	bot := newTestBot()
	assert.Equal(t, testAppID, bot.AppID())
	assert.Equal(t, ChatBot, bot.BotType())
	assert.Equal(t, "", bot.AccessToken())
	assert.Equal(t, "", bot.TenantAccessToken())
}

func TestSetLogger(t *testing.T) {
	bot := newTestBot()
	newLogger := log.New(os.Stderr, "new logger", log.LstdFlags)
	bot.SetLogger(newLogger)
	assert.Equal(t, newLogger, bot.logger)
	assert.Equal(t, "new logger", bot.logger.Prefix())
}
