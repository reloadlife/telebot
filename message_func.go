package telebot

import (
	"encoding/json"
	"fmt"
)

func (b *bot) SendMessage(recipient Recipient, text string, options ...any) (*AccessibleMessage, error) {
	params := sendMessageParams{
		ChatID: recipient,
		Text:   text,
	}

	for _, opt := range options {
		switch v := opt.(type) {

		default:
			if !b.format(&params, options...) {
				panic(fmt.Errorf(GeneralBadInputError, v, methodSendMessage))
			}
		}
	}

	var resp struct {
		Result *AccessibleMessage
	}

	req, err := b.sendMethodRequest(methodSendMessage, params)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, err
}
