package telebot

import "fmt"

type ReplyParameters struct {
	MessageID                int      `json:"message_id"`
	ChatID                   any      `json:"chat_id,omitempty"`
	AllowSendingWithoutReply *bool    `json:"allow_sending_without_reply,omitempty"`
	Quote                    *string  `json:"quote,omitempty"`
	QuoteParseMode           *string  `json:"quote_parse_mode,omitempty"`
	QuoteEntities            []Entity `json:"quote_entities,omitempty"`
	QuotePosition            *int     `json:"quote_position,omitempty"`
}

func (c *ReplyParameters) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *ReplyParameters) Type() string        { return "ReplyParameters" }
