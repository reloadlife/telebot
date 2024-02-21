package telebot

type Row []Button

type Button interface {
	Button()

	String() string

	MarshalJSON() ([]byte, error)
	UnmarshalJSON(data []byte) error

	deepEqual(other any) bool // private for internal use
}

func (*InlineKeyboardButton) Button() {}
func (*KeyboardButton) Button()       {}
