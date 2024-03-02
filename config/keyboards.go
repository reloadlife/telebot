package config

import (
	tele "go.mamad.dev/telebot"
	"strings"
)

type keyboardMap struct {
	Type string `yaml:"type" json:"type"`

	Placeholder string `yaml:"placeholder" json:"placeholder"`
	OneTime     bool   `yaml:"one_time" json:"one_time"`
	Resize      bool   `yaml:"resize" json:"resize"`
	Selective   bool   `yaml:"selective" json:"selective"`
	Persistent  bool   `yaml:"persistent" json:"persistent"`

	Buttons [][]string `yaml:"keyboard" json:"keyboard"`
}

type keyboards struct {
	conf *config

	keyboards map[string]map[string]tele.ReplyMarkup
}

type Keyboards interface {
	GetKeyboard(key string, l ...string) tele.ReplyMarkup
}

func (k *keyboards) GetKeyboard(key string, l ...string) tele.ReplyMarkup {
	locale := k.conf.GetDefaultLocale()
	if len(l) > 0 {
		locale = l[0]
	}
	return k.keyboards[locale][key]
}

func (c *config) GetKeyboards() Keyboards {
	k := keyboards{
		conf:      c,
		keyboards: make(map[string]map[string]tele.ReplyMarkup),
	}

	keys := c.conf.Keyboards
	gb := c.GetButtons()

	for _, _l := range c.GetLocales() {
		if _, ok := k.keyboards[_l]; !ok {
			k.keyboards[_l] = make(map[string]tele.ReplyMarkup)
		}

		for i, v := range keys {
			markup := tele.NewMarkup()
			switch v.Type {
			case "inline":
				markup.Inline()
				for _, row := range v.Buttons {
					butts := make([]tele.Button, 0)
					for _, btn := range row {
						butts = append(butts, gb.GetButton(btn, _l))
					}
					markup.AddRow(butts...)
				}
			case "keyboard":
				var opts []any
				if v.OneTime {
					opts = append(opts, tele.OneTimeKeyboard)
				}

				if v.Resize {
					opts = append(opts, tele.ResizeKeyboard)
				}

				if v.Persistent {
					opts = append(opts, tele.PersistentKeyboard)
				}

				if v.Selective {
					opts = append(opts, tele.Selective)
				}

				if v.Placeholder != "" {
					plHolder := v.Placeholder
					if strings.HasPrefix(plHolder, "locale:") {
						pl := strings.TrimPrefix(plHolder, "locale:")
						plHolder = c.L(_l, pl)
					}
					opts = append(opts, plHolder)
				}

				markup.Keyboard(opts...)
				for _, row := range v.Buttons {
					butts := make([]tele.Button, 0)
					for _, btn := range row {
						butts = append(butts, gb.GetButton(btn, _l))
					}
					markup.AddRow(butts...)
				}
			}
			k.keyboards[_l][i] = markup
		}
	}
	return &k
}
