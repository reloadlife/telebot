package telebot

type BusinessConnection struct {
	// ID  Unique identifier of the business connection
	ID string `json:"id"`

	// User  Business account user that created the business connection
	User User `json:"user"`

	// UserChatID Identifier of a private chat with the user who created the business connection.
	UserChatID int64 `json:"user_chat_id"`

	// Date  Point in time (Unix timestamp) when the business connection was created
	Date int64 `json:"date"`

	// CanReply if the bot can act on behalf of the business account in chats that were active in the last 24 hours
	CanReply bool `json:"can_reply"`

	// Enabled if the connection is active
	Enabled bool `json:"is_enabled"`
}
