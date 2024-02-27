package telebot

import (
	"encoding/json"
	"fmt"
	"go.mamad.dev/telebot/validation"
)

func (b *bot) SendVoice(to Recipient, voice File, options ...any) (*AccessibleMessage, error) {
	params := sendVoiceRequest{
		ChatID: to.Recipient(),
		Voice:  &voice,
	}

	for _, option := range options {
		switch v := option.(type) {
		case string:
			err := validation.ValidateCaptionLength(v)
			if err != nil {
				panic("telebot: Bad Caption: " + err.Error())
			}
			params.Caption = &v

		default:
			if !b.format(&params, options...) {
				panic(fmt.Errorf(GeneralBadInputError, v, methodSendVoice))
			}
		}
	}

	var resp struct {
		Result *AccessibleMessage
	}

	req, err := b.sendMethodRequest(methodSendVoice, params)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
