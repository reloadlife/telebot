package telebot

// Story object represents a message about a forwarded story in the chat
type Story struct {
	ID   int64 `json:"id"`
	Chat *Chat `json:"chat"`
}
