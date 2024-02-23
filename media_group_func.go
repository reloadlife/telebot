package telebot

import (
	"encoding/json"
	"fmt"
)

func (b *bot) SendMediaGroup(to Recipient, media []InputMedia, options ...any) ([]*AccessibleMessage, error) {
	if len(media) < 2 || len(media) > 10 {
		panic("telebot: to send media as a group, you have to provide at least 2 and upto 10 medias")
	}

	var mediaType InputMediaType
	for _, m := range media {
		if mediaType != "" {
			if mediaType != m.MediaType {
				panic("telebot: only one MediaType is able to be provided in SendMediaGroup")
			}
		}
		mediaType = m.MediaType
	}

	params := sendMediaGroupRequest{
		ChatID: to.Recipient(),
		Media:  media,
	}

	for _, option := range options {
		switch v := option.(type) {
		case *MessageThreadID:
			params.ThreadID = v

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
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in SendMediaGroup.")
		}
	}

	var resp struct {
		Result []*AccessibleMessage
	}

	req, err := b.sendMethodRequest(methodSendMediaGroup, params)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
