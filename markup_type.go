package telebot

import "fmt"

// ForceReplyMarkup requests clients to display a reply interface to the user (act as if the user selected the bot's message and tapped 'Reply').
// This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
// Available in private chats only
type ForceReplyMarkup struct {
	// ForceReply requests clients to display a reply interface to the user (act as if the user selected the bot's message and tapped 'Reply').
	// This can be extremely useful if you want to create user-friendly step-by-step interfaces without having to sacrifice privacy mode.
	// Available in private chats only
	ForceReply bool `json:"force_reply"`
	// InputFieldPlaceholder is the placeholder to be shown in the input field when the reply interface is active; 1-64 characters
	InputFieldPlaceholder *string `json:"input_field_placeholder,omitempty"`
	// Selective Use this parameter if you want to force reply from specific users only
	Selective *bool `json:"selective,omitempty"`
}

// ReplyKeyboardRemove removes the current custom keyboard and displays the default letter-keyboard.
type ReplyKeyboardRemove struct {
	// RemoveKeyboard Requests clients to remove the custom keyboard
	RemoveKeyboard bool `json:"remove_keyboard"`
	// Selective Use this parameter to remove the keyboard for specific users only
	Selective *bool `json:"selective,omitempty"`
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	// InlineKeyboard is an array of button rows, each represented by an array of InlineKeyboardButton objects.
	InlineKeyboard []Row `json:"inline_keyboard"`
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options.
type ReplyKeyboardMarkup struct {
	// Keyboard is an array of button rows, each represented by an array of KeyboardButton objects.
	Keyboard []Row `json:"keyboard"`

	// IsPersistent requests clients to always show the keyboard when the regular keyboard is hidden. Defaults to false, in which case the custom keyboard can be hidden and opened with a keyboard icon.
	IsPersistent *bool `json:"is_persistent,omitempty"`

	// ResizeKeyboard requests clients to resize the keyboard vertically for optimal fit (e.g., make the keyboard smaller if there are just two rows of buttons). Defaults to false, in which case the custom keyboard is always of the same height as the app's standard keyboard.
	ResizeKeyboard *bool `json:"resize_keyboard,omitempty"`

	// OneTimeKeyboard requests clients to hide the keyboard as soon as it's been used. The keyboard will still be available, but clients will automatically display the usual letter-keyboard in the chat - the user can press a special button in the input field to see the custom keyboard again. Defaults to false.
	OneTimeKeyboard *bool `json:"one_time_keyboard,omitempty"`

	// InputFieldPlaceholder is the placeholder to be shown in the input field when the keyboard is active; 1-64 characters.
	InputFieldPlaceholder *string `json:"input_field_placeholder,omitempty"`

	// Selective is used if you want to show the keyboard to specific users only.
	// Targets:
	// 1) Users that are @mentioned in the text of the AccessibleMessage object.
	// 2) If the bot's message is a reply to a message in the same chat and forum topic, the sender of the original message.
	Selective *bool `json:"selective,omitempty"`
}

// Type Functions

func (m *ForceReplyMarkup) Type() string            { return "ForceReplyMarkup" }
func (m *ForceReplyMarkup) ReflectType() string     { return fmt.Sprintf("%T", m) }
func (m *ReplyKeyboardRemove) Type() string         { return "ReplyKeyboardRemove" }
func (m *ReplyKeyboardRemove) ReflectType() string  { return fmt.Sprintf("%T", m) }
func (m *InlineKeyboardMarkup) Type() string        { return "InlineKeyboardMarkup" }
func (m *InlineKeyboardMarkup) ReflectType() string { return fmt.Sprintf("%T", m) }
func (m *ReplyKeyboardMarkup) Type() string         { return "ReplyKeyboardMarkup" }
func (m *ReplyKeyboardMarkup) ReflectType() string  { return fmt.Sprintf("%T", m) }
