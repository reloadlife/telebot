package telebot

// Contact represents a phone contact.
type Contact struct {
	// PhoneNumber is the contact's phone number.
	PhoneNumber string `json:"phone_number"`

	// FirstName is the contact's first name.
	FirstName string `json:"first_name"`

	// LastName is an optional field representing the contact's last name.
	LastName *string `json:"last_name,omitempty"`

	// UserID is an optional field representing the contact's user identifier in Telegram.
	// This number may have more than 32 significant bits, and some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a 64-bit integer or double-precision float type is safe for storing this identifier.
	UserID *int64 `json:"user_id,omitempty"`

	// VCard is an optional field providing additional data about the contact in the form of a vCard.
	VCard *string `json:"vcard,omitempty"`
}
