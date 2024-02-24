package telebot

type sendLocationRequest struct {
	ChatID    any     `json:"chat_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	ThreadID             *MessageThreadID `json:"message_thread_id,omitempty"`
	HorizontalAccuracy   *float64         `json:"horizontal_accuracy,omitempty"`
	LivePeriod           *int             `json:"live_period,omitempty"`
	Heading              *int             `json:"heading,omitempty"`
	ProximityAlertRadius *int             `json:"proximity_alert_radius,omitempty"`

	DisableNotification *bool            `json:"disable_notification,omitempty"`
	Protected           *bool            `json:"protect_content,omitempty"`
	ReplyParameters     *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup         ReplyMarkup      `json:"reply_markup,omitempty"`
}
