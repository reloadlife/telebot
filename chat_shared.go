package telebot

import "fmt"

// ChatShared contains information about the chat whose identifier was shared with the bot
// using a KeyboardButtonRequestChat button.
type ChatShared struct {
	// RequestID is the identifier of the request.
	RequestID int `json:"request_id"`

	// ChatID is the identifier of the shared chat.
	// This number may have more than 32 significant bits, and some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	// The bot may not have access to the chat and could be unable to use this identifier unless the chat is already known to the bot by some other means.
	ChatID int64 `json:"chat_id"`
}

func (c *ChatShared) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *ChatShared) Type() string        { return "ChatShared" }
