package telebot

import (
	"encoding/json"
	"fmt"
)

func (b *bot) SendContact(to Recipient, contact Contact, options ...any) (*AccessibleMessage, error) {
	if contact.UserID != nil {
		panic("telebot: contact.UserID must be nil in SendContact.")
	}
	params := sendContactRequest{
		ChatID:      to.Recipient(),
		PhoneNumber: contact.PhoneNumber,
		FirstName:   contact.FirstName,
		LastName:    contact.LastName,
		VCard:       contact.VCard,
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
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in SendContact.")
		}
	}

	var resp struct {
		Result *AccessibleMessage
	}

	req, err := b.sendMethodRequest(methodSendContact, params)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
