package config

import (
	"bytes"
	"fmt"
	"text/template"
)

func (c *config) GetDefaultLocale() string {
	return string(c.defaultLocale)
}

func (c *config) GetLocales() []string {
	return c.conf.Locales
}

func (c *config) L(locale, key string, args ...any) string {
	lMap, ok := c.locales[localeKey(locale)]
	if !ok {
		lMap = c.locales[c.defaultLocale]
	}
	lResp, has := lMap[key]
	if !has {
		panic("telebot-config: key, " + key + ", not found in locale, " + locale + ", or default locale")
	}
	_tpl := template.New(key)
	_tpl.Funcs(c.mapFunctions)
	tpl, err := _tpl.Parse(lResp)
	if err != nil {
		panic(fmt.Errorf("telebot-config/text: %w", err))
	}
	var arg any
	if len(args) > 0 {
		arg = args[0]
	}
	var buf bytes.Buffer
	if err := c.parse(tpl, locale).ExecuteTemplate(&buf, key, arg); err != nil {
		panic(fmt.Errorf("telebot-config/text: %w", err))
	}
	return buf.String()
}
