package telebot

import "fmt"

type MenuButton struct {
	Menu   MenuButtonType `json:"type"`
	Text   *string        `json:"text,omitempty"`
	WebApp *WebAppInfo    `json:"web_app,omitempty"`
}

func (c *MenuButton) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *MenuButton) Type() string {
	if c.Menu == "" {
		return Unknown
	}

	return string(c.Menu)
}
