package telebot

// Row Represents a row of Buttons
// Maximum number of Buttons in a ReplyMarkup is 100 (telegram limits).
type Row []Button

// Button is the interface for all buttons
// It is implemented by InlineKeyboardButton and KeyboardButton
type Button interface {
	Button()

	String() string

	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error

	deepEqual(other any) bool // private for internal use
}

// Button is here so InlineKeyboardButton implements Button
func (*InlineKeyboardButton) Button() {}

// Button is here so InlineKeyboardButton implements Button
func (*KeyboardButton) Button() {}
