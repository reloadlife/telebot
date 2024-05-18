package telebot

import "fmt"

// Poll contains information about a poll.
type Poll struct {
	// ID is the unique poll identifier.
	ID string `json:"id"`

	// Question is the poll question (1-300 characters).
	Question string `json:"question"`

	// Options is the list of poll options.
	Options []PollOption `json:"options"`

	// TotalVoterCount is the total number of users that voted in the poll.
	TotalVoterCount int `json:"total_voter_count"`

	// IsClosed is true if the poll is closed.
	IsClosed bool `json:"is_closed"`

	// IsAnonymous is true if the poll is anonymous.
	IsAnonymous bool `json:"is_anonymous"`

	// Type is the poll type, currently can be “regular” or “quiz”.
	PollType PollType `json:"type"`

	// AllowsMultipleAnswers is true if the poll allows multiple answers.
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`

	// CorrectOptionID is the 0-based identifier of the correct answer option (optional).
	// Available only for polls in quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	CorrectOptionID int `json:"correct_option_id,omitempty"`

	// Explanation is the text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll (optional, 0-200 characters).
	Explanation string `json:"explanation,omitempty"`

	// ExplanationEntities is the list of special entities like usernames, URLs, bot commands, etc. that appear in the explanation (optional).
	ExplanationEntities []Entity `json:"explanation_entities,omitempty"`

	// OpenPeriod is the amount of time in seconds the poll will be active after creation (optional).
	OpenPeriod int `json:"open_period,omitempty"`

	// CloseDate is the point in time (Unix timestamp) when the poll will be automatically closed (optional).
	CloseDate int `json:"close_date,omitempty"`

	QuestionEntities []Entity `json:"question_entities,omitempty"`
}

func (p *Poll) ReflectType() string {
	return fmt.Sprintf("%T", p)
}

func (p *Poll) Type() string {
	if p.PollType == "" {
		return Unknown
	}
	return string(p.PollType)
}
