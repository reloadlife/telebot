package telebot

import "fmt"

// InlineQuery represents an incoming inline query. When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	// ID is the unique identifier for this query.
	ID string `json:"id"`

	// From is the sender of the inline query.
	From User `json:"from"`

	// Query is the text of the query (up to 256 characters).
	Query string `json:"query"`

	// Offset is the offset of the results to be returned, which can be controlled by the bot.
	Offset string `json:"offset"`

	// ChatType is the type of the chat from which the inline query was sent.
	// It can be either “sender” for a private chat with the inline query sender, “private”, “group”, “supergroup”, or “channel”.
	// The chat type should be always known for requests sent from official clients and most third-party clients unless the request was sent from a secret chat.
	ChatType ChatType `json:"chat_type,omitempty"`

	// Location is the sender's location (optional, only for bots that request user location).
	Location *Location `json:"location,omitempty"`
}

func (c *InlineQuery) ReflectType() string {
	return fmt.Sprintf("%T", c)
}

func (c *InlineQuery) Type() string {
	return "InlineQuery"
}

type ChosenInlineResult struct {
	// ResultID is the unique identifier for the result that was chosen.
	ID string `json:"result_id"`

	// From is the user that chose the result.
	From User `json:"from"`

	// Location is the sender's location (optional, only for bots that require user location).
	Location *Location `json:"location,omitempty"`

	// InlineMessageID is the identifier of the sent inline message. Available only if there is an inline keyboard attached to the message.
	// Will be also received in callback queries and can be used to edit the message.
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// Query is the query that was used to obtain the result.
	Query string `json:"query"`
}

func (c *ChosenInlineResult) ReflectType() string {
	return fmt.Sprintf("%T", c)
}

func (c *ChosenInlineResult) Type() string {
	return "ChosenInlineResult"
}
