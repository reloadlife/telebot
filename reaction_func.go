package telebot

import "errors"

func (b *bot) SetMessageReaction(chatId Recipient, messageID int, options ...any) error {
	params := setMessageReactionRequest{
		ChatID:    chatId.Recipient(),
		MessageID: messageID,
		Reaction:  []ReactionType{},
	}

	for _, option := range options {
		switch value := option.(type) {
		case ReactionType:
			params.Reaction = append(params.Reaction, value)
			break
		case []ReactionType:
			params.Reaction = append(params.Reaction, value...)
			break
		case bool:
			params.IsBig = value
		default:
			return errors.New("telebot: invalid option type for SetMessageReaction method")
		}
	}

	_, err := b.sendMethodRequest(methodSetMessageReaction, params)
	return err
}
