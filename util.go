package telebot

func toPtr[T any](v T) *T {
	return &v
}
