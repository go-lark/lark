package lark

import "fmt"

const (
	messageURL                = "/open-apis/im/v1/messages?receive_id_type=%s"
	replyMessageURL           = "/open-apis/im/v1/messages/%s/reply"
	reactionsMessageUrl       = "/open-apis/im/v1/messages/%s/reactions"
	deleteReactionsMessageUrl = "/open-apis/im/v1/messages/%s/reactions/%s"
	getMessageURL             = "/open-apis/im/v1/messages/%s"
	updateMessageURL          = "/open-apis/im/v1/messages/%s"
	recallMessageURL          = "/open-apis/im/v1/messages/%s"
	messageReceiptURL         = "/open-apis/message/v4/read_info/"
	ephemeralMessageURL       = "/open-apis/ephemeral/v1/send"
	deleteEphemeralMessageURL = "/open-apis/ephemeral/v1/delete"
	pinMessageURL             = "/open-apis/im/v1/pins"
	unpinMessageURL           = "/open-apis/im/v1/pins/%s"
)

type EmojiType string

const (
	EmojiTypeOK                       EmojiType = "OK"
	EmojiTypeTHUMBSUP                 EmojiType = "THUMBSUP"
	EmojiTypeTHANKS                   EmojiType = "THANKS"
	EmojiTypeMUSCLE                   EmojiType = "MUSCLE"
	EmojiTypeFINGERHEART              EmojiType = "FINGERHEART"
	EmojiTypeAPPLAUSE                 EmojiType = "APPLAUSE"
	EmojiTypeFISTBUMP                 EmojiType = "FISTBUMP"
	EmojiTypeJIAYI                    EmojiType = "JIAYI"
	EmojiTypeDONE                     EmojiType = "DONE"
	EmojiTypeSMILE                    EmojiType = "SMILE"
	EmojiTypeBLUSH                    EmojiType = "BLUSH"
	EmojiTypeLAUGH                    EmojiType = "LAUGH"
	EmojiTypeSMIRK                    EmojiType = "SMIRK"
	EmojiTypeLOL                      EmojiType = "LOL"
	EmojiTypeFACEPALM                 EmojiType = "FACEPALM"
	EmojiTypeLOVE                     EmojiType = "LOVE"
	EmojiTypeWINK                     EmojiType = "WINK"
	EmojiTypePROUD                    EmojiType = "PROUD"
	EmojiTypeWITTY                    EmojiType = "WITTY"
	EmojiTypeSMART                    EmojiType = "SMART"
	EmojiTypeSCOWL                    EmojiType = "SCOWL"
	EmojiTypeTHINKING                 EmojiType = "THINKING"
	EmojiTypeSOB                      EmojiType = "SOB"
	EmojiTypeCRY                      EmojiType = "CRY"
	EmojiTypeERROR                    EmojiType = "ERROR"
	EmojiTypeNOSEPICK                 EmojiType = "NOSEPICK"
	EmojiTypeHAUGHTY                  EmojiType = "HAUGHTY"
	EmojiTypeSLAP                     EmojiType = "SLAP"
	EmojiTypeSPITBLOOD                EmojiType = "SPITBLOOD"
	EmojiTypeTOASTED                  EmojiType = "TOASTED"
	EmojiTypeGLANCE                   EmojiType = "GLANCE"
	EmojiTypeDULL                     EmojiType = "DULL"
	EmojiTypeINNOCENTSMILE            EmojiType = "INNOCENTSMILE"
	EmojiTypeJOYFUL                   EmojiType = "JOYFUL"
	EmojiTypeWOW                      EmojiType = "WOW"
	EmojiTypeTRICK                    EmojiType = "TRICK"
	EmojiTypeYEAH                     EmojiType = "YEAH"
	EmojiTypeENOUGH                   EmojiType = "ENOUGH"
	EmojiTypeTEARS                    EmojiType = "TEARS"
	EmojiTypeEMBARRASSED              EmojiType = "EMBARRASSED"
	EmojiTypeKISS                     EmojiType = "KISS"
	EmojiTypeSMOOCH                   EmojiType = "SMOOCH"
	EmojiTypeDROOL                    EmojiType = "DROOL"
	EmojiTypeOBSESSED                 EmojiType = "OBSESSED"
	EmojiTypeMONEY                    EmojiType = "MONEY"
	EmojiTypeTEASE                    EmojiType = "TEASE"
	EmojiTypeSHOWOFF                  EmojiType = "SHOWOFF"
	EmojiTypeCOMFORT                  EmojiType = "COMFORT"
	EmojiTypeCLAP                     EmojiType = "CLAP"
	EmojiTypePRAISE                   EmojiType = "PRAISE"
	EmojiTypeSTRIVE                   EmojiType = "STRIVE"
	EmojiTypeXBLUSH                   EmojiType = "XBLUSH"
	EmojiTypeSILENT                   EmojiType = "SILENT"
	EmojiTypeWAVE                     EmojiType = "WAVE"
	EmojiTypeWHAT                     EmojiType = "WHAT"
	EmojiTypeFROWN                    EmojiType = "FROWN"
	EmojiTypeSHY                      EmojiType = "SHY"
	EmojiTypeDIZZY                    EmojiType = "DIZZY"
	EmojiTypeLOOKDOWN                 EmojiType = "LOOKDOWN"
	EmojiTypeCHUCKLE                  EmojiType = "CHUCKLE"
	EmojiTypeWAIL                     EmojiType = "WAIL"
	EmojiTypeCRAZY                    EmojiType = "CRAZY"
	EmojiTypeWHIMPER                  EmojiType = "WHIMPER"
	EmojiTypeHUG                      EmojiType = "HUG"
	EmojiTypeBLUBBER                  EmojiType = "BLUBBER"
	EmojiTypeWRONGED                  EmojiType = "WRONGED"
	EmojiTypeHUSKY                    EmojiType = "HUSKY"
	EmojiTypeSHHH                     EmojiType = "SHHH"
	EmojiTypeSMUG                     EmojiType = "SMUG"
	EmojiTypeANGRY                    EmojiType = "ANGRY"
	EmojiTypeHAMMER                   EmojiType = "HAMMER"
	EmojiTypeSHOCKED                  EmojiType = "SHOCKED"
	EmojiTypeTERROR                   EmojiType = "TERROR"
	EmojiTypePETRIFIED                EmojiType = "PETRIFIED"
	EmojiTypeSKULL                    EmojiType = "SKULL"
	EmojiTypeSWEAT                    EmojiType = "SWEAT"
	EmojiTypeSPEECHLESS               EmojiType = "SPEECHLESS"
	EmojiTypeSLEEP                    EmojiType = "SLEEP"
	EmojiTypeDROWSY                   EmojiType = "DROWSY"
	EmojiTypeYAWN                     EmojiType = "YAWN"
	EmojiTypeSICK                     EmojiType = "SICK"
	EmojiTypePUKE                     EmojiType = "PUKE"
	EmojiTypeBETRAYED                 EmojiType = "BETRAYED"
	EmojiTypeHEADSET                  EmojiType = "HEADSET"
	EmojiTypeEatingFood               EmojiType = "EatingFood"
	EmojiTypeMeMeMe                   EmojiType = "MeMeMe"
	EmojiTypeSigh                     EmojiType = "Sigh"
	EmojiTypeTyping                   EmojiType = "Typing"
	EmojiTypeLemon                    EmojiType = "Lemon"
	EmojiTypeGet                      EmojiType = "Get"
	EmojiTypeLGTM                     EmojiType = "LGTM"
	EmojiTypeOnIt                     EmojiType = "OnIt"
	EmojiTypeOneSecond                EmojiType = "OneSecond"
	EmojiTypeVRHeadset                EmojiType = "VRHeadset"
	EmojiTypeYouAreTheBest            EmojiType = "YouAreTheBest"
	EmojiTypeSALUTE                   EmojiType = "SALUTE"
	EmojiTypeSHAKE                    EmojiType = "SHAKE"
	EmojiTypeHIGHFIVE                 EmojiType = "HIGHFIVE"
	EmojiTypeUPPERLEFT                EmojiType = "UPPERLEFT"
	EmojiTypeThumbsDown               EmojiType = "ThumbsDown"
	EmojiTypeSLIGHT                   EmojiType = "SLIGHT"
	EmojiTypeTONGUE                   EmojiType = "TONGUE"
	EmojiTypeEYESCLOSED               EmojiType = "EYESCLOSED"
	EmojiTypeRoarForYou               EmojiType = "RoarForYou"
	EmojiTypeCALF                     EmojiType = "CALF"
	EmojiTypeBEAR                     EmojiType = "BEAR"
	EmojiTypeBULL                     EmojiType = "BULL"
	EmojiTypeRAINBOWPUKE              EmojiType = "RAINBOWPUKE"
	EmojiTypeROSE                     EmojiType = "ROSE"
	EmojiTypeHEART                    EmojiType = "HEART"
	EmojiTypePARTY                    EmojiType = "PARTY"
	EmojiTypeLIPS                     EmojiType = "LIPS"
	EmojiTypeBEER                     EmojiType = "BEER"
	EmojiTypeCAKE                     EmojiType = "CAKE"
	EmojiTypeGIFT                     EmojiType = "GIFT"
	EmojiTypeCUCUMBER                 EmojiType = "CUCUMBER"
	EmojiTypeDrumstick                EmojiType = "Drumstick"
	EmojiTypePepper                   EmojiType = "Pepper"
	EmojiTypeCANDIEDHAWS              EmojiType = "CANDIEDHAWS"
	EmojiTypeBubbleTea                EmojiType = "BubbleTea"
	EmojiTypeCoffee                   EmojiType = "Coffee"
	EmojiTypeYes                      EmojiType = "Yes"
	EmojiTypeNo                       EmojiType = "No"
	EmojiTypeOKR                      EmojiType = "OKR"
	EmojiTypeCheckMark                EmojiType = "CheckMark"
	EmojiTypeCrossMark                EmojiType = "CrossMark"
	EmojiTypeMinusOne                 EmojiType = "MinusOne"
	EmojiTypeHundred                  EmojiType = "Hundred"
	EmojiTypeAWESOMEN                 EmojiType = "AWESOMEN"
	EmojiTypePin                      EmojiType = "Pin"
	EmojiTypeAlarm                    EmojiType = "Alarm"
	EmojiTypeLoudspeaker              EmojiType = "Loudspeaker"
	EmojiTypeTrophy                   EmojiType = "Trophy"
	EmojiTypeFire                     EmojiType = "Fire"
	EmojiTypeBOMB                     EmojiType = "BOMB"
	EmojiTypeMusic                    EmojiType = "Music"
	EmojiTypeXmasTree                 EmojiType = "XmasTree"
	EmojiTypeSnowman                  EmojiType = "Snowman"
	EmojiTypeXmasHat                  EmojiType = "XmasHat"
	EmojiTypeFIREWORKS                EmojiType = "FIREWORKS"
	EmojiType2022                     EmojiType = "2022"
	EmojiTypeREDPACKET                EmojiType = "REDPACKET"
	EmojiTypeFORTUNE                  EmojiType = "FORTUNE"
	EmojiTypeLUCK                     EmojiType = "LUCK"
	EmojiTypeFIRECRACKER              EmojiType = "FIRECRACKER"
	EmojiTypeStickyRiceBalls          EmojiType = "StickyRiceBalls"
	EmojiTypeHEARTBROKEN              EmojiType = "HEARTBROKEN"
	EmojiTypePOOP                     EmojiType = "POOP"
	EmojiTypeStatusFlashOfInspiration EmojiType = "StatusFlashOfInspiration"
	EmojiType18X                      EmojiType = "18X"
	EmojiTypeCLEAVER                  EmojiType = "CLEAVER"
	EmojiTypeSoccer                   EmojiType = "Soccer"
	EmojiTypeBasketball               EmojiType = "Basketball"
	EmojiTypeGeneralDoNotDisturb      EmojiType = "GeneralDoNotDisturb"
	EmojiTypeStatusPrivateMessage     EmojiType = "Status_PrivateMessage"
	EmojiTypeGeneralInMeetingBusy     EmojiType = "GeneralInMeetingBusy"
	EmojiTypeStatusReading            EmojiType = "StatusReading"
	EmojiTypeStatusInFlight           EmojiType = "StatusInFlight"
	EmojiTypeGeneralBusinessTrip      EmojiType = "GeneralBusinessTrip"
	EmojiTypeGeneralWorkFromHome      EmojiType = "GeneralWorkFromHome"
	EmojiTypeStatusEnjoyLife          EmojiType = "StatusEnjoyLife"
	EmojiTypeGeneralTravellingCar     EmojiType = "GeneralTravellingCar"
	EmojiTypeStatusBus                EmojiType = "StatusBus"
	EmojiTypeGeneralSun               EmojiType = "GeneralSun"
	EmojiTypeGeneralMoonRest          EmojiType = "GeneralMoonRest"
	EmojiTypePursueUltimate           EmojiType = "PursueUltimate"
	EmojiTypePatient                  EmojiType = "Patient"
	EmojiTypeAmbitious                EmojiType = "Ambitious"
	EmojiTypeCustomerSuccess          EmojiType = "CustomerSuccess"
	EmojiTypeResponsible              EmojiType = "Responsible"
	EmojiTypeReliable                 EmojiType = "Reliable"
)

// PostMessageResponse .
type PostMessageResponse struct {
	BaseResponse

	Data IMMessage `json:"data"`
}

// IMMessageRequest .
type IMMessageRequest struct {
	ReceiveID string `json:"receive_id"`
	Content   string `json:"content"`
	MsgType   string `json:"msg_type"`
	UUID      string `json:"uuid,omitempty"`
}

// IMSender .
type IMSender struct {
	ID         string `json:"id"`
	IDType     string `json:"id_type"`
	SenderType string `json:"sender_type"`
	TenantKey  string `json:"tenant_key"`
}

// IMMention .
type IMMention struct {
	ID     string `json:"id"`
	IDType string `json:"id_type"`
	Key    string `json:"key"`
	Name   string `json:"name"`
}

// IMBody .
type IMBody struct {
	Content string `json:"content"`
}

// IMMessage .
type IMMessage struct {
	MessageID      string `json:"message_id"`
	UpperMessageID string `json:"upper_message_id"`
	RootID         string `json:"root_id"`
	ParentID       string `json:"parent_id"`
	ChatID         string `json:"chat_id"`
	MsgType        string `json:"msg_type"`
	CreateTime     string `json:"create_time"`
	UpdateTime     string `json:"update_time"`
	Deleted        bool   `json:"deleted"`
	Updated        bool   `json:"updated"`
	Sender         IMSender
	Mentions       []IMMention
	Body           IMBody
}

type ReactionResponse struct {
	BaseResponse
	Data struct {
		ReactionID string `json:"reaction_id"`
		Operator   struct {
			OperatorID   string `json:"operator_id"`
			OperatorType string `json:"operator_type"`
			ActionTime   string `json:"action_time"`
		} `json:"operator"`
		ReactionType struct {
			EmojiType EmojiType `json:"emoji_type"`
		} `json:"reaction_type"`
	} `json:"data"`
}

// GetMessageResponse .
type GetMessageResponse struct {
	BaseResponse

	Data struct {
		Items []IMMessage `json:"items"`
	} `json:"data"`
}

// PostEphemeralMessageResponse .
type PostEphemeralMessageResponse struct {
	BaseResponse
	Data struct {
		MessageID string `json:"message_id"`
	} `json:"data"`
}

// DeleteEphemeralMessageResponse .
type DeleteEphemeralMessageResponse = BaseResponse

// RecallMessageResponse .
type RecallMessageResponse = BaseResponse

// UpdateMessageResponse .
type UpdateMessageResponse = BaseResponse

// MessageReceiptResponse .
type MessageReceiptResponse struct {
	BaseResponse
	Data struct {
		ReadUsers []struct {
			OpenID    string `json:"open_id"`
			UserID    string `json:"user_id"`
			Timestamp string `json:"timestamp"`
		} `json:"read_users"`
	} `json:"data"`
}

// PinMessageResponse .
type PinMessageResponse struct {
	BaseResponse
	Data struct {
		Pin struct {
			MessageID      string `json:"message_id"`
			ChatID         string `json:"chat_id"`
			OperatorID     string `json:"operator_id"`
			OperatorIDType string `json:"operator_id_type"`
			CreateTime     string `json:"create_time"`
		} `json:"pin"`
	} `json:"data"`
}

// UnpinMessageResponse .
type UnpinMessageResponse = BaseResponse

func newMsgBufWithOptionalUserID(msgType string, userID *OptionalUserID) *MsgBuffer {
	mb := NewMsgBuffer(msgType)
	realID := userID.RealID
	switch userID.UIDType {
	case "email":
		mb.BindEmail(realID)
	case "open_id":
		mb.BindOpenID(realID)
	case "chat_id":
		mb.BindChatID(realID)
	case "user_id":
		mb.BindUserID(realID)
	case "union_id":
		mb.BindUnionID(realID)
	default:
		return nil
	}
	return mb
}

// PostText is a simple way to send text messages
func (bot Bot) PostText(text string, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgText, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	om := mb.Text(text).Build()
	return bot.PostMessage(om)
}

// PostRichText is a simple way to send rich text messages
func (bot Bot) PostRichText(postContent *PostContent, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgPost, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	om := mb.Post(postContent).Build()
	return bot.PostMessage(om)
}

// PostTextMention is a simple way to send text messages with @user
func (bot Bot) PostTextMention(text string, atUserID string, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgText, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	tb := NewTextBuilder()
	om := mb.Text(tb.Text(text).Mention(atUserID).Render()).Build()
	return bot.PostMessage(om)
}

// PostTextMentionAll is a simple way to send text messages with @all
func (bot Bot) PostTextMentionAll(text string, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgText, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	tb := NewTextBuilder()
	om := mb.Text(tb.Text(text).MentionAll().Render()).Build()
	return bot.PostMessage(om)
}

// PostTextMentionAndReply is a simple way to send text messages with @user and reply a message
func (bot Bot) PostTextMentionAndReply(text string, atUserID string, userID *OptionalUserID, replyID string) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgText, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	tb := NewTextBuilder()
	om := mb.Text(tb.Text(text).Mention(atUserID).Render()).BindReply(replyID).Build()
	return bot.PostMessage(om)
}

// PostImage is a simple way to send image
func (bot Bot) PostImage(imageKey string, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgImage, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	om := mb.Image(imageKey).Build()
	return bot.PostMessage(om)
}

// PostShareChat is a simple way to share chat
func (bot Bot) PostShareChat(chatID string, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgShareCard, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	om := mb.ShareChat(chatID).Build()
	return bot.PostMessage(om)
}

// PostShareUser is a simple way to share user
func (bot Bot) PostShareUser(openID string, userID *OptionalUserID) (*PostMessageResponse, error) {
	mb := newMsgBufWithOptionalUserID(MsgShareUser, userID)
	if mb == nil {
		return nil, ErrParamUserID
	}
	om := mb.ShareUser(openID).Build()
	return bot.PostMessage(om)
}

// PostMessage posts message
func (bot Bot) PostMessage(om OutcomingMessage) (*PostMessageResponse, error) {
	req, err := BuildMessage(om)
	if err != nil {
		return nil, err
	}
	var respData PostMessageResponse
	if om.RootID == "" {
		err = bot.PostAPIRequest("PostMessage", fmt.Sprintf(messageURL, om.UIDType), true, req, &respData)
	} else {
		resp, err := bot.ReplyMessage(om)
		return resp, err
	}
	return &respData, err
}

// ReplyMessage replies messages
func (bot Bot) ReplyMessage(om OutcomingMessage) (*PostMessageResponse, error) {
	req, err := BuildMessage(om)
	if err != nil {
		return nil, err
	}
	if om.RootID == "" {
		return nil, ErrParamMessageID
	}
	var respData PostMessageResponse
	err = bot.PostAPIRequest("ReplyMessage", fmt.Sprintf(replyMessageURL, om.RootID), true, req, &respData)
	return &respData, err
}

// ReactionMessage reactions messages
func (bot Bot) ReactionMessage(messageID string, emojiType EmojiType) (*ReactionResponse, error) {
	req := map[string]interface{}{
		"reaction_type": map[string]interface{}{
			"emoji_type": emojiType,
		},
	}
	var respData ReactionResponse
	err := bot.PostAPIRequest("ReactionMessage", fmt.Sprintf(reactionsMessageUrl, messageID), true, req, &respData)
	return &respData, err
}

// DeleteReactionMessage delete reactions messages
func (bot Bot) DeleteReactionMessage(messageID string, reactionID string) (*ReactionResponse, error) {
	var respData ReactionResponse
	err := bot.DeleteAPIRequest("DeleteReactionMessage", fmt.Sprintf(deleteReactionsMessageUrl, messageID, reactionID), true, nil, &respData)
	return &respData, err
}

// UpdateMessage update message card
func (bot Bot) UpdateMessage(messageID string, om OutcomingMessage) (*UpdateMessageResponse, error) {
	if om.MsgType != MsgInteractive {
		return nil, ErrMessageType
	}
	req, err := BuildMessage(om)
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf(updateMessageURL, messageID)
	var respData UpdateMessageResponse
	err = bot.PatchAPIRequest("UpdateMessage", url, true, req, &respData)
	return &respData, err
}

// GetMessage posts message with im/v1
func (bot Bot) GetMessage(messageID string) (*GetMessageResponse, error) {
	var respData GetMessageResponse
	err := bot.GetAPIRequest("GetMessage", fmt.Sprintf(getMessageURL, messageID), true, nil, &respData)
	return &respData, err
}

// RecallMessage recalls a message with ID
func (bot Bot) RecallMessage(messageID string) (*RecallMessageResponse, error) {
	url := fmt.Sprintf(recallMessageURL, messageID)
	var respData RecallMessageResponse
	err := bot.DeleteAPIRequest("RecallMessage", url, true, nil, &respData)
	return &respData, err
}

// MessageReadReceipt queries message read receipt
func (bot Bot) MessageReadReceipt(messageID string) (*MessageReceiptResponse, error) {
	params := map[string]interface{}{
		"message_id": messageID,
	}
	var respData MessageReceiptResponse
	err := bot.PostAPIRequest("MessageReadReceipt", messageReceiptURL, true, params, &respData)
	return &respData, err
}

// PostEphemeralMessage posts an ephemeral message
func (bot Bot) PostEphemeralMessage(om OutcomingMessage) (*PostEphemeralMessageResponse, error) {
	if om.UIDType == UIDUnionID {
		return nil, ErrUnsupportedUIDType
	}
	params := BuildOutcomingMessageReq(om)
	var respData PostEphemeralMessageResponse
	err := bot.PostAPIRequest("PostEphemeralMessage", ephemeralMessageURL, true, params, &respData)
	return &respData, err
}

// DeleteEphemeralMessage deletes an ephemeral message
func (bot Bot) DeleteEphemeralMessage(messageID string) (*DeleteEphemeralMessageResponse, error) {
	params := map[string]interface{}{
		"message_id": messageID,
	}
	var respData DeleteEphemeralMessageResponse
	err := bot.PostAPIRequest("DeleteEphemeralMessage", deleteEphemeralMessageURL, true, params, &respData)
	return &respData, err
}

// PinMessage pin a message
func (bot Bot) PinMessage(messageID string) (*PinMessageResponse, error) {
	params := map[string]interface{}{
		"message_id": messageID,
	}
	var respData PinMessageResponse
	err := bot.PostAPIRequest("PinMessage", pinMessageURL, true, params, &respData)
	return &respData, err
}

// UnpinMessage unpin a message
func (bot Bot) UnpinMessage(messageID string) (*UnpinMessageResponse, error) {
	url := fmt.Sprintf(unpinMessageURL, messageID)
	var respData UnpinMessageResponse
	err := bot.DeleteAPIRequest("PinMessage", url, true, nil, &respData)
	return &respData, err
}
