package lark

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
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
	testAppID = os.Getenv("LARK_APP_ID")
	testAppSecret = os.Getenv("LARK_APP_SECRET")

	testUserEmail = os.Getenv("LARK_USER_EMAIL")
	testUserOpenID = os.Getenv("LARK_OPEN_ID")
	testGroupChatID = os.Getenv("LARK_CHAT_ID")
	testMessageID = os.Getenv("LARK_MESSAGE_ID")
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
	_, _ = bot.GetTenantAccessTokenInternal(true)
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
	assert.Equal(t, DomainFeishu, bot.domain)
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

func TestSetClient(t *testing.T) {
	bot := &Bot{}
	assert.Nil(t, bot.client)
	bot.SetClient(&http.Client{})
	assert.NotNil(t, bot.client)
}

type customHTTPWrapper struct {
	client *http.Client
}

func (c customHTTPWrapper) Do(ctx context.Context, method, url string, header http.Header, body io.Reader) (io.ReadCloser, error) {
	return nil, nil
}

func TestCustomClient(t *testing.T) {
	bot := &Bot{}
	assert.Nil(t, bot.customClient)
	var c customHTTPWrapper
	bot.SetCustomClient(c)
	assert.NotNil(t, bot.customClient)
}
