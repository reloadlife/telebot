package telebot

import (
	"encoding/json"
	"fmt"
	"time"
)

func (b *bot) SendVideo(to Recipient, video File, options ...any) (*AccessibleMessage, error) {
	params := sendVideoParams{
		ChatID: to.Recipient(),
		Video:  &video,
	}

	for _, option := range options {
		switch v := option.(type) {
		case string:
			params.Caption = &v

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
			case Streamable:
				params.Streamable = toPtr(true)
			}

		default:
			if !b.format(&params, options...) {
				panic(fmt.Errorf(GeneralBadInputError, v, methodSendVideo))
			}
		}
	}

	var resp struct {
		Result *AccessibleMessage
	}

	req, err := b.sendMethodRequest(methodSendVideo, params)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
