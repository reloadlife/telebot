package telebot

type MenuButton struct {
	Menu   MenuButtonType `json:"type"`
	Text   *string        `json:"text,omitempty"`
	WebApp *WebAppInfo    `json:"web_app,omitempty"`
}
