package telebot

import "fmt"

type InputMessageContent interface {
	JSONer
	InputMessageContent()
}

type InputTextMessageContent struct {
	MessageText        string             `json:"message_text"`
	ParseMode          string             `json:"parse_mode,omitempty"`
	Entities           []Entity           `json:"entities,omitempty"`
	LinkPreviewOptions LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

func (c *InputTextMessageContent) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *InputTextMessageContent) Type() string        { return "InputTextMessageContent" }

type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

func (c *InputLocationMessageContent) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *InputLocationMessageContent) Type() string        { return "InputLocationMessageContent" }

type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    string  `json:"foursquare_id,omitempty"`
	FoursquareType  string  `json:"foursquare_type,omitempty"`
	GooglePlaceID   string  `json:"google_place_id,omitempty"`
	GooglePlaceType string  `json:"google_place_type,omitempty"`
}

func (c *InputVenueMessageContent) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *InputVenueMessageContent) Type() string        { return "InputVenueMessageContent" }

type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	VCard       string `json:"vcard,omitempty"`
}

func (c *InputContactMessageContent) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *InputContactMessageContent) Type() string        { return "InputContactMessageContent" }

type InputInvoiceMessageContent struct {
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             string         `json:"provider_token"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              int            `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int          `json:"suggested_tip_amounts,omitempty"`
	ProviderData              string         `json:"provider_data,omitempty"`
	PhotoURL                  string         `json:"photo_url,omitempty"`
	PhotoSize                 int            `json:"photo_size,omitempty"`
	PhotoWidth                int            `json:"photo_width,omitempty"`
	PhotoHeight               int            `json:"photo_height,omitempty"`
	NeedName                  bool           `json:"need_name,omitempty"`
	NeedPhoneNumber           bool           `json:"need_phone_number,omitempty"`
	NeedEmail                 bool           `json:"need_email,omitempty"`
	NeedShippingAddress       bool           `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool           `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool           `json:"is_flexible,omitempty"`
}

func (c *InputInvoiceMessageContent) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *InputInvoiceMessageContent) Type() string        { return "InputInvoiceMessageContent" }

func (i *InputTextMessageContent) InputMessageContent()     {}
func (i *InputLocationMessageContent) InputMessageContent() {}
func (i *InputVenueMessageContent) InputMessageContent()    {}
func (i *InputContactMessageContent) InputMessageContent()  {}
func (i *InputInvoiceMessageContent) InputMessageContent()  {}
