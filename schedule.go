package telebot

import "fmt"

// VideoChatScheduled represents a service message about a video chat scheduled in the chat.
type VideoChatScheduled struct {
	// StartDate is the point in time (Unix timestamp) when the video chat is supposed to be started by a chat administrator.
	StartDate int `json:"start_date"`
}

func (c *VideoChatScheduled) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *VideoChatScheduled) Type() string        { return "VideoChatScheduled" }

// VideoChatStarted represents a service message about a video chat started in the chat.
// Currently holds no information.
type VideoChatStarted struct{}

func (c *VideoChatStarted) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *VideoChatStarted) Type() string        { return "VideoChatStarted" }

// VideoChatEnded represents a service message about a video chat ended in the chat.
type VideoChatEnded struct {
	// Duration is the video chat duration in seconds.
	Duration int `json:"duration"`
}

func (c *VideoChatEnded) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *VideoChatEnded) Type() string        { return "VideoChatEnded" }

// VideoChatParticipantsInvited represents a service message about new members invited to a video chat.
type VideoChatParticipantsInvited struct {
	// Users is the list of new members that were invited to the video chat.
	Users []User `json:"users"`
}

func (c *VideoChatParticipantsInvited) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *VideoChatParticipantsInvited) Type() string        { return "VideoChatParticipantsInvited" }
