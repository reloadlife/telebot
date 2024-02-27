package telebot

import "go.mamad.dev/gtb/log"

func (b *bot) Debug(logAttrs ...any) {
	for _, l := range logAttrs {
		log.Debugf("%v", l)
	}
}
