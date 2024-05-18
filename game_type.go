package telebot

type sendGameRequest struct {
	ChatID        any    `json:"chat_id"`
	GameShortName string `json:"game_short_name"`

	BusinessID          *BusinessID      `json:"business_connection_id,omitempty"`
	ThreadID            *MessageThreadID `json:"message_thread_id,omitempty"`
	DisableNotification *bool            `json:"disable_notification,omitempty"`
	Protected           *bool            `json:"protect_content,omitempty"`
	ReplyParameters     *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup         ReplyMarkup      `json:"reply_markup,omitempty"`
}

type setGameScore struct {
	ChatID          any    `json:"chat_id"`
	MessageID       int64  `json:"message_id"`
	InlineMessageID string `json:"inline_message_id,omitempty"`

	DisableEditMessage *bool `json:"disable_edit_message,omitempty"`

	UserID any   `json:"user_id"`
	Score  uint  `json:"score"`
	Force  *bool `json:"force,omitempty"`
}

type getGameHighScores struct {
	UserID any `json:"user_id"`

	ChatID          any    `json:"chat_id"`
	MessageID       int64  `json:"message_id"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
}
