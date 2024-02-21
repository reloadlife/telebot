package telebot

import "encoding/json"

// markupType is the type of the markup.
// telegram as if v7.1 has 4 types of markup:
type markupType int

const (
	markupTypeInline markupType = iota
	markupTypeKeyboard
	markupTypeRemoveKeyboard
	markupTypeForceReply
)

type MenuButton struct {
}

// SwitchInlineQueryChosenChat represents an inline button that switches the current user to inline mode
// in a chosen chat, with an optional default inline query.
type SwitchInlineQueryChosenChat struct {
	// Query is the default inline query to be inserted in the input field.
	// If left empty, only the bot's username will be inserted.
	Query *string `json:"query,omitempty"`

	// AllowUserChats, if true, allows private chats with users to be chosen.
	AllowUserChats *bool `json:"allow_user_chats,omitempty"`

	// AllowBotChats, if true, allows private chats with bots to be chosen.
	AllowBotChats *bool `json:"allow_bot_chats,omitempty"`

	// AllowGroupChats, if true, allows group and supergroup chats to be chosen.
	AllowGroupChats *bool `json:"allow_group_chats,omitempty"`

	// AllowChannelChats, if true, allows channel chats to be chosen.
	AllowChannelChats *bool `json:"allow_channel_chats,omitempty"`
}

// KeyboardButtonRequestChat defines the criteria used to request a suitable chat.
type KeyboardButtonRequestChat struct {
	// RequestID Signed 32-bit identifier of the request
	RequestID int32 `json:"request_id"`
	// ChatIsChannel Pass true to request a channel chat, pass false to request a group or a supergroup chat
	ChatIsChannel bool `json:"chat_is_channel"`
	// ChatIsForum  Optional. Pass true to request a forum supergroup, pass false to request a non-forum chat
	ChatIsForum bool `json:"chat_is_forum,omitempty"`
	// ChatHasUsernameOptional. Pass true to request a supergroup or a channel with a username, pass false to request a chat without a username
	ChatHasUsername bool `json:"chat_has_username,omitempty"`
	// ChatIsCreated Optional. Pass true to request a chat owned by the user. Otherwise, no additional restrictions are applied
	ChatIsCreated bool `json:"chat_is_created,omitempty"`
	// UserRights Optional. Administrator rights required for the user in the chat
	UserRights Rights `json:"user_administrator_rights,omitempty"`
	// BotRights  Optional. Administrator rights required for the bot in the chat
	BotRights Rights `json:"bot_administrator_rights,omitempty"`
	// BotIsMember Optional. Pass true to request a chat with the bot as a member. Otherwise, no additional restrictions are applied
	BotIsMember bool `json:"bot_is_member,omitempty"`
}

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

// KeyboardButtonRequestUsers defines the criteria used to request suitable users.
type KeyboardButtonRequestUsers struct {
	// RequestID  Signed 32-bit identifier of the request
	RequestID int32 `json:"request_id"`
	// UserIsBot. Request bots if true, regular users if false
	UserIsBot *bool `json:"user_is_bot,omitempty"`
	// UserIsPremium. Request premium users if true, non-premium users if false
	UserIsPremium *bool `json:"user_is_premium,omitempty"`
	// MaxQuantity. Maximum number of users to be selected; 1-10. Defaults to 1.
	MaxQuantity *int `json:"max_quantity,omitempty"`
}

// KeyboardButtonPollType represents the type of a poll, which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	// Type
	// If quiz is passed, the user will be allowed to create only polls in the quiz mode.
	// If regular is passed,
	// only regular polls will be allowed.
	// Otherwise, the user will be allowed to create a poll of any type.
	Type *string `json:"type,omitempty"`
}

func (m *InlineKeyboardMarkup) ReplyMarkup() {}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to the message it belongs to.
type InlineKeyboardMarkup struct {
	// InlineKeyboard is an array of button rows, each represented by an array of InlineKeyboardButton objects.
	InlineKeyboard []Row `json:"inline_keyboard"`
}

func (m *InlineKeyboardMarkup) UnmarshalJSON(data []byte) error {
	var raw struct {
		InlineKeyboard []json.RawMessage `json:"inline_keyboard"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	m.InlineKeyboard = make([]Row, len(raw.InlineKeyboard))
	for i, row := range raw.InlineKeyboard {
		r := Row{}
		err := r.UnmarshalJSON(row)
		if err != nil {
			return err
		}
		m.InlineKeyboard[i] = r
	}
	return nil
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

func (m *ReplyKeyboardMarkup) UnmarshalJSON(data []byte) error {
	var raw struct {
		Keyboard              []json.RawMessage `json:"inline_keyboard"`
		IsPersistent          *bool             `json:"is_persistent,omitempty"`
		ResizeKeyboard        *bool             `json:"resize_keyboard,omitempty"`
		OneTimeKeyboard       *bool             `json:"one_time_keyboard,omitempty"`
		InputFieldPlaceholder *string           `json:"input_field_placeholder,omitempty"`
		Selective             *bool             `json:"selective,omitempty"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	m.Keyboard = make([]Row, len(raw.Keyboard))
	m.IsPersistent = raw.IsPersistent
	m.ResizeKeyboard = raw.ResizeKeyboard
	m.OneTimeKeyboard = raw.OneTimeKeyboard
	m.InputFieldPlaceholder = raw.InputFieldPlaceholder
	m.Selective = raw.Selective

	m.Keyboard = make([]Row, len(raw.Keyboard))
	for i, row := range raw.Keyboard {
		r := Row{}
		err := r.UnmarshalJSON(row)
		if err != nil {
			return err
		}
		m.Keyboard[i] = r
	}
	return nil
}
