package telebot

// todo: better documentation and file naming.

import (
	"fmt"
)

// SwitchInlineQueryChosenChat represents an inline button that switches the current user to inline mode
// in a chosen chat, with an optional default inline query.
type SwitchInlineQueryChosenChat struct {
	// Query is the default inline query to be inserted in the input field.
	// If left empty, only the bot's username will be inserted.
	Query *string `json:"query,omitempty" yaml:"query,omitempty"`

	// AllowUserChats, if true, allows private chats with users to be chosen.
	AllowUserChats *bool `json:"allow_user_chats,omitempty" yaml:"allow_user_chats,omitempty"`

	// AllowBotChats, if true, allows private chats with bots to be chosen.
	AllowBotChats *bool `json:"allow_bot_chats,omitempty" yaml:"allow_bot_chats,omitempty"`

	// AllowGroupChats, if true, allows group and supergroup chats to be chosen.
	AllowGroupChats *bool `json:"allow_group_chats,omitempty" yaml:"allow_group_chats,omitempty"`

	// AllowChannelChats, if true, allows channel chats to be chosen.
	AllowChannelChats *bool `json:"allow_channel_chats,omitempty" yaml:"allow_channel_chats,omitempty"`
}

// KeyboardButtonRequestChat defines the criteria used to request a suitable chat.
type KeyboardButtonRequestChat struct {
	// RequestID Signed 32-bit identifier of the request
	RequestID int32 `json:"request_id" yaml:"request_id"`
	// ChatIsChannel Pass true to request a channel chat, pass false to request a group or a supergroup chat
	ChatIsChannel bool `json:"chat_is_channel" yaml:"chat_is_channel"`
	// ChatIsForum  Optional. Pass true to request a forum supergroup, pass false to request a non-forum chat
	ChatIsForum bool `json:"chat_is_forum,omitempty" yaml:"chat_is_forum,omitempty"`
	// ChatHasUsernameOptional. Pass true to request a supergroup or a channel with a username, pass false to request a chat without a username
	ChatHasUsername bool `json:"chat_has_username,omitempty" yaml:"chat_has_username,omitempty"`
	// ChatIsCreated Optional. Pass true to request a chat owned by the user. Otherwise, no additional restrictions are applied
	ChatIsCreated bool `json:"chat_is_created,omitempty" yaml:"chat_is_created,omitempty"`
	// UserRights Optional. Administrator rights required for the user in the chat
	UserRights Rights `json:"user_administrator_rights,omitempty" yaml:"user_rights,omitempty"`
	// BotRights  Optional. Administrator rights required for the bot in the chat
	BotRights Rights `json:"bot_administrator_rights,omitempty" yaml:"bot_rights,omitempty"`
	// BotIsMember Optional. Pass true to request a chat with the bot as a member. Otherwise, no additional restrictions are applied
	BotIsMember bool `json:"bot_is_member,omitempty" yaml:"bot_is_member,omitempty"`
}

// KeyboardButtonRequestUsers defines the criteria used to request suitable users.
type KeyboardButtonRequestUsers struct {
	// RequestID  Signed 32-bit identifier of the request
	RequestID int32 `json:"request_id" yaml:"request_id"`
	// UserIsBot. Request bots if true, regular users if false
	UserIsBot *bool `json:"user_is_bot,omitempty" yaml:"user_is_bot,omitempty"`
	// UserIsPremium. Request premium users if true, non-premium users if false
	UserIsPremium *bool `json:"user_is_premium,omitempty" yaml:"user_is_premium,omitempty"`
	// MaxQuantity. Maximum number of users to be selected; 1-10. Defaults to 1.
	MaxQuantity *int `json:"max_quantity,omitempty" yaml:"max_quantity,omitempty"`
}

// KeyboardButtonPollType represents the type of a poll, which is allowed to be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	// Type
	// If quiz is passed, the user will be allowed to create only polls in the quiz mode.
	// If regular is passed,
	// only regular polls will be allowed.
	// Otherwise, the user will be allowed to create a poll of any type.
	PollType PollType `json:"type,omitempty"`
}

// Actual Button Types =>

// InlineKeyboardButton represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	// Text is the label text on the button.
	Text string `json:"text"`

	// URL is the optional HTTP or tg:// URL to be opened when the button is pressed.
	// Links tg://user?id=<user_id> can be used to mention a user by their identifier without using a username,
	// if this is allowed by their privacy settings.
	URL *string `json:"url,omitempty"`

	// CallbackData is the optional data to be sent in a callback query to the bot when the button is pressed, 1-64 bytes.
	CallbackData *string `json:"callback_data,omitempty"`

	// WebApp is the optional description of the Web App that will be launched when the user presses the button.
	// The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery.
	// Available only in private chats between a user and the bot.
	WebApp *WebAppInfo `json:"web_app,omitempty"`

	// LoginURL is the optional HTTPS URL used to automatically authorize the user.
	// Can be used as a replacement for the Telegram Login Widget.
	LoginURL *LoginURL `json:"login_url,omitempty"`

	// SwitchInlineQuery is the optional inline query to prompt the user to select one of their chats and insert the bot's username and the specified inline query in the input field.
	// May be empty, in which case just the bot's username will be inserted.
	SwitchInlineQuery *string `json:"switch_inline_query,omitempty"`

	// SwitchInlineQueryCurrentChat is the optional inline query to insert the bot's username and the specified inline query in the current chat's input field.
	// May be empty, in which case only the bot's username will be inserted.
	SwitchInlineQueryCurrentChat *string `json:"switch_inline_query_current_chat,omitempty"`

	// SwitchInlineQueryChosenChat is the optional inline query to prompt the user to select one of their chats of the specified type,
	// open that chat and insert the bot's username and the specified inline query in the input field.
	SwitchInlineQueryChosenChat *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`

	// CallbackGame is the optional description of the game that will be launched when the user presses the button.
	// NOTE: This type of button must always be the first button in the first row.
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`

	// Pay is the optional boolean to specify True, to send a Pay button.
	// NOTE: This type of button must always be the first button in the first row and can only be used in invoice messages.
	Pay *bool `json:"pay,omitempty"`
}

// KeyboardButton represents one button of the reply keyboard. For simple text buttons, String can be used instead of this object to specify the button text.
// The optional fields web_app, request_users, request_chat, request_contact, request_location, and request_poll are mutually exclusive.
type KeyboardButton struct {
	// Text is the text of the button. If none of the optional fields are used, it will be sent as a message when the button is pressed.
	Text string `json:"text"`

	// RequestUsers, if specified, pressing the button will open a list of suitable users.
	// Identifiers of selected users will be sent to the bot in a “users_shared” service message.
	// Available in private chats only.
	RequestUsers *KeyboardButtonRequestUsers `json:"request_users,omitempty"`

	// RequestChat, if specified, pressing the button will open a list of suitable chats.
	// Tapping on a chat will send its identifier to the bot in a “chat_shared” service message.
	// Available in private chats only.
	RequestChat *KeyboardButtonRequestChat `json:"request_chat,omitempty"`

	// RequestContact, if True, the user's phone number will be sent as a contact when the button is pressed.
	// Available in private chats only.
	RequestContact *bool `json:"request_contact,omitempty"`

	// RequestLocation, if True, the user's current location will be sent when the button is pressed.
	// Available in private chats only.
	RequestLocation *bool `json:"request_location,omitempty"`

	// RequestPoll, if specified, the user will be asked to create a poll and send it to the bot when the button is pressed.
	// Available in private chats only.
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`

	// WebApp, if specified, the described Web App will be launched when the button is pressed.
	// The Web App will be able to send a “web_app_data” service message.
	// Available in private chats only.
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

func (i *InlineKeyboardButton) Type() string               { return "InlineKeyboardButton" }
func (i *InlineKeyboardButton) ReflectType() string        { return fmt.Sprintf("%T", i) }
func (i *KeyboardButton) Type() string                     { return "KeyboardButton" }
func (i *KeyboardButton) ReflectType() string              { return fmt.Sprintf("%T", i) }
func (i *SwitchInlineQueryChosenChat) Type() string        { return "SwitchInlineQueryChosenChat" }
func (i *SwitchInlineQueryChosenChat) ReflectType() string { return fmt.Sprintf("%T", i) }
func (i *KeyboardButtonRequestChat) Type() string          { return "KeyboardButtonRequestChat" }
func (i *KeyboardButtonRequestChat) ReflectType() string   { return fmt.Sprintf("%T", i) }
func (i *KeyboardButtonRequestUsers) Type() string         { return "KeyboardButtonRequestUsers" }
func (i *KeyboardButtonRequestUsers) ReflectType() string  { return fmt.Sprintf("%T", i) }
func (i *KeyboardButtonPollType) Type() string {
	if i.PollType == "" {
		return Unknown
	}
	return string(i.PollType)
}
func (i *KeyboardButtonPollType) ReflectType() string { return fmt.Sprintf("%T", i) }
