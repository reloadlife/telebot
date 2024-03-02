package telebot

import (
	"encoding/json"
	"fmt"
)

func (b *bot) SendPhoto(to Recipient, photo File, options ...any) (*AccessibleMessage, error) {
	params := sendPhotoRequest{
		ChatID: to.Recipient(),
		Photo:  &photo,
	}

	for _, option := range options {
		switch v := option.(type) {
		case string:
			params.Caption = &v

		default:
			if !b.format(&params, options...) {
				panic(fmt.Errorf(GeneralBadInputError, v, methodSendPhoto))
			}
		}
	}

	var resp struct {
		Result *AccessibleMessage
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
