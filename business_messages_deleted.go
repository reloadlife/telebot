package telebot

type BusinessMessagesDeleted struct {
	// ID Unique identifier of the business connection
	ID string `json:"business_connection_id"`

	// Chat Information about a chat in the business account. The bot may not have access to the chat or the corresponding user.
	Chat *Chat `json:"chat"`

	// MessageIds The list of identifiers of deleted messages in the chat of the business account
	MessageIds []int64 `json:"message_ids"`
}
