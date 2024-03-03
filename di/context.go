package di

import (
	tele "go.mamad.dev/telebot"
	"go.mamad.dev/telebot/config"
)

type Ctx interface {
	GetLocales() []string

	Text(key string, l string, args ...any) string
	Keyboard(text string, l ...string) tele.ReplyMarkup
}

type context struct {
	locales []string
	conf    config.Config
}

func (c *context) GetLocales() []string { return c.locales }
func (c *context) Text(key string, l string, args ...any) string {
	return c.conf.L(l, key, args...)
}
func (c *context) Keyboard(text string, l ...string) tele.ReplyMarkup {
	locale := c.conf.GetDefaultLocale()
	if len(l) > 0 {
		locale = l[0]
	}
	return c.conf.Keyboard(text, locale)
}
