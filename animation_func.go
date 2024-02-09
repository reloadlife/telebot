package telebot

import (
	"encoding/json"
	"time"
)

func (b *bot) SendAnimation(to Recipient, animation File, options ...any) (*Message, error) {
	params := sendAnimationRequest{
		ChatID:    to.Recipient(),
		Animation: animation,
	}

	for _, option := range options {
		switch v := option.(type) {
		case *MessageThreadID:
			params.ThreadID = v
		case *string:
			params.Caption = v
		case string:
			params.Caption = &v
		case *ParseMode:
			params.ParseMode = v
		case []Entity:
			params.Entities = v
		case *ReplyMarkup:
			params.ReplyMarkup = v

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
			panic("telebot: unknown option type")
		}
	}

	var resp struct {
		Result *Message
	}

	req, err := b.sendMethodRequest(methodSendPhoto, params)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
