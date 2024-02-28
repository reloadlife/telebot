package telebot

import (
	"fmt"
	"net/url"
)

// NewInlineKeyboardButton creates a new inline keyboard button with the given text and options.
func NewInlineKeyboardButton(text string, options ...any) Button {
	btn := &InlineKeyboardButton{
		Text: text,
	}

	for _, option := range options {
		switch value := option.(type) {
		case string:
			btn.CallbackData = toPtr(fmt.Sprintf("\f%s", value))

		case *string:
			btn.CallbackData = toPtr(fmt.Sprintf("\f%s", *value))

		case *url.URL:
			btn.URL = toPtr(value.String())

		case *WebAppInfo:
			btn.WebApp = value

		case *LoginURL:
			btn.LoginURL = value

		case *SwitchInlineQueryStringType:
			btn.SwitchInlineQuery = toPtr(string(*value))

		case *SwitchInlineQueryChosenChatStringType:
			btn.SwitchInlineQueryCurrentChat = toPtr(string(*value))

		case *SwitchInlineQueryChosenChat:
			btn.SwitchInlineQueryChosenChat = value

		case *CallbackGame:
			btn.CallbackGame = value

		case bool:
			btn.Pay = &value
		}
	}

	if len(options) == 0 {
		panic("telebot: NewInlineKeyboardButton: no options provided")
	}

	return btn
}

// NewKeyboardButton creates a new keyboard button with the given text and options.
func NewKeyboardButton(text string, options ...any) Button {
	btn := &KeyboardButton{
		Text: text,
	}

	for _, option := range options {
		switch value := option.(type) {
		case *KeyboardButtonRequestUsers:
			btn.RequestUsers = value

		case *KeyboardButtonRequestChat:
			btn.RequestChat = value

		case *RequestContact:
			btn.RequestContact = toPtr(bool(*value))

		case *RequestLocation:
			btn.RequestLocation = toPtr(bool(*value))

		case *KeyboardButtonPollType:
			btn.RequestPoll = value

		case *WebAppInfo:
			btn.WebApp = value
		}
	}

	return btn
}
