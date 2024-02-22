package telebot

import "fmt"

// UsersShared contains information about the users whose identifiers were shared with the bot
// using a KeyboardButtonRequestUsers button.
type UsersShared struct {
	// RequestID is the identifier of the request.
	RequestID int `json:"request_id"`

	// UserIDs is an array of identifiers of the shared users.
	// These numbers may have more than 32 significant bits, and some programming languages may have difficulty/silent defects in interpreting them.
	// But they have at most 52 significant bits, so 64-bit integers or double-precision float types are safe for storing these identifiers.
	// The bot may not have access to the users and could be unable to use these identifiers unless the users are already known to the bot by some other means.
	UserIDs []int64 `json:"user_ids"`
}

func (c *UsersShared) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *UsersShared) Type() string        { return "UsersShared" }
