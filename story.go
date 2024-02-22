package telebot

import "fmt"

// Story object represents a message about a forwarded story in the chat
type Story struct {
	ID   int64 `json:"id"`
	Chat *Chat `json:"chat"`
}

func (c *Story) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *Story) Type() string        { return "Story" }
