package config

func (c *config) GetDefaultLocale() string {
	return string(c.defaultLocale)
}

func (c *config) GetLocales() []string {
	return c.conf.Locales
}

func (c *config) l(locale, key string) string {
	lMap, ok := c.locales[localeKey(locale)]
	if !ok {
		lMap = c.locales[c.defaultLocale]
	}

	lResp, has := lMap[key]
	if !has {
		panic("telebot-config: key, " + key + ", not found in locale, " + locale + ", or default locale")
	}
	return lResp
}
