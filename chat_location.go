package telebot

import "fmt"

// ChatLocation represents a location to which a chat is connected.
type ChatLocation struct {
	Location Location `json:"location,omitempty"`
	Address  string   `json:"address,omitempty"`
}

func (c *ChatLocation) Type() string        { return "chat_location" }
func (c *ChatLocation) ReflectType() string { return fmt.Sprintf("%T", c) }
