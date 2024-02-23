package telebot

type sendDocumentParams struct {
	ChatID   any   `json:"chat_id"`
	Document *File `json:"document" file:"1"`

	ThreadID            *MessageThreadID `json:"message_thread_id,omitempty"`
	Caption             *string          `json:"caption,omitempty"`
	ParseMode           *ParseMode       `json:"parse_mode,omitempty"`
	Entities            []Entity         `json:"caption_entities,omitempty"`
	DisableNotification *bool            `json:"disable_notification,omitempty"`
	Protected           *bool            `json:"protect_content,omitempty"`
	ReplyParameters     *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup         ReplyMarkup      `json:"reply_markup,omitempty"`

	Thumbnail *File `json:"thumbnail,omitempty" file:"1"`

	DisableContentTypeDetection *bool `json:"disable_content_type_detection,omitempty"`
}
