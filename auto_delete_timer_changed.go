package telebot

import "fmt"

// AutoDeleteTimerChanged represents a service message about a change in auto-delete timer settings.
// <a href="https://core.telegram.org/bots/api#messageautodeletetimerchanged">/bots/api#messageautodeletetimerchanged</a>
type AutoDeleteTimerChanged struct {
	// AutoDeleteTime is the new auto-delete time for messages in the chat; in seconds.
	AutoDeleteTime int `json:"message_auto_delete_time"`
}

// ReflectType returns the reflect type of the struct
func (a *AutoDeleteTimerChanged) ReflectType() string { return fmt.Sprintf("%T", a) }

// Type returns the type of the message
func (a *AutoDeleteTimerChanged) Type() string { return "AutoDeleteTimerChanged" }
