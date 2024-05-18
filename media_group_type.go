package telebot

type sendMediaGroupRequest struct {
	ChatID any          `json:"chat_id"`
	Media  []InputMedia `json:"media"`

	BusinessID          *BusinessID      `json:"business_connection_id,omitempty"`
	ThreadID            *MessageThreadID `json:"message_thread_id,omitempty"`
	DisableNotification *bool            `json:"disable_notification,omitempty"`
	Protected           *bool            `json:"protect_content,omitempty"`
	ReplyParameters     *ReplyParameters `json:"reply_parameters,omitempty"`
}
