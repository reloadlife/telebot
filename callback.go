package telebot

import "fmt"

type Callback struct {
	ID            string                    `json:"id"`
	Sender        *User                     `json:"from"`
	Message       *MaybeInaccessibleMessage `json:"message,omitempty"`
	MessageID     string                    `json:"inline_message_id"`
	ChatInstance  string                    `json:"chat_instance,omitempty"`
	Data          string                    `json:"data"`
	GameShortName string                    `json:"game_short_name,omitempty"`

	Unique string `json:"-"`
}

func (c *Callback) IsInline() bool {
	return c.MessageID != ""
}

type CallbackEndpoint interface {
	CallbackUnique() string
}

func (c *Callback) ReflectType() string {
	return fmt.Sprintf("%T", c)
}

func (c *Callback) Type() string {
	return "Callback"
}
