package telebot

type sendMessageParams struct {
	ChatID              any                 `json:"chat_id" validate:"required"`
	MessageThreadID     *MessageThreadID    `json:"message_thread_id,omitempty"`
	Text                string              `json:"text" validate:"required,max=4096"`
	ParseMode           *ParseMode          `json:"parse_mode,omitempty"`
	Entities            []Entity            `json:"entities,omitempty"`
	LinkPreviewOptions  *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	DisableNotification *bool               `json:"disable_notification,omitempty"`
	ProtectContent      *bool               `json:"protect_content,omitempty"`
	ReplyParameters     *ReplyParameters    `json:"reply_parameters,omitempty"`
	ReplyMarkup         ReplyMarkup         `json:"reply_markup,omitempty"`
}
