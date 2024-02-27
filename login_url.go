package telebot

import "fmt"

// LoginURL represents a parameter of the inline keyboard button used to automatically authorize a user.
// Serves as a great replacement for the Telegram Login Widget when the user is coming from Telegram.
// All the user needs to do is tap/click a button and confirm that they want to log in.
type LoginURL struct {
	// URL is an HTTPS URL to be opened with user authorization data added to the query string when the button is pressed.
	// If the user refuses to provide authorization data, the original URL without information about the user will be opened.
	// The data added is the same as described in Receiving authorization data.
	URL string `json:"url" yaml:"url"`

	// ForwardText is the new text of the button in forwarded messages.
	ForwardText string `json:"forward_text,omitempty" yaml:"forward_text"`

	// BotUsername is the username of a bot, which will be used for user authorization.
	// See Setting up a bot for more details.
	// If not specified, the current bot's username will be assumed.
	// The URL's domain must be the same as the domain linked with the bot.
	// See Linking your domain to the bot for more details.
	BotUsername string `json:"bot_username,omitempty" yaml:"bot_username"`

	// RequestWriteAccess, if true, passes True to request the permission for your bot to send messages to the user.
	RequestWriteAccess bool `json:"request_write_access,omitempty" yaml:"request_write_access"`
}

func (c *LoginURL) Type() string        { return "login_url" }
func (c *LoginURL) ReflectType() string { return fmt.Sprintf("%T", c) }
