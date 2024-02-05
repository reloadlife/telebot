package telebot

func (b *bot) SendChatAction(recipient Recipient, action ChatAction) error {
	params := sendChatActionRequest{
		ChatID: recipient.Recipient(),
		Action: action,
	}

	_, err := b.sendMethodRequest(methodSendChatAction, params)
	return err
}
