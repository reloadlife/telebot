package telebot

// SuccessfulPayment contains basic information about a successful payment.
type SuccessfulPayment struct {
	// Currency is the three-letter ISO 4217 currency code.
	Currency string `json:"currency"`

	// TotalAmount is the total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in currencies.json; it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`

	// InvoicePayload is the bot-specified invoice payload.
	InvoicePayload string `json:"invoice_payload"`

	// ShippingOptionID is an optional field representing the identifier of the shipping option chosen by the user.
	ShippingOptionID string `json:"shipping_option_id,omitempty"`

	// OrderInfo is an optional field representing order information provided by the user.
	OrderInfo OrderInfo `json:"order_info,omitempty"`

	// TelegramPaymentChargeID is the Telegram payment identifier.
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`

	// ProviderPaymentChargeID is the provider payment identifier.
	ProviderPaymentChargeID string `json:"provider_payment_charge_id"`
}
