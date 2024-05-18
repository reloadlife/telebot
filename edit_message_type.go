package telebot

type editMessageParams struct {
	ChatID          any    `json:"chat_id"`
	MessageID       int64  `json:"message_id"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
	ReplyMarkup     ReplyMarkup
}

type editMessageTextParams struct {
	ChatID          any         `json:"chat_id"`
	MessageID       int64       `json:"message_id"`
	InlineMessageID string      `json:"inline_message_id,omitempty"`
	ReplyMarkup     ReplyMarkup `json:"reply_markup,omitempty"`

	Text               string              `json:"text" validate:"required,max=4096"`
	ParseMode          ParseMode           `json:"parse_mode,omitempty"`
	Entities           []Entity            `json:"entities,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

type editMessageCaptionParams struct {
	ChatID          any         `json:"chat_id"`
	MessageID       int64       `json:"message_id"`
	InlineMessageID string      `json:"inline_message_id,omitempty"`
	ReplyMarkup     ReplyMarkup `json:"reply_markup,omitempty"`

	Caption   string    `json:"text" validate:"required,max=4096"`
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	Entities  []Entity  `json:"caption_entities,omitempty"`
}

type editMessageMediaParams struct {
	ChatID          any         `json:"chat_id"`
	MessageID       int64       `json:"message_id"`
	InlineMessageID string      `json:"inline_message_id,omitempty"`
	ReplyMarkup     ReplyMarkup `json:"reply_markup,omitempty"`

	Media InputMedia `json:"media"`
}

type editMessageLiveLocationParams struct {
	ChatID          any         `json:"chat_id"`
	MessageID       int64       `json:"message_id"`
	InlineMessageID string      `json:"inline_message_id,omitempty"`
	ReplyMarkup     ReplyMarkup `json:"reply_markup,omitempty"`

	// LivePeriod new period in seconds during which the location can be updated,
	// starting from the message send date. If 0x7FFFFFFF is specified, then the location can be updated forever.
	// Otherwise, the new value must not exceed the current live_period by more than a day,
	// and the live location expiration date must remain within the next 90 days.
	// If not specified, then live_period remains unchanged
	LivePeriod *int `json:"live_period,omitempty"`

	Latitude  float64 `json:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" validate:"required"`

	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

type deleteMessageParams struct {
	ChatID     any     `json:"chat_id"`
	MessageID  int64   `json:"message_id"`
	MessageIDs []int64 `json:"message_ids,omitempty"`
}
