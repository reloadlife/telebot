package telebot

type Performer string
type Title string

type sendAudioRequest struct {
	ChatID any   `json:"chat_id"`
	Audio  *File `json:"audio" file:"1"`

	ThreadID            *MessageThreadID `json:"message_thread_id,omitempty"`
	Caption             *string          `json:"caption,omitempty"`
	ParseMode           *ParseMode       `json:"parse_mode,omitempty"`
	Entities            []Entity         `json:"caption_entities,omitempty"`
	DisableNotification *bool            `json:"disable_notification,omitempty"`
	Protected           *bool            `json:"protect_content,omitempty"`
	ReplyParameters     *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup         *ReplyMarkup     `json:"reply_markup,omitempty"`

	Duration  *int    `json:"duration,omitempty"`
	Performer *string `json:"performer,omitempty"`
	Title     *string `json:"title,omitempty"`
	Thumbnail *File   `json:"thumbnail,omitempty"`
}
