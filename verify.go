package telebot

type Verifiable interface {
	Verify() error
}
