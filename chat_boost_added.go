package telebot

import "fmt"

// ChatBoostAdded represents a service message about a user boosting a chat.
type ChatBoostAdded struct {
	BoostCount int `json:"boost_count"`
}

func (c *ChatBoostAdded) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *ChatBoostAdded) Type() string        { return "ChatBoostAdded" }
