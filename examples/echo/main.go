package main

import (
	tele "go.mamad.dev/telebot"
	"os"
)

func main() {
	tg := tele.New(tele.BotSettings{
		Token: os.Getenv("TELEGRAM_TOKEN"),
	})

	tg.Handle("/start", func(c tele.Context) error {
		// return c.Send("Hello Sir, I'm Echo OldBot, Please send me something to echo.")

		_, err := tg.SendMessage(c.Message().Chat, "Hello Sir, I'm Echo OldBot, Please send me something to echo.")
		return err
	})

	tg.Handle(tele.OnText, func(c tele.Context) error {
		// return c.Send(c.Text())

		_, err := tg.SendMessage(c.Message().Chat, *c.Message().Text)
		return err
	})

	tg.Start()
}
