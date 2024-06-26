package telebot

type sendVoiceRequest struct {
	ChatID any   `json:"chat_id"`
	Voice  *File `json:"voice" file:"1"`

	BusinessID          *BusinessID      `json:"business_connection_id,omitempty"`
	ThreadID            *MessageThreadID `json:"message_thread_id,omitempty"`
	Caption             *string          `json:"caption,omitempty"`
	ParseMode           ParseMode        `json:"parse_mode,omitempty"`
	Entities            []Entity         `json:"caption_entities,omitempty"`
	DisableNotification *bool            `json:"disable_notification,omitempty"`
	Protected           *bool            `json:"protect_content,omitempty"`
	ReplyParameters     *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup         ReplyMarkup      `json:"reply_markup,omitempty"`

	Duration *int `json:"duration,omitempty"`
}
