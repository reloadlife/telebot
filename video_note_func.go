package telebot

import (
	"encoding/json"
	"fmt"
	"time"
)

func (b *bot) SendVideoNote(to Recipient, videoNote File, options ...any) (*AccessibleMessage, error) {
	params := sendVideoNoteRequest{
		ChatID:    to.Recipient(),
		VideoNote: &videoNote,
	}

	for _, option := range options {
		switch v := option.(type) {
		case *MessageThreadID:
			params.ThreadID = v
		case ReplyMarkup:
			params.ReplyMarkup = v
		case *ReplyParameters:
			params.ReplyParameters = v

		case time.Duration:
			params.Duration = toPtr(int(v.Seconds()))
		case Height:
			params.Length = toPtr(int(v))
		case Width:
			params.Length = toPtr(int(v))

		case *File:
			params.Thumbnail = v

		case Option:
			switch v {
			case Silent:
				params.DisableNotification = toPtr(true)
			case Protected:
				params.Protected = toPtr(true)
			}

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in SendVideoNote.")
		}
	}

	var resp struct {
		Result *AccessibleMessage
	}

	req, err := b.sendMethodRequest(methodSendVideoNote, params)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
