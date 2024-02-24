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
		case *MessageThreadID:
			params.ThreadID = v

		case DiceEmoji:
			params.Emoji = v

		case ReplyMarkup:
			params.ReplyMarkup = v
		case *ReplyParameters:
			params.ReplyParameters = v

		case Option:
			switch v {
			case Silent:
				params.DisableNotification = toPtr(true)
			case Protected:
				params.Protected = toPtr(true)
			}

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in SendDice.")
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
