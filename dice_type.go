package telebot

type sendDiceRequest struct {
	ChatID any `json:"chat_id"`

	Emoji DiceEmoji `json:"emoji,omitempty"`

	BusinessID          *BusinessID      `json:"business_connection_id,omitempty"`
	ThreadID            *MessageThreadID `json:"message_thread_id,omitempty"`
	DisableNotification *bool            `json:"disable_notification,omitempty"`
	Protected           *bool            `json:"protect_content,omitempty"`
	ReplyParameters     *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup         ReplyMarkup      `json:"reply_markup,omitempty"`
}
