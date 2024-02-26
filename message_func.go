package telebot

import (
	"encoding/json"
	"fmt"
)

func (b *bot) SendMessage(recipient Recipient, text string, option ...any) (*AccessibleMessage, error) {
	params := sendMessageParams{
		ChatID: recipient,
		Text:   text,
	}

	for _, opt := range option {
		switch v := opt.(type) {
		case *MessageThreadID:
			params.MessageThreadID = v

		case ParseMode:
			params.ParseMode = &v

		case []Entity:
			params.Entities = v

		case *LinkPreviewOptions:
			params.LinkPreviewOptions = v

		case *ReplyParameters:
			params.ReplyParameters = v

		case ReplyMarkup:
			params.ReplyMarkup = v

		case Option:
			switch v {
			case Silent:
				params.DisableNotification = toPtr(true)
			case Protected:
				params.ProtectContent = toPtr(true)
			}

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in SendPhoto.")
		}
	}

	var resp struct {
		Result *AccessibleMessage
	}

	if b.offlineMode {

		return resp.Result, nil
	}

	req, err := b.sendMethodRequest(methodSendMessage, params)

	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, err
}
