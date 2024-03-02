package telebot

import "reflect"

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

	GetText() string
	SetText(string)

	Clone() Button
}

// Button is here so InlineKeyboardButton implements Button
func (*InlineKeyboardButton) Button() {}

// Button is here so InlineKeyboardButton implements Button
func (*KeyboardButton) Button() {}

func (i *InlineKeyboardButton) GetText() string  { return i.Text }
func (i *InlineKeyboardButton) SetText(t string) { i.Text = t }
func (i *KeyboardButton) GetText() string        { return i.Text }
func (i *KeyboardButton) SetText(t string)       { i.Text = t }

func (i *InlineKeyboardButton) Clone() Button {
	clone := reflect.New(reflect.TypeOf(i).Elem()).Interface()

	// Use reflection to copy the fields
	valueOfReceiver := reflect.ValueOf(i).Elem()
	valueOfClone := reflect.ValueOf(clone).Elem()

	for i := 0; i < valueOfReceiver.NumField(); i++ {
		field := valueOfReceiver.Field(i)
		valueOfClone.Field(i).Set(field)
	}

	return clone.(Button)
}

func (i *KeyboardButton) Clone() Button {
	clone := reflect.New(reflect.TypeOf(i).Elem()).Interface()

	// Use reflection to copy the fields
	valueOfReceiver := reflect.ValueOf(i).Elem()
	valueOfClone := reflect.ValueOf(clone).Elem()

	for i := 0; i < valueOfReceiver.NumField(); i++ {
		field := valueOfReceiver.Field(i)
		valueOfClone.Field(i).Set(field)
	}

	return clone.(Button)
}
