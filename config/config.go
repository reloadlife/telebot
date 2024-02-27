package config

import (
	tele "go.mamad.dev/gtb"
	"go.mamad.dev/gtb/log"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type config struct {
	filePath string

	locales       localeMap
	defaultLocale localeKey

	conf *parsedConf

	keyboards Keyboards
}

func NewConfigFromFile(path string) Config {
	yamlRaw, err := os.ReadFile(path)
	if err != nil {
		panic("telebot-config: failed to read config file, " + err.Error())
	}

	c := &config{
		filePath: path,
		conf:     &parsedConf{},
		locales:  make(localeMap),
	}

	err = yaml.Unmarshal(yamlRaw, c.conf)
	if err != nil {
		panic("telebot-config: failed to parse config file, " + err.Error())
	}

	parseEnvironmentIntoConfig(c.conf)

	localePath := c.conf.Settings.RootDir + c.conf.Settings.LocalesDir

	loadLocales(localePath, c, c.conf.Locales...)

	c.defaultLocale = localeKey(c.conf.Locales[0])

	_k := c.GetKeyboards()
	c.keyboards = _k

	return c
}

func parseEnvironmentIntoConfig(c *parsedConf) {
	if strings.HasPrefix(c.Settings.Token, "env:") {
		envVar := strings.TrimPrefix(c.Settings.Token, "env:")
		token, ok := os.LookupEnv(envVar)
		c.Settings.Token = token
		if !ok {
			panic("telebot-config: failed to get token from environment, variable, " + envVar + " not found")
		}
	}

	if strings.HasPrefix(c.Settings.URL, "env:") {
		envVar := strings.TrimPrefix(c.Settings.URL, "env:")
		url, ok := os.LookupEnv(envVar)
		c.Settings.URL = url
		if !ok {
			panic("telebot-config: failed to get url from environment, variable, " + envVar + " not found")
		}
	}

	if strings.HasPrefix(c.Settings.RootDir, "env:") {
		envVar := strings.TrimPrefix(c.Settings.RootDir, "env:")
		rootDir, ok := os.LookupEnv(envVar)
		c.Settings.RootDir = rootDir
		if !ok {
			panic("telebot-config: failed to get rootDir from environment, variable, " + envVar + " not found")
		}
	}

	if strings.HasPrefix(c.Settings.LocalesDir, "env:") {
		envVar := strings.TrimPrefix(c.Settings.LocalesDir, "env:")
		localesDir, ok := os.LookupEnv(envVar)
		c.Settings.LocalesDir = localesDir
		if !ok {
			panic("telebot-config: failed to get localesDir from environment, variable, " + envVar + " not found")
		}
	}

	if len(c.Settings.AllowedUpdates) == 0 {
		log.Warnf("telebot-config: no allowed updates specified, telegram will chose the default for you.")
	}

	for _, update := range c.Settings.AllowedUpdates {
		if !update.IsValid() {
			panic("telebot-config: invalid update type, " + update.String())
		}
	}
}

func loadLocales(path string, c *config, locales ...string) {
	possibleTransSuffix := []string{"yaml", "yml", "json"}

	for _, l := range locales {
		for _, s := range possibleTransSuffix {
			f := path + "/" + l + "." + s
			if _, err := os.Stat(f); err == nil {
				localeRaw, err := os.ReadFile(f)
				if err != nil {
					panic("telebot-config: failed to read locale file, " + f + ", " + err.Error())
				}

				locale := make(langMap)
				err = yaml.Unmarshal(localeRaw, &locale)
				if err != nil {
					panic("telebot-config: failed to parse locale file, " + f + ", " + err.Error())
				}

				c.locales[localeKey(l)] = locale
			}
		}
	}
}

func (c *config) Keyboard(text string, l ...string) tele.ReplyMarkup {
	return c.keyboards.GetKeyboard(text, l...)
}
