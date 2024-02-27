package telebot

import (
	"encoding/json"
	"fmt"
)

func (b *bot) SendDice(to Recipient, options ...any) (*AccessibleMessage, error) {
	params := sendDiceRequest{
		ChatID: to.Recipient(),
	}

	for _, option := range options {
		switch v := option.(type) {
		case DiceEmoji:
			params.Emoji = v

		default:
			if !b.format(&params, options...) {
				panic(fmt.Errorf(GeneralBadInputError, v, methodSendDice))
			}
		}
	}

	var resp struct {
		Result *AccessibleMessage
	}

	req, err := b.sendMethodRequest(methodSendDice, params)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
