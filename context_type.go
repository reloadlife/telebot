package telebot

// Context wraps an update and represents the context of the current event.
type Context interface {
	// Bot returns the bot instance.
	Bot() Bot

	// Update returns the original update.
	Update() Update

	// Message returns stored *AccessibleMessage if such presented.
	Message() *AccessibleMessage

	// Callback returns stored callback if such presented.
	Callback() *Callback

	// Query returns stored query if such presented.
	Query() *InlineQuery

	// Get retrieves data from the context.
	Get(key string) any

	// Set saves data in the context.
	Set(key string, val any)

	ReplyTo(msg Message, text string, options ...any) (Message, error)
	Reply(text string, options ...any) (Message, error)

	Send(s any, options ...any) (Message, error)

	Text() string

	Sender() *User
}
