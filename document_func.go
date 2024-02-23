package telebot

import (
	"encoding/json"
	"fmt"
)

func (b *bot) SendDocument(to Recipient, doc File, options ...any) (*AccessibleMessage, error) {
	params := sendDocumentParams{
		ChatID:   to.Recipient(),
		Document: &doc,
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
			params.ParseMode = *v
		case []Entity:
			params.Entities = v
		case ReplyMarkup:
			params.ReplyMarkup = v
		case *ReplyParameters:
			params.ReplyParameters = v

		case *File:
			params.Thumbnail = v

		case Option:
			switch v {
			case Silent:
				params.DisableNotification = toPtr(true)
			case Protected:
				params.Protected = toPtr(true)
			case DisableContentTypeDetection:
				params.DisableContentTypeDetection = toPtr(true)
			}

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in SendDocument.")
		}
	}

	var resp struct {
		Result *AccessibleMessage
	}

	req, err := b.sendMethodRequest(methodSendDocument, params)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
