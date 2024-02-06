package telebot

type MessageThreadID int

type forwardMessageRequest struct {
	ChatID              any              `json:"chat_id"`
	FromChat            string           `json:"from_chat_id"`
	MessageID           int              `json:"message_id,omitempty"`
	MessageIDs          []int            `json:"message_ids,omitempty"`
	MessageThreadID     *MessageThreadID `json:"message_thread_id,omitempty"`
	Protect             *bool            `json:"protect_content,omitempty"`
	DisableNotification *bool            `json:"disable_notification,omitempty"`
}

type copyMessageRequest struct {
	ChatID          any              `json:"chat_id"`
	FromChat        string           `json:"from_chat_id"`
	MessageID       int              `json:"message_id,omitempty"`
	MessageIDs      []int            `json:"message_ids,omitempty"`
	MessageThreadID *MessageThreadID `json:"message_thread_id,omitempty"`

	Caption             *string          `json:"caption,omitempty"`
	RemoveCaption       *bool            `json:"remove_caption,omitempty"`
	ParseMode           *ParseMode       `json:"parse_mode,omitempty"`
	CaptionEntities     []Entity         `json:"caption_entities,omitempty"`
	DisableNotification *bool            `json:"disable_notification,omitempty"`
	Protect             *bool            `json:"protect_content,omitempty"`
	ReplyParams         *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup         *ReplyMarkup     `json:"reply_markup,omitempty"`
}
