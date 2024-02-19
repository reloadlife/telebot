package telebot

import (
	"encoding/json"
)

func (b *bot) Forward(to Recipient, From Recipient, messageID int, opts ...any) (*AccessibleMessage, error) {
	params := forwardMessageRequest{
		ChatID:    to.Recipient(),
		FromChat:  From.Recipient(),
		MessageID: messageID,
	}

	for _, opt := range opts {
		switch v := opt.(type) {
		case Option:
			switch v {
			case Silent:
				params.DisableNotification = toPtr(true)
			case Protected:
				params.Protect = toPtr(true)
			}
		case *MessageThreadID:
			params.MessageThreadID = v
		}
	}

	req, err := b.sendMethodRequest(methodForwardMessage, params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result AccessibleMessage
	}
	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return &resp.Result, nil
}

func (b *bot) Forwards(to Recipient, From Recipient, messageIDs []int, opts ...any) ([]int, error) {
	params := forwardMessageRequest{
		ChatID:     to.Recipient(),
		FromChat:   From.Recipient(),
		MessageIDs: messageIDs,
	}

	for _, opt := range opts {
		switch v := opt.(type) {
		case Option:
			switch v {
			case Silent:
				params.DisableNotification = toPtr(true)
			case Protected:
				params.Protect = toPtr(true)
			}
		case *MessageThreadID:
			params.MessageThreadID = v
		}
	}

	req, err := b.sendMethodRequest(methodForwardMessages, params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result []int
	}
	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

func (b *bot) Copy(to Recipient, from Recipient, messageIDs []int, opts ...any) ([]int, error) {
	params := copyMessageRequest{
		ChatID:     to.Recipient(),
		FromChat:   from.Recipient(),
		MessageIDs: messageIDs,
	}

	for _, opt := range opts {
		switch v := opt.(type) {
		case Option:
			switch v {
			case Silent:
				params.DisableNotification = toPtr(true)
			case Protected:
				params.Protect = toPtr(true)
			case RemoveCaption:
				params.RemoveCaption = toPtr(true)
			}
		case *MessageThreadID:
			params.MessageThreadID = v
		}
	}

	req, err := b.sendMethodRequest(methodCopyMessages, params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result []int
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
