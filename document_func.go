package telebot

import (
	"encoding/json"
	"fmt"
	"go.mamad.dev/telebot/validation"
)

func (b *bot) SendDocument(to Recipient, doc File, options ...any) (*AccessibleMessage, error) {
	params := sendDocumentParams{
		ChatID:   to.Recipient(),
		Document: &doc,
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
			if !b.format(&params, options...) {
				panic(fmt.Errorf(GeneralBadInputError, v, methodSendDocument))
			}
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
