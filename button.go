package telebot

import (
	"encoding/json"
)

type Row []Button

// UnmarshalJSON implements json.Unmarshaler.
// It supports both InlineKeyboardButton and KeyboardButton.
// It is a bit complicated to unmarshal into interface types, believe me.
// this thing works perfectly for our goal, if you have a better idea, a PR is appreciated.
func (r *Row) UnmarshalJSON(data []byte) error {
	var buttons []json.RawMessage
	if err := json.Unmarshal(data, &buttons); err != nil {
		return err
	}

	butts := make([]Button, 0, len(buttons))

	for _, buttonData := range buttons {
		var rawButton struct {
			URL                          *string                      `json:"url,omitempty"`
			CallbackData                 *string                      `json:"callback_data,omitempty"`
			LoginURL                     *LoginUrl                    `json:"login_url,omitempty"`
			SwitchInlineQuery            *string                      `json:"switch_inline_query,omitempty"`
			SwitchInlineQueryCurrentChat *string                      `json:"switch_inline_query_current_chat,omitempty"`
			SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
			CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
			Pay                          *bool                        `json:"pay,omitempty"`

			RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
			RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`
			RequestContact  *bool                       `json:"request_contact,omitempty"`
			RequestLocation *bool                       `json:"request_location,omitempty"`
			RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`

			Text   string      `json:"text"`
			WebApp *WebAppInfo `json:"web_app,omitempty"`
		}

		if err := json.Unmarshal(buttonData, &rawButton); err != nil {
			return err
		}

		switch {
		case rawButton.URL != nil || rawButton.LoginURL != nil || rawButton.CallbackData != nil ||
			rawButton.SwitchInlineQuery != nil || rawButton.SwitchInlineQueryCurrentChat != nil ||
			rawButton.SwitchInlineQueryChosenChat != nil || rawButton.CallbackGame != nil || rawButton.Pay != nil:
			butts = append(butts, &InlineKeyboardButton{
				Text:                         rawButton.Text,
				URL:                          rawButton.URL,
				CallbackData:                 rawButton.CallbackData,
				WebApp:                       rawButton.WebApp,
				LoginURL:                     rawButton.LoginURL,
				SwitchInlineQuery:            rawButton.SwitchInlineQuery,
				SwitchInlineQueryCurrentChat: rawButton.SwitchInlineQueryCurrentChat,
				SwitchInlineQueryChosenChat:  rawButton.SwitchInlineQueryChosenChat,
				CallbackGame:                 rawButton.CallbackGame,
				Pay:                          rawButton.Pay,
			})

		case rawButton.RequestUsers != nil || rawButton.RequestChat != nil ||
			rawButton.RequestContact != nil || rawButton.RequestLocation != nil || rawButton.RequestPoll != nil:
			butts = append(butts, &KeyboardButton{
				Text:            rawButton.Text,
				RequestUsers:    rawButton.RequestUsers,
				RequestChat:     rawButton.RequestChat,
				RequestContact:  rawButton.RequestContact,
				RequestLocation: rawButton.RequestLocation,
				RequestPoll:     rawButton.RequestPoll,
			})

		default:
			return ErrUnsupportedButton
		}
	}

	*r = butts
	return nil
}

type Button interface {
	Button()

	String() string

	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}
