package telebot

type sendInvoiceRequest struct {
	ChatID any `json:"chat_id"`

	Title         string         `json:"title"`
	Description   string         `json:"description"`
	Payload       string         `json:"payload"`
	ProviderToken string         `json:"provider_token"`
	Currency      string         `json:"currency"`
	Prices        []LabeledPrice `json:"prices"`

	MessageThreadID           *MessageThreadID `json:"message_thread_id,omitempty"`
	MaxTipAmount              *int             `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int            `json:"suggested_tip_amounts,omitempty"`
	StartParameter            *string          `json:"start_parameter,omitempty"`
	ProviderData              *string          `json:"provider_data,omitempty"`
	PhotoURL                  *string          `json:"photo_url,omitempty"`
	PhotoSize                 *int             `json:"photo_size,omitempty"`
	PhotoWidth                *int             `json:"photo_width,omitempty"`
	PhotoHeight               *int             `json:"photo_height,omitempty"`
	NeedName                  *bool            `json:"need_name,omitempty"`
	NeedPhoneNumber           *bool            `json:"need_phone_number,omitempty"`
	NeedEmail                 *bool            `json:"need_email,omitempty"`
	NeedShippingAddress       *bool            `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider *bool            `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       *bool            `json:"send_email_to_provider,omitempty"`
	IsFlexible                *bool            `json:"is_flexible,omitempty"`
	DisableNotification       *bool            `json:"disable_notification,omitempty"`
	ProtectContent            *bool            `json:"protect_content,omitempty"`
	ReplyParameters           *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup               ReplyMarkup      `json:"reply_markup,omitempty"`
}
