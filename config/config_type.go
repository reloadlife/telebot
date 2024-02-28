package config

import tele "go.mamad.dev/telebot"

type Config interface {
	L(locale, key string) string
	GetDefaultLocale() string
	GetLocales() []string

	GetSettings() Settings
	GetBot() BotConfig

	GetHandles() []Handle

	Keyboard(text string, l ...string) tele.ReplyMarkup
}
