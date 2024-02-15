package telebot

import "encoding/json"

type ReplyMarkup interface {
	ReplyMarkup()

	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	// InlineKeyboard is an array of button rows, each represented by an array of InlineKeyboardButton objects.
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

func (m *InlineKeyboardMarkup) ReplyMarkup() {}
func (m *InlineKeyboardMarkup) MarshalJSON() ([]byte, error) {
	return json.Marshal(m)
}
func (m *InlineKeyboardMarkup) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, m)
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options.
type ReplyKeyboardMarkup struct {
	// Keyboard is an array of button rows, each represented by an array of KeyboardButton objects.
	Keyboard [][]KeyboardButton `json:"keyboard"`

	// IsPersistent requests clients to always show the keyboard when the regular keyboard is hidden. Defaults to false, in which case the custom keyboard can be hidden and opened with a keyboard icon.
	IsPersistent bool `json:"is_persistent,omitempty"`

	// ResizeKeyboard requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`

	// OneTimeKeyboard requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat - the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`

	// InputFieldPlaceholder is the placeholder to be shown in the input field when the keyboard is active; 1-64 characters.
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`

	// Selective is used if you want to show the keyboard to specific users only.
	// Targets:
	// 1) Users that are @mentioned in the text of the Message object.
	// 2) If the bot's message is a reply to a message in the same chat and forum topic, the sender of the original message.
	Selective bool `json:"selective,omitempty"`
}
