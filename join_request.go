package telebot

import (
	"fmt"
	"strconv"
)

// ChatJoinRequest represents a join request sent to a chat.
type ChatJoinRequest struct {
	// Chat is the chat to which the request was sent.
	Chat Chat `json:"chat"`

	// From is the user that sent the join request.
	From User `json:"from"`

	// UserChatID is the identifier of a private chat with the user who sent the join request.
	// This number may have more than 32 significant bits, and some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a 64-bit integer or double-precision float type is safe for storing this identifier.
	// The bot can use this identifier for 5 minutes to send messages until the join request is processed, assuming no other administrator contacted the user.
	UserChatID int64 `json:"user_chat_id"`

	// Date is the date the request was sent in Unix time.
	Date int64 `json:"date"`

	// Bio is the bio of the user (optional).
	Bio string `json:"bio,omitempty"`

	// InviteLink is the chat invite link that was used by the user to send the join request (optional).
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

func (c *ChatJoinRequest) Recipient() string {
	return strconv.FormatInt(c.Chat.ID, 10)
}

func (c *ChatJoinRequest) ReflectType() string {
	return fmt.Sprintf("%T", c)
}

func (c *ChatJoinRequest) Type() string {
	return "ChatJoinRequest"
}
