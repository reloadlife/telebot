package telebot

import "fmt"

type TextQuote struct {
	Text     string   `json:"text"`
	Entities []Entity `json:"entities,omitempty"`
	Position int      `json:"position"`
	IsManual *bool    `json:"is_manual,omitempty"`
}

func (c *TextQuote) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *TextQuote) Type() string        { return "TextQuote" }
