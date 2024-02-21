package telebot

import "fmt"

// CallbackGame represents a placeholder for game-related information.
// Use BotFather to set up your game.
type CallbackGame struct{}

func (c *CallbackGame) Type() string        { return "callback_game" }
func (c *CallbackGame) ReflectType() string { return fmt.Sprintf("%T", c) }
