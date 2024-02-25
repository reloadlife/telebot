package telebot

import "fmt"

// Invoice contains basic information about an invoice.
type Invoice struct {
	// Title is the product name.
	Title string `json:"title"`

	// Description is the product description.
	Description string `json:"description"`

	// StartParameter is a unique bot deep-linking parameter that can be used to generate this invoice.
	StartParameter string `json:"start_parameter"`

	// Currency is the three-letter ISO 4217 currency code.
	Currency string `json:"currency"`

	// TotalAmount is the total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in currencies.json; it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
}

func (i *Invoice) ReflectType() string { return fmt.Sprintf("%T", i) }
func (i *Invoice) Type() string        { return "Invoice" }
