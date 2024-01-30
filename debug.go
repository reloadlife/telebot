package telebot

import "fmt"

func (b *bot) Debug(log ...any) {
	for _, l := range log {
		fmt.Println("[telebot] Debug: ", l)
	}
}
