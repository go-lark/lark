package lark

import "fmt"

const (
	getChatURL    = "/open-apis/im/v1/chats/%s?user_id_type=%s"
	updateChatURL = "/open-apis/im/v1/chats/%s?user_id_type=%s"
	createChatURL = "/open-apis/im/v1/chats?user_id_type=%s"
	deleteChatURL = "/open-apis/im/v1/chats/%s"
)

// GetChatResponse .
type GetChatResponse struct {
	BaseResponse

	Data ChatInfo `json:"data"`
}

// ChatInfo entity of a chat, not every field is available for every API.
type ChatInfo struct {
	ChatID      string `json:"chat_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	Description string `json:"description,omitempty"`
	I18NNames   struct {
		ZhCN string `json:"zh_cn,omitempty"`
		EnUS string `json:"en_us,omitempty"`
		JaJP string `json:"ja_jp,omitempty"`
	} `json:"i18n_names,omitempty"`
	AddMemberPermission    string `json:"add_member_permission,omitempty"`
	ShareCardPermission    string `json:"share_card_permission,omitempty"`
	AtAllPermission        string `json:"at_all_permission,omitempty"`
	EditPermission         string `json:"edit_permission,omitempty"`
	OwnerIDType            string `json:"owner_id_type,omitempty"`
	OwnerID                string `json:"owner_id,omitempty"`
	ChatMode               string `json:"chat_mode,omitempty"`
	ChatType               string `json:"chat_type,omitempty"`
	ChatTag                string `json:"chat_tag,omitempty"`
	JoinMessageVisibility  string `json:"join_message_visibility,omitempty"`
	LeaveMessageVisibility string `json:"leave_message_visibility,omitempty"`
	MembershipApproval     string `json:"membership_approval,omitempty"`
	ModerationPermission   string `json:"moderation_permission,omitempty"`
	External               bool   `json:"external,omitempty"`
}

// CreateChatRequest .
type CreateChatRequest struct {
	Name        string `json:"name,omitempty"`
	Avatar      string `json:"Avatar,omitempty"`
	Description string `json:"description,omitempty"`
	I18NNames   struct {
		ZhCN string `json:"zh_cn,omitempty"`
		EnUS string `json:"en_us,omitempty"`
		JaJP string `json:"ja_jp,omitempty"`
	} `json:"i18n_names,omitempty"`
	OwnerID                string `json:"owner_id,omitempty"`
	ChatMode               string `json:"chat_mode,omitempty"`
	ChatType               string `json:"chat_type,omitempty"`
	JoinMessageVisibility  string `json:"join_message_visibility,omitempty"`
	LeaveMessageVisibility string `json:"leave_message_visibility,omitempty"`
	MembershipApproval     string `json:"membership_approval,omitempty"`
	External               bool   `json:"external,omitempty"`
}

// CreateChatResponse .
type CreateChatResponse struct {
	BaseResponse

	Data ChatInfo `json:"data"`
}

// DeleteChatResponse .
type DeleteChatResponse struct {
	BaseResponse
}

// UpdateChatRequest .
type UpdateChatRequest struct {
	Name        string `json:"name,omitempty"`
	Avatar      string `json:"Avatar,omitempty"`
	Description string `json:"description,omitempty"`
	I18NNames   struct {
		ZhCN string `json:"zh_cn,omitempty"`
		EnUS string `json:"en_us,omitempty"`
		JaJP string `json:"ja_jp,omitempty"`
	} `json:"i18n_names,omitempty"`
	AddMemberPermission    string `json:"add_member_permission,omitempty"`
	ShareCardPermission    string `json:"share_card_permission,omitempty"`
	AtAllPermission        string `json:"at_all_permission,omitempty"`
	EditPermission         string `json:"edit_permission,omitempty"`
	OwnerID                string `json:"owner_id,omitempty"`
	JoinMessageVisibility  string `json:"join_message_visibility,omitempty"`
	LeaveMessageVisibility string `json:"leave_message_visibility,omitempty"`
	MembershipApproval     string `json:"membership_approval,omitempty"`
}

// UpdateChatResponse .
type UpdateChatResponse struct {
	BaseResponse
}

// WithUserIDType .
func (bot *Bot) WithUserIDType(userIDType string) *Bot {
	bot.userIDType = userIDType
	return bot
}

// GetChat .
func (bot Bot) GetChat(chatID string) (*GetChatResponse, error) {
	var respData GetChatResponse
	err := bot.GetAPIRequest("GetChatInfo", fmt.Sprintf(getChatURL, chatID, bot.userIDType), true, nil, &respData)
	return &respData, err
}

// CreateChat .
func (bot Bot) CreateChat(req CreateChatRequest) (*CreateChatResponse, error) {
	var respData CreateChatResponse
	err := bot.PostAPIRequest("CreateChat", fmt.Sprintf(createChatURL, bot.userIDType), true, req, &respData)
	return &respData, err
}

// DeleteChat .
func (bot Bot) DeleteChat(chatID string) (*DeleteChatResponse, error) {
	var respData DeleteChatResponse
	err := bot.DeleteAPIRequest("DeleteChat", fmt.Sprintf(deleteChatURL, chatID), true, nil, &respData)
	return &respData, err
}

// UpdateChat .
func (bot Bot) UpdateChat(chatID string, req UpdateChatRequest) (*UpdateChatResponse, error) {
	var respData UpdateChatResponse
	err := bot.PutAPIRequest("UpdateChat", fmt.Sprintf(updateChatURL, chatID, bot.userIDType), true, req, &respData)
	return &respData, err
}
