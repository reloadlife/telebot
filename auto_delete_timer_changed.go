package telebot

import "fmt"

// AutoDeleteTimerChanged represents a service message about a change in auto-delete timer settings.
type AutoDeleteTimerChanged struct {
	// AutoDeleteTime is the new auto-delete time for messages in the chat; in seconds.
	AutoDeleteTime int `json:"message_auto_delete_time"`
}

func (c *AutoDeleteTimerChanged) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *AutoDeleteTimerChanged) Type() string        { return "AutoDeleteTimerChanged" }
