package telebot

// ShippingQuery contains information about an incoming shipping query.
type ShippingQuery struct {
	// ID is the unique query identifier.
	ID string `json:"id"`

	// From is the user who sent the query.
	From User `json:"from"`

	// InvoicePayload is the bot-specified invoice payload.
	InvoicePayload string `json:"invoice_payload"`

	// ShippingAddress is the user-specified shipping address.
	ShippingAddress ShippingAddress `json:"shipping_address"`
}

// PreCheckoutQuery contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	// ID is the unique query identifier.
	ID string `json:"id"`

	// From is the user who sent the query.
	From User `json:"from"`

	// Currency is the three-letter ISO 4217 currency code.
	Currency string `json:"currency"`

	// TotalAmount is the total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45, pass amount = 145.
	// See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// InvoicePayload is the bot-specified invoice payload.
	InvoicePayload string `json:"invoice_payload"`

	// ShippingOptionID is the identifier of the shipping option chosen by the user (optional).
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// OrderInfo is the order information provided by the user (optional).
	OrderInfo OrderInfo `json:"order_info,omitempty"`
}
