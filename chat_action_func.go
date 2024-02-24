package telebot

import "fmt"

func (b *bot) SendChatAction(recipient Recipient, action ChatAction, options ...any) error {
	params := sendChatActionRequest{
		ChatID: recipient.Recipient(),
		Action: action,
	}

	for _, option := range options {
		switch v := option.(type) {
		case *MessageThreadID:
			params.ThreadID = v
		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in SendChatAction.")
		}
	}

	_, err := b.sendMethodRequest(methodSendChatAction, params)
	return err
}
