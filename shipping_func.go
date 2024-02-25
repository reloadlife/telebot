package telebot

import (
	"fmt"
)

func (b *bot) AnswerShippingQuery(ID string, ok bool, options ...any) error {
	params := answerShippingQueryRequest{
		ID:              ID,
		OK:              ok,
		ShippingOptions: []ShippingOption{},
	}

	for _, option := range options {
		switch v := option.(type) {
		case ShippingOption:
			params.ShippingOptions = append(params.ShippingOptions, v)
		case []ShippingOption:
			params.ShippingOptions = append(params.ShippingOptions, v...)

		case string:
			params.ErrorMessage = v

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in AnswerShippingQuery.")
		}
	}

	_, err := b.sendMethodRequest(methodAnswerShippingQuery, params)
	return err
}

func (b *bot) AnswerPreCheckoutQuery(queryID string, ok bool, errorMessage *string) error {
	params := answerShippingQueryRequest{
		ID: queryID,
		OK: ok,
	}

	if errorMessage != nil {
		params.ErrorMessage = *errorMessage
	}

	_, err := b.sendMethodRequest(methodAnswerPreCheckoutQuery, params)
	return err
}
