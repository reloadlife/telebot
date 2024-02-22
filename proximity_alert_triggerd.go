package telebot

import "fmt"

// ProximityAlertTriggered represents the content of a service message, sent whenever a user in the chat triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {
	// Traveler is the user that triggered the alert.
	Traveler User `json:"traveler"`

	// Watcher is the user that set the alert.
	Watcher User `json:"watcher"`

	// Distance is the distance between the users.
	Distance int `json:"distance"`
}

func (c *ProximityAlertTriggered) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *ProximityAlertTriggered) Type() string        { return "ProximityAlertTriggered" }
