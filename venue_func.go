package telebot

import (
	"encoding/json"
	"fmt"
)

func (b *bot) SendVenue(to Recipient, venue Venue, options ...any) (*AccessibleMessage, error) {
	params := sendVenueRequest{
		ChatID:    to.Recipient(),
		Latitude:  venue.Location.Latitude,
		Longitude: venue.Location.Longitude,

		Title:           venue.Title,
		Address:         venue.Address,
		Foursquare:      &venue.FoursquareID,
		FoursquareType:  &venue.FoursquareType,
		GooglePlaceID:   &venue.GooglePlaceID,
		GooglePlaceType: &venue.GooglePlaceType,
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
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in SendVenue.")
		}
	}

	var resp struct {
		Result *AccessibleMessage
	}

	req, err := b.sendMethodRequest(methodSendVenue, params)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
