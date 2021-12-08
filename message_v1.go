package lark

// BuildOutcomingMessageReq for msg builder
func BuildOutcomingMessageReq(om OutcomingMessage) map[string]interface{} {
	params := map[string]interface{}{
		"msg_type":     om.MsgType,
		"root_id":      om.RootID,
		"update_multi": om.UpdateMulti,
	}
	params[om.UIDType] = buildReceiveID(om)
	content := make(map[string]interface{})
	if om.Content.Text != nil {
		content["text"] = *&om.Content.Text.Text
	}
	if om.Content.Image != nil {
		content["image_key"] = *&om.Content.Image.ImageKey
	}
	if om.Content.ShareChat != nil {
		content["share_open_chat_id"] = *&om.Content.ShareChat.ChatID
	}
	if om.Content.Post != nil {
		content["post"] = *om.Content.Post
	}
	if om.MsgType == MsgInteractive && om.Content.Card != nil {
		params["card"] = *om.Content.Card
	}
	params["content"] = content
	return params
}
