package telebot

import (
	"encoding/json"
	"fmt"
)

type MaxTipAmount int
type SuggestedTopAmounts []int
type StartParameter string
type ProviderData string
type PhotoURL string
type PictureSize int

func (b *bot) SendInvoice(to Recipient, title, description, payload, providerToken, currency string, prices []LabeledPrice,
	options ...any) (*AccessibleMessage, error) {

	params := sendInvoiceRequest{
		ChatID:        to.Recipient(),
		Title:         title,
		Description:   description,
		Payload:       payload,
		ProviderToken: providerToken,
		Currency:      currency,
		Prices:        prices,
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case *MessageThreadID:
			params.MessageThreadID = v

		case *ReplyParameters:
			params.ReplyParameters = v

		case ReplyMarkup:
			params.ReplyMarkup = v

		case MaxTipAmount:
			params.MaxTipAmount = toPtr(int(v))

		case SuggestedTopAmounts:
			var suggestedTopAmounts []int
			for _, amount := range v {
				suggestedTopAmounts = append(suggestedTopAmounts, int(amount))
			}
			params.SuggestedTipAmounts = suggestedTopAmounts

		case StartParameter:
			params.StartParameter = toPtr(string(v))

		case ProviderData:
			params.ProviderData = toPtr(string(v))

		case PhotoURL:
			params.PhotoURL = toPtr(string(v))

		case PictureSize:
			params.PhotoSize = toPtr(int(v))

		case Width:
			params.PhotoWidth = toPtr(int(v))

		case Height:
			params.PhotoHeight = toPtr(int(v))

		case Option:
			switch v {
			case Silent:
				params.DisableNotification = toPtr(true)
			case Protected:
				params.ProtectContent = toPtr(true)
			case InvoiceNeedName:
				params.NeedName = toPtr(true)
			case InvoiceNeedPhoneNumber:
				params.NeedPhoneNumber = toPtr(true)
			case InvoiceNeedEmail:
				params.NeedEmail = toPtr(true)
			case InvoiceNeedShippingAddress:
				params.NeedShippingAddress = toPtr(true)
			case InvoiceSendPhoneNumberToProvider:
				params.SendPhoneNumberToProvider = toPtr(true)
			case InvoiceSendEmailToProvider:
				params.SendEmailToProvider = toPtr(true)
			case InvoiceIsFlexible:
				params.IsFlexible = toPtr(true)

			}

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in SendInvoice.")
		}
	}

	res, err := b.sendMethodRequest(methodSendInvoice, params)
	if err != nil {
		return nil, err
	}

	var result struct {
		Result *AccessibleMessage `json:"result"`
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return result.Result, nil
}

func (b *bot) CreateInvoiceLink(title, description, payload, providerToken, currency string, prices []LabeledPrice,
	options ...any) (*string, error) {

	params := sendInvoiceRequest{
		Title:         title,
		Description:   description,
		Payload:       payload,
		ProviderToken: providerToken,
		Currency:      currency,
		Prices:        prices,
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case MaxTipAmount:
			params.MaxTipAmount = toPtr(int(v))

		case SuggestedTopAmounts:
			var suggestedTopAmounts []int
			for _, amount := range v {
				suggestedTopAmounts = append(suggestedTopAmounts, int(amount))
			}
			params.SuggestedTipAmounts = suggestedTopAmounts

		case StartParameter:
			params.StartParameter = toPtr(string(v))

		case ProviderData:
			params.ProviderData = toPtr(string(v))

		case PhotoURL:
			params.PhotoURL = toPtr(string(v))

		case PictureSize:
			params.PhotoSize = toPtr(int(v))

		case Width:
			params.PhotoWidth = toPtr(int(v))

		case Height:
			params.PhotoHeight = toPtr(int(v))

		case Option:
			switch v {
			case InvoiceNeedName:
				params.NeedName = toPtr(true)
			case InvoiceNeedPhoneNumber:
				params.NeedPhoneNumber = toPtr(true)
			case InvoiceNeedEmail:
				params.NeedEmail = toPtr(true)
			case InvoiceNeedShippingAddress:
				params.NeedShippingAddress = toPtr(true)
			case InvoiceSendPhoneNumberToProvider:
				params.SendPhoneNumberToProvider = toPtr(true)
			case InvoiceSendEmailToProvider:
				params.SendEmailToProvider = toPtr(true)
			case InvoiceIsFlexible:
				params.IsFlexible = toPtr(true)

			}

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in CreateInvoiceLink.")
		}
	}

	res, err := b.sendMethodRequest(methodCreateInvoiceLink, params)
	if err != nil {
		return nil, err
	}

	var result struct {
		Result *string `json:"result"`
	}

	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}

	return result.Result, nil
}
