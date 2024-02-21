package telebot

type Option int

const (
	Silent Option = iota
	Protected
	HasSpoiler

	RemoveCaption

	DisableContentTypeDetection
	Streamable

	ResizeKeyboard
	OneTimeKeyboard
	PersistentKeyboard
	Selective
)
