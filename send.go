package telebot

type Sendable interface {
	Send(b Bot, to Recipient, opts ...any) (Message, error)
}
