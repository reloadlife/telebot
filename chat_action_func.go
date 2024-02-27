package telebot

import "fmt"

func (b *bot) SendChatAction(recipient Recipient, action ChatAction, options ...any) error {
	params := sendChatActionRequest{
		ChatID: recipient.Recipient(),
		Action: action,
	}

	for _, option := range options {
		switch v := option.(type) {
		default:
			if !b.format(&params, options...) {
				panic(fmt.Errorf(GeneralBadInputError, v, methodSendChatAction))
			}
		}
	}

	_, err := b.sendMethodRequest(methodSendChatAction, params)
	return err
}
