package config

import (
	tele "go.mamad.dev/telebot"
	"net/url"
	"strings"
)

type ButtonMap struct {
	Type string `yaml:"type" json:"type"`
	Text string `yaml:"text" json:"text"`

	RequestContact  bool `yaml:"request_contact" json:"request_contact"`
	RequestLocation bool `yaml:"request_location" json:"request_location"`

	RequestUsers *tele.KeyboardButtonRequestUsers `yaml:"request_users" json:"request_users"`
	RequestChat  *tele.KeyboardButtonRequestChat  `yaml:"request_chat" json:"request_chat"`

	PollRequest *tele.PollType `yaml:"poll_request" json:"poll_request"`
	WebApp      string         `yaml:"web_app" json:"web_app"`

	URL                          string                            `yaml:"url" json:"url"`
	Data                         string                            `yaml:"callback_data" json:"callback_data"`
	Pay                          bool                              `yaml:"pay" json:"pay"`
	LoginURL                     *tele.LoginURL                    `yaml:"login_url" json:"login_url"`
	Game                         bool                              `yaml:"game" json:"game"`
	SwitchInlineQuery            string                            `yaml:"switch_inline_query" json:"switch_inline_query"`
	SwitchInlineQueryCurrentChat string                            `yaml:"switch_inline_query_current_chat" json:"switch_inline_query_current_chat"`
	SwitchInlineQueryChosenChat  *tele.SwitchInlineQueryChosenChat `yaml:"switch_inline_query_chosen_chat" json:"switch_inline_query_chosen_chat"`
}

type buttons struct {
	conf *config

	buttons map[string]tele.Button
}

type Buttons interface {
	GetButton(string, ...string) tele.Button
}

func (c *config) GetButtons() Buttons {
	btns := c.conf.Buttons

	but := buttons{
		conf:    c,
		buttons: make(map[string]tele.Button),
	}

	for k, v := range btns {
		t := v.Type
		if t == "" {
			t = "keyboardButton"
		}

		text := v.Text

		if text == "" {
			panic("telebot-config: text is required for button")
		}

		extraOptions := make([]any, 0)

		switch t {
		case "keyboardButton":
			if v.RequestContact {
				extraOptions = append(extraOptions, tele.RequestContact(true))
			}

			if v.RequestLocation {
				extraOptions = append(extraOptions, tele.RequestLocation(true))
			}

			if v.RequestChat != nil {
				extraOptions = append(extraOptions, v.RequestChat)
			}

			if v.RequestUsers != nil {
				extraOptions = append(extraOptions, v.RequestUsers)
			}

			if v.PollRequest != nil {
				extraOptions = append(extraOptions, v.PollRequest)
			}

			if v.WebApp != "" {
				extraOptions = append(extraOptions, tele.WebAppInfo{
					URL: v.WebApp,
				})
			}

			but.buttons[k] = tele.NewKeyboardButton(text, extraOptions...)
		case "inlineKeyboardButton":

			if v.URL != "" {
				u, err := url.Parse(v.URL)
				if err != nil {
					panic("telebot-config: invalid URL provided for inlineKeyboardButton, " + err.Error())
				}
				extraOptions = append(extraOptions, u)
			}

			if v.Data != "" {
				extraOptions = append(extraOptions, v.Data)
			}

			if v.Pay {
				extraOptions = append(extraOptions, true)
			}

			if v.LoginURL != nil {
				extraOptions = append(extraOptions, v.LoginURL)
			}

			if v.Game {
				extraOptions = append(extraOptions, &tele.CallbackGame{})
			}

			if v.SwitchInlineQuery != "" {
				i := tele.SwitchInlineQueryStringType(v.SwitchInlineQuery)
				extraOptions = append(extraOptions, &i)
			}

			if v.SwitchInlineQueryCurrentChat != "" {
				i := tele.SwitchInlineQueryChosenChatStringType(v.SwitchInlineQueryCurrentChat)
				extraOptions = append(extraOptions, &i)
			}

			if v.SwitchInlineQueryChosenChat != nil {
				extraOptions = append(extraOptions, v.SwitchInlineQueryChosenChat)
			}

			if v.WebApp != "" {
				extraOptions = append(extraOptions, tele.WebAppInfo{
					URL: v.WebApp,
				})
			}

			but.buttons[k] = tele.NewInlineKeyboardButton(
				text,
				extraOptions...,
			)
		}
	}
	return &but
}

func (b *buttons) GetButton(name string, l ...string) tele.Button {
	locale := b.conf.GetDefaultLocale()
	if len(l) > 0 {
		locale = l[0]
	}

	btn, ok := b.buttons[name]
	if !ok {
		panic("telebot-config: button not found")
	}
	text := btn.GetText()
	but := btn.Clone()
	if strings.HasPrefix(text, "locale:") {
		lKey := strings.TrimPrefix(text, "locale:")
		but.SetText(b.conf.L(locale, lKey))
	}
	return but
}
