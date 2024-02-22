package telebot

type Width int
type Height int

type sendAnimationRequest struct {
	ChatID    any   `json:"chat_id"`
	Animation *File `json:"animation" file:"1"`

	ThreadID            *MessageThreadID `json:"message_thread_id,omitempty"`
	Caption             *string          `json:"caption,omitempty"`
	ParseMode           *ParseMode       `json:"parse_mode,omitempty"`
	Entities            []Entity         `json:"caption_entities,omitempty"`
	HasSpoiler          *bool            `json:"has_spoiler,omitempty"`
	DisableNotification *bool            `json:"disable_notification,omitempty"`
	Protected           *bool            `json:"protect_content,omitempty"`
	ReplyParameters     *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup         *ReplyMarkup     `json:"reply_markup,omitempty"`

	Duration  *int  `json:"duration,omitempty"`
	Width     *int  `json:"width,omitempty"`
	Height    *int  `json:"height,omitempty"`
	Thumbnail *File `json:"thumbnail,omitempty"`
}
