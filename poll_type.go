package telebot

import "fmt"

// PollOption represents an option in a poll.
type PollOption struct {
	// Text is the text of the option.
	Text string `json:"text"`

	// VoterCount is the number of users who voted for this option.
	VoterCount int `json:"voter_count"`
}

func (c *PollOption) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *PollOption) Type() string        { return "PollOption" }
