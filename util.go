package telebot

import "os"

func toPtr[T any](v T) *T {
	return &v
}

func readFromEnv(key string) string {
	return os.Getenv(key)
}
