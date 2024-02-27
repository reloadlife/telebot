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

		default:
			if !b.format(&params, options...) {
				panic(fmt.Errorf(GeneralBadInputError, v, methodSendContact))
			}
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
