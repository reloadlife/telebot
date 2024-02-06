package telebot

type Option int

const (
	Silent Option = iota
	Protected
	HasSpoiler

	RemoveCaption
)
