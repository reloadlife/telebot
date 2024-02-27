package telebot

import "fmt"

// WebAppData describes data sent from a Web App to the bot.
type WebAppData struct {
	// Data is the data sent from the Web App. Be aware that a bad client can send arbitrary data in this field.
	Data string `json:"data"`

	// ButtonText is the text of the web_app keyboard button from which the Web App was opened.
	// Be aware that a bad client can send arbitrary data in this field.
	ButtonText string `json:"button_text"`
}

func (w *WebAppData) ReflectType() string { return fmt.Sprintf("%T", w) }
func (w *WebAppData) Type() string        { return "WebAppData" }

// WebAppInfo describes a Web App.
type WebAppInfo struct {
	// URL is an HTTPS URL of a Web App to be opened with additional data as specified in Initializing Web Apps.
	URL string `json:"url" yaml:"url"`
}

func (w *WebAppInfo) ReflectType() string { return fmt.Sprintf("%T", w) }
func (w *WebAppInfo) Type() string        { return "WebAppInfo" }

type SentWebAppMessage struct {
	InlineMessageID *string `json:"inline_message_id,omitempty"`
}

func (s *SentWebAppMessage) ReflectType() string { return fmt.Sprintf("%T", s) }
func (s *SentWebAppMessage) Type() string        { return "SentWebAppMessage" }
