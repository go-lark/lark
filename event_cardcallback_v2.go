package lark

type CardV2CallbackEvent struct {
	Operator struct {
		TenantKey string `json:"tenant_key,omitempty"`
		UserID    string `json:"user_id,omitempty"`
		UnionID   string `json:"union_id,omitempty"`
		OpenID    string `json:"open_id,omitempty"`
	} `json:"operator,omitempty"`
	Token  string `json:"token,omitempty"`
	Action struct {
		Value      interface{} `json:"value,omitempty"`
		Tag        string      `json:"tag,omitempty"`
		Timezone   string      `json:"timezone,omitempty"`
		Name       string      `json:"name,omitempty"`
		FormValue  string      `json:"form_value,omitempty"`
		InputValue string      `json:"input_value,omitempty"`
		Option     string      `json:"option,omitempty"`
		Options    []string    `json:"options,omitempty"`
		Checked    bool        `json:"checked,omitempty"`
	} `json:"action,omitempty"`
	Host         string `json:"host,omitempty"`
	DeliveryType string `json:"delivery_type,omitempty"`
	Context      struct {
		Url           string `json:"url,omitempty"`
		PreviewToken  string `json:"preview_token,omitempty"`
		OpenMessageID string `json:"open_message_id,omitempty"`
		OpenChatID    string `json:"open_chat_id,omitempty"`
	} `json:"context,omitempty"`
}
