package telebot

import (
	"encoding/json"
	"fmt"
)

func (b *bot) SendLocation(to Recipient, location Location, options ...any) (*AccessibleMessage, error) {
	params := sendLocationRequest{
		ChatID:    to.Recipient(),
		Longitude: location.Longitude,
		Latitude:  location.Latitude,

		// The following fields are optional.
		HorizontalAccuracy:   &location.HorizontalAccuracy,
		LivePeriod:           &location.LivePeriod,
		Heading:              &location.Heading,
		ProximityAlertRadius: &location.ProximityAlertRadius,
	}

	for _, option := range options {
		switch v := option.(type) {
		case *MessageThreadID:
			params.ThreadID = v
		case ReplyMarkup:
			params.ReplyMarkup = v
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
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in SendLocation.")
		}
	}

	var resp struct {
		Result *AccessibleMessage
	}

	req, err := b.sendMethodRequest(methodSendLocation, params)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
