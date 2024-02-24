package telebot

import (
	"encoding/json"
	"fmt"
	httpc "go.mamad.dev/telebot/http"
	"io"
	"strconv"
)

func (b *bot) SendMediaGroup(to Recipient, media []InputMedia, options ...any) ([]*AccessibleMessage, error) {
	if len(media) < 2 || len(media) > 10 {
		panic("telebot: to send media as a group, you have to provide at least 2 and upto 10 medias")
	}

	files := make([]httpc.File, 0)

	var mediaType InputMediaType
	for i, m := range media {
		if mediaType != "" {
			if mediaType != m.MediaType {
				panic("telebot: only one MediaType is able to be provided in SendMediaGroup")
			}
		}
		mediaType = m.MediaType
		switch {
		case m.Media.GetFileReader() != nil:
			m.Repr = "attach://" + strconv.Itoa(i)
			r, _ := io.ReadAll(m.Media.GetFileReader())
			files = append(files, httpc.File{
				Name:     strconv.Itoa(i),
				FileName: m.Media.GetFileName(),
				DATA:     r,
			})
		case m.Media != nil:

		default:
			return nil, fmt.Errorf("telebot: album entry #%d does not exist", i)
		}
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

	req, err := b.sendMethodRequest(methodSendMediaGroup, params, files...)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
