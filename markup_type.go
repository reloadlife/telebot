package telebot

type MenuButton struct {
}

// InlineKeyboardButton represents one button of an inline keyboard. You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	// Text is the label text on the button.
	Text string `json:"text"`

	// URL is the optional HTTP or tg:// URL to be opened when the button is pressed.
	// Links tg://user?id=<user_id> can be used to mention a user by their identifier without using a username,
	// if this is allowed by their privacy settings.
	URL string `json:"url,omitempty"`

	// CallbackData is the optional data to be sent in a callback query to the bot when the button is pressed, 1-64 bytes.
	CallbackData string `json:"callback_data,omitempty"`

	// WebApp is the optional description of the Web App that will be launched when the user presses the button.
	// The Web App will be able to send an arbitrary message on behalf of the user using the method answerWebAppQuery.
	// Available only in private chats between a user and the bot.
	WebApp *WebAppInfo `json:"web_app,omitempty"`

	// LoginURL is the optional HTTPS URL used to automatically authorize the user.
	// Can be used as a replacement for the Telegram Login Widget.
	LoginURL *LoginUrl `json:"login_url,omitempty"`

	// SwitchInlineQuery is the optional inline query to prompt the user to select one of their chats and insert the bot's username and the specified inline query in the input field.
	// May be empty, in which case just the bot's username will be inserted.
	SwitchInlineQuery string `json:"switch_inline_query,omitempty"`

	// SwitchInlineQueryCurrentChat is the optional inline query to insert the bot's username and the specified inline query in the current chat's input field.
	// May be empty, in which case only the bot's username will be inserted.
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`

	// SwitchInlineQueryChosenChat is the optional inline query to prompt the user to select one of their chats of the specified type,
	// open that chat and insert the bot's username and the specified inline query in the input field.
	SwitchInlineQueryChosenChat *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`

	// CallbackGame is the optional description of the game that will be launched when the user presses the button.
	// NOTE: This type of button must always be the first button in the first row.
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`

	// Pay is the optional boolean to specify True, to send a Pay button.
	// NOTE: This type of button must always be the first button in the first row and can only be used in invoice messages.
	Pay bool `json:"pay,omitempty"`
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
	RequestContact bool `json:"request_contact,omitempty"`

	// RequestLocation, if True, the user's current location will be sent when the button is pressed.
	// Available in private chats only.
	RequestLocation bool `json:"request_location,omitempty"`

	// RequestPoll, if specified, the user will be asked to create a poll and send it to the bot when the button is pressed.
	// Available in private chats only.
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`

	// WebApp, if specified, the described Web App will be launched when the button is pressed.
	// The Web App will be able to send a “web_app_data” service message.
	// Available in private chats only.
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}
