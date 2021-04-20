package lark

import (
	"log"
	"net/http"
	"time"
)

const (
	// ChatBot should call NewChatBot
	// Create from https://open.feishu.cn/ or https://open.larksuite.com/
	ChatBot = iota
	// NotificationBot for webhook, behave as a simpler notification bot
	// Create from Lark group
	NotificationBot
)

// Bot definition
type Bot struct {
	// bot type
	botType int
	// Auth info
	appID             string
	appSecret         string
	accessToken       string
	tenantAccessToken string
	// webhook for NotificationBot
	webhook string
	// API Domain
	domain string
	// http client
	client *http.Client
	// auth heartbeat
	heartbeat      chan bool
	debugHeartbeat int

	logger *log.Logger
}

// default domain
const defaultDomain = "https://open.feishu.cn"

// NewChatBot with appID and appSecret
func NewChatBot(appID, appSecret string) *Bot {
	return &Bot{
		botType:   ChatBot,
		appID:     appID,
		appSecret: appSecret,
		client:    initClient(),
		domain:    defaultDomain,
		logger:    initDefaultLogger(),
	}
}

// NewNotificationBot with URL
func NewNotificationBot(hookURL string) *Bot {
	return &Bot{
		botType: NotificationBot,
		webhook: hookURL,
		client:  initClient(),
		logger:  initDefaultLogger(),
	}
}

// requireType checks whether the action is allowed in a list of bot types
func (bot Bot) requireType(botType ...int) bool {
	for _, iterType := range botType {
		if bot.botType == iterType {
			return true
		}
	}
	return false
}

// SetClient assigns a new client to bot.client
func (bot *Bot) SetClient(c *http.Client) {
	bot.client = c
}

func initClient() *http.Client {
	return &http.Client{
		Timeout: 5 * time.Second,
	}
}

// SetDomain set domain of endpoint, so we could call Feishu/Lark
// go-lark does not check your host, just use the right one or fail.
func (bot *Bot) SetDomain(domain string) {
	bot.domain = domain
}

// AppID returns bot.appID for external use
func (bot Bot) AppID() string {
	return bot.appID
}

// BotType returns bot.botType for external use
func (bot Bot) BotType() int {
	return bot.botType
}

// AccessToken returns bot.accessToken for external use
func (bot Bot) AccessToken() string {
	return bot.accessToken
}

// TenantAccessToken returns bot.tenantAccessToken for external use
func (bot Bot) TenantAccessToken() string {
	return bot.tenantAccessToken
}
