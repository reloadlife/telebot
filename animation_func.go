package telebot

import (
	"fmt"
	"go.mamad.dev/telebot/validation"
	"time"
)

func (b *bot) SendAnimation(to Recipient, animation File, options ...any) (*AccessibleMessage, error) {
	params := sendAnimationRequest{
		ChatID:    to,
		Animation: &animation,
	}

	for _, option := range options {
		switch v := option.(type) {
		case *MessageThreadID:
			params.ThreadID = v

		case string:
			err := validation.ValidateCaptionLength(v)
			if err != nil {
				panic("telebot: Bad Caption: " + err.Error())
			}
			params.Caption = &v

		case ParseMode:
			params.ParseMode = v

		case []Entity:
			params.Entities = append(params.Entities, v...)
		case Entity:
			params.Entities = append(params.Entities, v)

		case ReplyMarkup:
			params.ReplyMarkup = v
		case *ReplyParameters:
			params.ReplyParameters = v

		case time.Duration:
			params.Duration = toPtr(int(v.Seconds()))

		case Width:
			params.Width = toPtr(int(v))
		case Height:
			params.Height = toPtr(int(v))

		case *File:
			params.Thumbnail = v

		case Option:
			switch v {
			case Silent:
				params.DisableNotification = toPtr(true)
			case Protected:
				params.Protected = toPtr(true)
			case HasSpoiler:
				params.HasSpoiler = toPtr(true)
			}

		default:
			panic(fmt.Errorf(GeneralBadInputError, v, methodSendAnimation))
		}
	}

	req, err := b.sendMethodRequest(methodSendAnimation, params)
	return b.resAsMessage(req, err)
}
