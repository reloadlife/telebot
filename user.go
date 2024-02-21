package telebot

import (
	"fmt"
)

// User represents a Telegram user or bot.
type User struct {
	// ID is the unique identifier for this user or bot.
	ID int64 `json:"id"`

	// IsBot indicates whether this user is a bot.
	IsBot bool `json:"is_bot"`

	// FirstName is the user's or bot's first name.
	FirstName string `json:"first_name"`

	// LastName is the optional last name of the user or bot.
	LastName *string `json:"last_name,omitempty"`

	// Username is the optional username of the user or bot.
	Username *string `json:"username,omitempty"`

	// LanguageCode is the optional IETF language tag of the user's language.
	LanguageCode *string `json:"language_code,omitempty"`

	// IsPremium is optional and indicates whether this user is a Telegram Premium user.
	IsPremium *bool `json:"is_premium,omitempty"`

	// AddedToAttachmentMenu is optional and indicates whether this user added
	// the bot to the attachment menu.
	AddedToAttachmentMenu *bool `json:"added_to_attachment_menu,omitempty"`

	// CanJoinGroups is optional and indicates whether the bot can be invited to groups.
	// This field is returned only in getMe.
	CanJoinGroups *bool `json:"can_join_groups,omitempty"`

	// CanReadAllGroupMessages is optional and indicates whether privacy mode
	// is disabled for the bot. This field is returned only in getMe.
	CanReadAllGroupMessages *bool `json:"can_read_all_group_messages,omitempty"`

	// SupportsInlineQueries is optional and indicates whether the bot supports
	// inline queries. This field is returned only in getMe.
	SupportsInlineQueries *bool `json:"supports_inline_queries,omitempty"`
}

func (u *User) Type() string {
	if u.IsBot {
		return "Bot"
	}
	return "User"
}
func (u *User) ReflectType() string {
	return fmt.Sprintf("%T", u)
}
