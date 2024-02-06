package telebot

import (
	"encoding/json"
	"fmt"
	"strconv"
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

// MarshalJSON to be JSON serializable, but only include non-empty fields.
func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(u)
}

// UnmarshalJSON to be JSON unserializable
func (u *User) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, u)
}

// String returns a string representation of this user.
func (u *User) String() string {
	indented, _ := json.MarshalIndent(u, "", "  ")
	isBot := ""
	if u.IsBot {
		isBot = "(Bot)"
	}
	return fmt.Sprintf("User{%sID: %d, User: @%v}\n%s\n", isBot, u.ID, u.Username, indented)
}

// Recipient Mention returns a string which mentions the user.
func (u *User) Recipient() string {
	return strconv.FormatInt(u.ID, 10)
}
