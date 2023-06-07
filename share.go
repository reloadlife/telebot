package telebot

type KeyboardButtonRequestUser struct {
	RequestID     string `json:"request_id"`
	UserIsBot     bool   `json:"user_is_bot,omitempty"`
	UserIsPremium bool   `json:"user_is_premium,omitempty"`
}

type KeyboardButtonRequestChat struct {
	RequestID               string  `json:"request_id"`
	ChatIsChannel           bool    `json:"chat_is_channel,omitempty"`
	ChatIsForum             bool    `json:"chat_is_forum,omitempty"`
	ChatHasUsername         bool    `json:"chat_has_username,omitempty"`
	ChatIsCreated           bool    `json:"chat_is_created,omitempty"`
	UserAdministratorRights *Rights `json:"user_administrator_rights,omitempty"`
	BotAdministratorRights  *Rights `json:"bot_administrator_rights,omitempty"`
	BotIsMember             bool    `json:"bot_is_member,omitempty"`
}

type UserShared struct {
	ID     int   `json:"request_id"`
	UserID int64 `json:"user_id"`
}

type ChatShared struct {
	ID     int   `json:"request_id"`
	ChatID int64 `json:"chat_id"`
}
