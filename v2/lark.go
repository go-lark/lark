// Package lark is an easy-to-use SDK for Feishu and Lark Open Platform,
// which implements messaging APIs, with full-fledged supports on building Chat Bot and Notification Bot.
package lark

import (
	"sync/atomic"
	"time"
)

const (
	// ChatBot should be created with NewChatBot
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
	// auth info
	appID     string
	appSecret string
	// access token
	tenantAccessToken atomic.Value
	autoRenew         bool
	userAccessToken   atomic.Value
	// user id type for api chat
	userIDType string
	// webhook for NotificationBot
	webhook string
	// API Domain
	domain string
	// http client
	client HTTPClient
	// logger
	logger LogWrapper
	debug  bool
}

// Domains
const (
	DomainFeishu = "https://open.feishu.cn"
	DomainLark   = "https://open.larksuite.com"
)

// TenantAccessToken .
type TenantAccessToken struct {
	TenantAccessToken string
	Expire            time.Duration
	LastUpdatedAt     *time.Time
	EstimatedExpireAt *time.Time
}

// NewChatBot with appID and appSecret
func NewChatBot(appID, appSecret string) *Bot {
	bot := &Bot{
		botType:   ChatBot,
		appID:     appID,
		appSecret: appSecret,
		client:    newDefaultClient(),
		domain:    DomainFeishu,
		logger:    initDefaultLogger(),
	}
	bot.autoRenew = true
	bot.tenantAccessToken.Store(TenantAccessToken{})

	return bot
}

// NewNotificationBot with URL
func NewNotificationBot(hookURL string) *Bot {
	bot := &Bot{
		botType: NotificationBot,
		webhook: hookURL,
		client:  newDefaultClient(),
		logger:  initDefaultLogger(),
	}
	bot.tenantAccessToken.Store(TenantAccessToken{})

	return bot
}

// requireType checks whether the action is allowed in a list of bot types
func (bot *Bot) requireType(botType ...int) bool {
	for _, iterType := range botType {
		if bot.botType == iterType {
			return true
		}
	}
	return false
}

// SetClient assigns a new client to bot.client
func (bot *Bot) SetClient(c HTTPClient) {
	bot.client = c
}

// SetDomain sets domain of endpoint, so we could call Feishu/Lark
// go-lark does not check your host, just use the right one or fail.
func (bot *Bot) SetDomain(domain string) {
	bot.domain = domain
}

// Domain returns current domain
func (bot *Bot) Domain() string {
	return bot.domain
}

// AppID returns bot.appID for external use
func (bot *Bot) AppID() string {
	return bot.appID
}

// BotType returns bot.botType for external use
func (bot *Bot) BotType() int {
	return bot.botType
}

// TenantAccessToken returns tenant access token for external use
func (bot *Bot) TenantAccessToken() string {
	token := bot.tenantAccessToken.Load().(TenantAccessToken)
	return token.TenantAccessToken
}

// SetTenantAccessToken sets tenant access token
func (bot *Bot) SetTenantAccessToken(t TenantAccessToken) {
	bot.tenantAccessToken.Store(t)
}

// SetAutoRenew sets autoRenew
func (bot *Bot) SetAutoRenew(onOff bool) {
	bot.autoRenew = onOff
}

// SetWebhook sets webhook URL
func (bot *Bot) SetWebhook(url string) {
	bot.webhook = url
}
