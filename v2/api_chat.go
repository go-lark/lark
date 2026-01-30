package lark

import (
	"context"
	"fmt"
	"net/url"
)

const (
	getChatURL          = "/open-apis/im/v1/chats/%s?user_id_type=%s"
	listChatURL         = "/open-apis/im/v1/chats?user_id_type=%s&sort_type=%s&page_token=%s&page_size=%d"
	searchChatURL       = "/open-apis/im/v1/chats/search?user_id_type=%s&query=%s&page_token=%s&page_size=%d"
	updateChatURL       = "/open-apis/im/v1/chats/%s?user_id_type=%s"
	createChatURL       = "/open-apis/im/v1/chats?user_id_type=%s"
	deleteChatURL       = "/open-apis/im/v1/chats/%s"
	joinChatURL         = "/open-apis/im/v1/chats/%s/members/me_join"
	addChatMemberURL    = "/open-apis/im/v1/chats/%s/members?member_id_type=%s"
	removeChatMemberURL = "/open-apis/im/v1/chats/%s/members?member_id_type=%s"
	isInChatURL         = "/open-apis/im/v1/chats/%s/members/is_in_chat"
	getChatMembersURL   = "/open-apis/im/v1/chats/%s/members?member_id_type=%s"
	setTopNoticeURL     = "/open-apis/im/v1/chats/%s/top_notice/put_top_notice"
	deleteTopNoticeURL  = "/open-apis/im/v1/chats/%s/top_notice/delete_top_notice"
)

// GetChatResponse .
type GetChatResponse struct {
	BaseResponse

	Data ChatInfo `json:"data"`
}

// ChatInfo is entity of a chat, not every field is available for every API.
type ChatInfo struct {
	ChatID                 string    `json:"chat_id,omitempty"`
	Name                   string    `json:"name,omitempty"`
	Avatar                 string    `json:"avatar,omitempty"`
	Description            string    `json:"description,omitempty"`
	I18NNames              I18NNames `json:"i18n_names,omitempty"`
	AddMemberPermission    string    `json:"add_member_permission,omitempty"`
	ShareCardPermission    string    `json:"share_card_permission,omitempty"`
	AtAllPermission        string    `json:"at_all_permission,omitempty"`
	EditPermission         string    `json:"edit_permission,omitempty"`
	OwnerIDType            string    `json:"owner_id_type,omitempty"`
	OwnerID                string    `json:"owner_id,omitempty"`
	ChatMode               string    `json:"chat_mode,omitempty"`
	ChatType               string    `json:"chat_type,omitempty"`
	ChatTag                string    `json:"chat_tag,omitempty"`
	JoinMessageVisibility  string    `json:"join_message_visibility,omitempty"`
	LeaveMessageVisibility string    `json:"leave_message_visibility,omitempty"`
	MembershipApproval     string    `json:"membership_approval,omitempty"`
	ModerationPermission   string    `json:"moderation_permission,omitempty"`
	External               bool      `json:"external,omitempty"`
}

// ListChatResponse .
type ListChatResponse struct {
	BaseResponse

	Data struct {
		Items     []ChatListInfo `json:"items"`
		PageToken string         `json:"page_token"`
		HasMore   bool           `json:"has_more"`
	} `json:"data"`
}

// ChatListInfo .
type ChatListInfo struct {
	ChatID      string `json:"chat_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Avatar      string `json:"avatar,omitempty"`
	Description string `json:"description,omitempty"`
	OwnerIDType string `json:"owner_id_type,omitempty"`
	OwnerID     string `json:"owner_id,omitempty"`
	External    bool   `json:"external,omitempty"`
	TenantKey   string `json:"tenant_key"`
}

// CreateChatRequest .
type CreateChatRequest struct {
	Name                   string    `json:"name,omitempty"`
	Avatar                 string    `json:"avatar,omitempty"`
	Description            string    `json:"description,omitempty"`
	I18NNames              I18NNames `json:"i18n_names,omitempty"`
	OwnerID                string    `json:"owner_id,omitempty"`
	ChatMode               string    `json:"chat_mode,omitempty"`
	ChatType               string    `json:"chat_type,omitempty"`
	JoinMessageVisibility  string    `json:"join_message_visibility,omitempty"`
	LeaveMessageVisibility string    `json:"leave_message_visibility,omitempty"`
	MembershipApproval     string    `json:"membership_approval,omitempty"`
	External               bool      `json:"external,omitempty"`
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
	Name                   string    `json:"name,omitempty"`
	Avatar                 string    `json:"avatar,omitempty"`
	Description            string    `json:"description,omitempty"`
	I18NNames              I18NNames `json:"i18n_names,omitempty"`
	AddMemberPermission    string    `json:"add_member_permission,omitempty"`
	ShareCardPermission    string    `json:"share_card_permission,omitempty"`
	AtAllPermission        string    `json:"at_all_permission,omitempty"`
	EditPermission         string    `json:"edit_permission,omitempty"`
	OwnerID                string    `json:"owner_id,omitempty"`
	JoinMessageVisibility  string    `json:"join_message_visibility,omitempty"`
	LeaveMessageVisibility string    `json:"leave_message_visibility,omitempty"`
	MembershipApproval     string    `json:"membership_approval,omitempty"`
}

// UpdateChatResponse .
type UpdateChatResponse struct {
	BaseResponse
}

// JoinChatResponse .
type JoinChatResponse struct {
	BaseResponse
}

// AddChatMemberRequest .
type AddChatMemberRequest struct {
	IDList []string `json:"id_list"`
}

// AddChatMemberResponse .
type AddChatMemberResponse struct {
	BaseResponse

	Data struct {
		InvalidIDList    []string `json:"invalid_id_list"`
		NotExistedIDList []string `json:"not_existed_id_list"`
	} `json:"data"`
}

// RemoveChatMemberRequest .
type RemoveChatMemberRequest struct {
	IDList []string `json:"id_list"`
}

// RemoveChatMemberResponse .
type RemoveChatMemberResponse struct {
	BaseResponse

	Data struct {
		InvalidIDList []string `json:"invalid_id_list"`
	} `json:"data"`
}

// IsInChatResponse .
type IsInChatResponse struct {
	BaseResponse

	Data struct {
		IsInChat bool `json:"is_in_chat"`
	} `json:"data"`
}

// GetChatMembersResponse .
type GetChatMembersResponse struct {
	BaseResponse

	Data struct {
		Items       []ChatMember `json:"items"`
		PageToken   string       `json:"page_token"`
		HasMore     bool         `json:"has_more"`
		MemberTotal int          `json:"member_total"`
	} `json:"data"`
}

// ChatMember .
type ChatMember struct {
	MemberIDType string `json:"member_id_type"`
	MemberID     string `json:"member_id"`
	Name         string `json:"name"`
	TenantKey    string `json:"tenant_key"`
}

// SetTopNoticeRequest .
type SetTopNoticeRequest struct {
	ChatTopNotice []ChatTopNoticeAction `json:"chat_top_notice"`
}

// ChatTopNoticeAction .
type ChatTopNoticeAction struct {
	ActionType string `json:"action_type"`
	MessageID  string `json:"message_id"`
}

// SetTopNoticeResponse .
type SetTopNoticeResponse = BaseResponse

// DeleteTopNoticeResponse .
type DeleteTopNoticeResponse = BaseResponse

// GetChat .
func (bot *Bot) GetChat(ctx context.Context, chatID string) (*GetChatResponse, error) {
	var respData GetChatResponse
	err := bot.GetAPIRequest(ctx, "GetChatInfo", fmt.Sprintf(getChatURL, chatID, bot.userIDType), true, nil, &respData)
	return &respData, err
}

// ListChat lists chats
// sortType: ByCreateTimeAsc/ByActiveTimeDesc
func (bot *Bot) ListChat(ctx context.Context, sortType string, pageToken string, pageSize int) (*ListChatResponse, error) {
	var respData ListChatResponse
	if sortType == "" {
		sortType = "ByCreateTimeAsc"
	}
	err := bot.GetAPIRequest(
		ctx,
		"ListChat",
		fmt.Sprintf(listChatURL, bot.userIDType, sortType, pageToken, pageSize),
		true, nil, &respData)
	return &respData, err
}

// SearchChat searches chat
func (bot *Bot) SearchChat(ctx context.Context, query string, pageToken string, pageSize int) (*ListChatResponse, error) {
	var respData ListChatResponse
	err := bot.GetAPIRequest(
		ctx,
		"SearchChat",
		fmt.Sprintf(searchChatURL, bot.userIDType, query, pageToken, pageSize),
		true, nil, &respData)
	return &respData, err
}

// CreateChat .
func (bot *Bot) CreateChat(ctx context.Context, req CreateChatRequest) (*CreateChatResponse, error) {
	var respData CreateChatResponse
	err := bot.PostAPIRequest(ctx, "CreateChat", fmt.Sprintf(createChatURL, bot.userIDType), true, req, &respData)
	return &respData, err
}

// DeleteChat .
func (bot *Bot) DeleteChat(ctx context.Context, chatID string) (*DeleteChatResponse, error) {
	var respData DeleteChatResponse
	err := bot.DeleteAPIRequest(ctx, "DeleteChat", fmt.Sprintf(deleteChatURL, chatID), true, nil, &respData)
	return &respData, err
}

// UpdateChat .
func (bot *Bot) UpdateChat(ctx context.Context, chatID string, req UpdateChatRequest) (*UpdateChatResponse, error) {
	var respData UpdateChatResponse
	err := bot.PutAPIRequest(ctx, "UpdateChat", fmt.Sprintf(updateChatURL, chatID, bot.userIDType), true, req, &respData)
	return &respData, err
}

// JoinChat .
func (bot *Bot) JoinChat(ctx context.Context, chatID string) (*JoinChatResponse, error) {
	var respData JoinChatResponse
	err := bot.PatchAPIRequest(ctx, "JoinChat", fmt.Sprintf(joinChatURL, chatID), true, nil, &respData)
	return &respData, err
}

// AddChatMember .
func (bot *Bot) AddChatMember(ctx context.Context, chatID string, idList []string) (*AddChatMemberResponse, error) {
	var respData AddChatMemberResponse
	req := AddChatMemberRequest{
		IDList: idList,
	}
	err := bot.PostAPIRequest(
		ctx,
		"AddChatMember",
		fmt.Sprintf(addChatMemberURL, chatID, bot.userIDType),
		true, req, &respData)
	return &respData, err
}

// RemoveChatMember .
func (bot *Bot) RemoveChatMember(ctx context.Context, chatID string, idList []string) (*RemoveChatMemberResponse, error) {
	var respData RemoveChatMemberResponse
	req := RemoveChatMemberRequest{
		IDList: idList,
	}
	err := bot.PostAPIRequest(
		ctx,
		"RemoveChatMember",
		fmt.Sprintf(removeChatMemberURL, chatID, bot.userIDType),
		true, req, &respData)
	return &respData, err
}

// IsInChat .
func (bot *Bot) IsInChat(ctx context.Context, chatID string) (*IsInChatResponse, error) {
	var respData IsInChatResponse
	err := bot.GetAPIRequest(ctx, "IsInChat", fmt.Sprintf(isInChatURL, chatID), true, nil, &respData)
	return &respData, err
}

// GetChatMembers .
// NOTICE: pageSize must be larger than 10, e.g. if you present pageSize=1, it returns the same pageToken as pageSize=10. So we recommend you just pass pageSize=10.
func (bot *Bot) GetChatMembers(ctx context.Context, chatID string, pageToken string, pageSize int) (*GetChatMembersResponse, error) {
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 10
	}
	var respData GetChatMembersResponse
	v := url.Values{}
	v.Add("page_size", fmt.Sprint(pageSize))
	if len(pageToken) > 0 {
		v.Add("page_token", pageToken)
	}
	fullURL := fmt.Sprintf(getChatMembersURL, chatID, bot.userIDType) + "&" + v.Encode()
	err := bot.GetAPIRequest(ctx, "GetChatMembers", fullURL, true, nil, &respData)
	return &respData, err
}

// SetTopNotice .
func (bot *Bot) SetTopNotice(ctx context.Context, chatID, actionType, messageID string) (*SetTopNoticeResponse, error) {
	var respData SetTopNoticeResponse
	req := SetTopNoticeRequest{
		ChatTopNotice: []ChatTopNoticeAction{
			{
				ActionType: actionType,
				MessageID:  messageID,
			},
		},
	}
	url := fmt.Sprintf(setTopNoticeURL, chatID)
	err := bot.PostAPIRequest(ctx, "SetTopNotice", url, true, req, &respData)
	return &respData, err
}

// DeleteTopNotice .
func (bot *Bot) DeleteTopNotice(ctx context.Context, chatID string) (*DeleteChatResponse, error) {
	var respData DeleteChatResponse
	url := fmt.Sprintf(deleteTopNoticeURL, chatID)
	err := bot.PostAPIRequest(ctx, "DeleteTopNotice", url, true, nil, &respData)
	return &respData, err
}
