package telebot

import (
	"fmt"
	"go.mamad.dev/telebot/validation"
)

func (b *bot) SendAnimation(to Recipient, animation File, options ...any) (*AccessibleMessage, error) {
	params := sendAnimationRequest{
		ChatID:    to,
		Animation: &animation,
	}

	for _, option := range options {
		switch v := option.(type) {
		case string:
			err := validation.ValidateCaptionLength(v)
			if err != nil {
				panic("telebot: Bad Caption: " + err.Error())
			}
			params.Caption = &v

		case *File:
			params.Thumbnail = v

		default:
			if !format(&params, options...) {
				panic(fmt.Errorf(GeneralBadInputError, v, methodSendAnimation))
			}
		}
	}

	req, err := b.sendMethodRequest(methodSendAnimation, params)
	return b.resAsMessage(req, err)
}
