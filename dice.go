package telebot

import "fmt"

// Dice represents an animated emoji that displays a random value.
type Dice struct {
	// Emoji is the emoji on which the dice throw animation is based.
	Emoji StickerEmoji `json:"emoji"`

	// Value is the value of the dice.
	// For “🎲”, “🎯”, and “🎳” base emoji, the value ranges from 1 to 6.
	// For “🏀” and “⚽” base emoji, the value ranges from 1 to 5.
	// For “🎰” base emoji, the value ranges from 1 to 64.
	Value int `json:"value"`
}

func (d *Dice) ReflectType() string { return fmt.Sprintf("%T", d) }
func (d *Dice) Type() string        { return "Dice" }
