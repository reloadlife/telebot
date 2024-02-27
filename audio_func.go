package telebot

import (
	"fmt"
)

func (b *bot) SendAudio(to Recipient, audio File, options ...any) (*AccessibleMessage, error) {
	params := sendAudioRequest{
		ChatID: to,
		Audio:  &audio,
	}

	for _, option := range options {
		switch v := option.(type) {
		case string:
			params.Caption = &v

		case *File:
			params.Thumbnail = v

		default:
			if !b.format(&params, options...) {
				panic(fmt.Errorf(GeneralBadInputError, v, methodSendAnimation))
			}
		}
	}

	req, err := b.sendMethodRequest(methodSendAudio, params)
	return b.resAsMessage(req, err)
}
