package telebot

import "time"

type Bot interface {
	// Debug sends a debug message to the bot.
	Debug(...any)

	// OnError Debug sends a debug message to the bot.
	OnError(error, Context)

	// Close this method to close the bot instance
	// before moving it from one local server to another.
	// You need to delete the webhook before calling this
	// method to ensure that the bot isn't launched again after the server restarts.
	// The method will return error 429 in the first 10 minutes after the bot is launched.
	Close() error

	// Logout Use this method to log out from the cloud
	// Bot API server before launching the bot locally.
	// You must log out the bot before running it locally,
	// otherwise there is no guarantee that the bot will receive updates.
	// After a successful call,
	// you can immediately log in on a local server,
	// but will not be able to log in back to the cloud Bot API server for 10 minutes.
	Logout() error

	// GetUpdates returns a list of updates (Long Polling).
	GetUpdates(offset, limit int, timeout time.Duration, allowed ...UpdateType) ([]Update, error)

	// GetMe returns basic information about the bot as a User object.
	GetMe() (*User, error)

	// SendMessage sends a text message.
	SendMessage(recipient Recipient, text string, options ...Option) (*Message, error)

	// Ban a user from a group, a supergroup or a channel.
	Ban(chatID Recipient, userID int64, untilDate *int64, revokeMessages *bool) error

	// Unban a user from a group, a supergroup or a channel.
	Unban(chatID Recipient, userID int64, onlyIfBanned *bool) error

	// Handle Register Routes
	Handle(endpoint any, h HandlerFunc, m ...MiddlewareFunc)

	Start()
	Stop()

	StartInWebhook()
	StopInWebhook()
}
