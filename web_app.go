package telebot

// WebAppData describes data sent from a Web App to the bot.
type WebAppData struct {
	// Data is the data sent from the Web App. Be aware that a bad client can send arbitrary data in this field.
	Data string `json:"data"`

	// ButtonText is the text of the web_app keyboard button from which the Web App was opened.
	// Be aware that a bad client can send arbitrary data in this field.
	ButtonText string `json:"button_text"`
}

// WebAppInfo describes a Web App.
type WebAppInfo struct {
	// URL is an HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps.
	URL string `json:"url"`
}

// LoginUrl represents a parameter of the inline keyboard button used to automatically authorize a user.
// Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram.
// All the user needs to do is tap/click a button and confirm that they want to log in.
type LoginUrl struct {
	// URL is an HTTPS URL to be opened with user authorization data added to the query string when the button is pressed.
	// If the user refuses to provide authorization data, the original URL without information about the user will be opened.
	// The data added is the same as described in Receiving authorization data.
	URL string `json:"url"`

	// ForwardText is the new text of the button in forwarded messages.
	ForwardText string `json:"forward_text,omitempty"`

	// BotUsername is the username of a bot, which will be used for user authorization.
	// See Setting up a bot for more details.
	// If not specified, the current bot's username will be assumed.
	// The URL's domain must be the same as the domain linked with the bot.
	// See Linking your domain to the bot for more details.
	BotUsername string `json:"bot_username,omitempty"`

	// RequestWriteAccess, if true, passes True to request the permission for your bot to send messages to the user.
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
}

type SentWebAppMessage struct {
	InlineMessageID *string `json:"inline_message_id,omitempty"`
}
