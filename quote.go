package telebot

type TextQuote struct {
	Text     string   `json:"text"`
	Entities []Entity `json:"entities,omitempty"`
	Position int      `json:"position"`
	IsManual *bool    `json:"is_manual,omitempty"`
}
