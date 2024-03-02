package telebot

import "fmt"

type sendPollRequest struct {
	ChatID any `json:"chat_id"`

	Question             string    `json:"question"`
	Options              []string  `json:"options"`
	IsAnonymous          *bool     `json:"is_anonymous"`
	Type                 *PollType `json:"type"`
	AllowMultipleAnswers *bool     `json:"allows_multiple_answers"`
	CorrectOptionID      *int      `json:"correct_option_id,omitempty"`
	Explanation          *string   `json:"explanation,omitempty"`
	ExplanationEntities  []Entity  `json:"explanation_entities,omitempty"`
	ExplanationParseMode ParseMode `json:"explanation_parse_mode,omitempty"`
	OpenPeriod           *int      `json:"open_period,omitempty"`
	CloseDate            *int      `json:"close_date,omitempty"`
	IsClosed             *bool     `json:"is_closed,omitempty"`

	ThreadID *MessageThreadID `json:"message_thread_id,omitempty"`

	DisableNotification *bool            `json:"disable_notification,omitempty"`
	Protected           *bool            `json:"protect_content,omitempty"`
	ReplyParameters     *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup         ReplyMarkup      `json:"reply_markup,omitempty"`
}

// PollOption represents an option in a poll.
type PollOption struct {
	// Text is the text of the option.
	Text string `json:"text"`

	// VoterCount is the number of users who voted for this option.
	VoterCount int `json:"voter_count"`
}

func (p *PollOption) ReflectType() string { return fmt.Sprintf("%T", p) }
func (p *PollOption) Type() string        { return "PollOption" }
