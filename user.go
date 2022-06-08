package lark

// UID types
const (
	UIDEmail   = "email"
	UIDUserID  = "user_id"
	UIDOpenID  = "open_id"
	UIDChatID  = "chat_id"
	UIDUnionID = "union_id"
)

// OptionalUserID to contain openID, chatID, userID, email
type OptionalUserID struct {
	UIDType string
	RealID  string
}

func withOneID(uidType, realID string) *OptionalUserID {
	return &OptionalUserID{
		UIDType: uidType,
		RealID:  realID,
	}
}

// WithEmail uses email as userID
func WithEmail(email string) *OptionalUserID {
	return withOneID("email", email)
}

// WithUserID uses userID as userID
func WithUserID(userID string) *OptionalUserID {
	return withOneID("user_id", userID)
}

// WithOpenID uses openID as userID
func WithOpenID(openID string) *OptionalUserID {
	return withOneID("open_id", openID)
}

// WithChatID uses chatID as userID
func WithChatID(chatID string) *OptionalUserID {
	return withOneID("chat_id", chatID)
}
