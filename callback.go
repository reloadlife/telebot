package telebot

import (
	"encoding/json"
	"fmt"
)

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

// MarshalJSON to be JSON serializable, but only include non-empty fields.
func (c *Callback) MarshalJSON() ([]byte, error) {
	return json.Marshal(c)
}

// UnmarshalJSON to be JSON unserializable
func (c *Callback) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, c)
}

// String returns a string representation of this user.
func (c *Callback) String() string {
	indented, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("Callback{ID: %v, Data: %v (%s)}\n%s\n", c.ID, c.Data, c.Unique, indented)
}

type CallbackEndpoint interface {
	CallbackUnique() string
}
