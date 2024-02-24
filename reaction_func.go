package telebot

import (
	"fmt"
)

func (b *bot) SetMessageReaction(to Recipient, messageID int64, options ...any) error {
	params := setMessageReactionRequest{
		ChatID:    to.Recipient(),
		MessageID: messageID,
		Reaction:  []ReactionType{},
	}

	for _, option := range options {
		switch v := option.(type) {
		case ReactionType:
			params.Reaction = append(params.Reaction, v)
		case []ReactionType:
			params.Reaction = append(params.Reaction, v...)
		case bool:
			params.IsBig = v
		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in SetMessageReaction.")
		}
	}

	_, err := b.sendMethodRequest(methodSetMessageReaction, params)
	return err
}
