package telebot

import "fmt"

type BotCommand struct {
	Command     string `json:"command"`
	Description string `json:"description"`
}

func (c *BotCommand) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *BotCommand) Type() string        { return "BotCommand" }

type BotCommandScope struct {
	ScopeType BotCommandScopeType `json:"type"`
	ChatID    any                 `json:"chat_id,omitempty"`
	UserID    any                 `json:"user_id,omitempty"`
}

func (c *BotCommandScope) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *BotCommandScope) Type() string {
	if c.ScopeType == "" {
		return Unknown
	}
	return string(c.ScopeType)
}
