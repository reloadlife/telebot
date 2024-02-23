package telebot

type sendVideoNoteRequest struct {
	ChatID    any   `json:"chat_id"`
	VideoNote *File `json:"video_note" file:"1"`

	ThreadID            *MessageThreadID `json:"message_thread_id,omitempty"`
	DisableNotification *bool            `json:"disable_notification,omitempty"`
	Protected           *bool            `json:"protect_content,omitempty"`
	ReplyParameters     *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup         ReplyMarkup      `json:"reply_markup,omitempty"`

	Thumbnail *File `json:"thumbnail,omitempty" file:"1"`
	Duration  *int  `json:"duration,omitempty"`
	Length    *int  `json:"length,omitempty"`
}
