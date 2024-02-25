package telebot

import "fmt"

type answerShippingQueryRequest struct {
	ID              string           `json:"shipping_query_id"`
	OK              bool             `json:"ok"`
	ShippingOptions []ShippingOption `json:"shipping_options,omitempty"`
	ErrorMessage    string           `json:"error_message,omitempty"`
}

// ShippingAddress represents a shipping address.
type ShippingAddress struct {
	// CountryCode is the two-letter ISO 3166-1 alpha-2 country code.
	CountryCode string `json:"country_code"`

	// State is the state, if applicable.
	State string `json:"state"`

	// City is the city.
	City string `json:"city"`

	// StreetLine1 is the first line for the address.
	StreetLine1 string `json:"street_line1"`

	// StreetLine2 is the second line for the address.
	StreetLine2 string `json:"street_line2"`

	// PostCode is the address post code.
	PostCode string `json:"post_code"`
}

func (c *ShippingAddress) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *ShippingAddress) Type() string        { return "ShippingAddress" }

// OrderInfo represents information about an order.
type OrderInfo struct {
	// Name is the user's name (optional).
	Name string `json:"name,omitempty"`

	// PhoneNumber is the user's phone number (optional).
	PhoneNumber string `json:"phone_number,omitempty"`

	// Email is the user's email (optional).
	Email string `json:"email,omitempty"`

	// ShippingAddress is the user's shipping address (optional).
	ShippingAddress ShippingAddress `json:"shipping_address,omitempty"`
}

func (c *OrderInfo) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *OrderInfo) Type() string        { return "OrderInfo" }

type ShippingOption struct {
	ID     string         `json:"id"`
	Title  string         `json:"title"`
	Prices []LabeledPrice `json:"prices"`
}
