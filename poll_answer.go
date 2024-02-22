package telebot

import "fmt"

// PollAnswer represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	// PollID is the unique poll identifier.
	ID string `json:"poll_id"`

	// VoterChat is the chat that changed the answer to the poll, if the voter is anonymous (optional).
	VoterChat *Chat `json:"voter_chat,omitempty"`

	// User is the user that changed the answer to the poll, if the voter isn't anonymous (optional).
	User *User `json:"user,omitempty"`

	// OptionIDs is the 0-based identifiers of chosen answer options.
	// May be empty if the vote was retracted.
	OptionIDs []int `json:"option_ids"`
}

func (c *PollAnswer) ReflectType() string {
	return fmt.Sprintf("%T", c)
}

func (c *PollAnswer) Type() string {
	return "PollAnswer"
}
