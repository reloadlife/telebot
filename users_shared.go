package telebot

import "fmt"

type SharedUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`

	Photo PhotoSizes `json:"photo"`
}

// UsersShared contains information about the users whose identifiers were shared with the bot
// using a KeyboardButtonRequestUsers button.
type UsersShared struct {
	// RequestID is the identifier of the request.
	RequestID int `json:"request_id"`

	// Users Information about users shared with the bot.
	Users []SharedUser `json:"user_ids"`
}

func (c *UsersShared) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *UsersShared) Type() string        { return "UsersShared" }
