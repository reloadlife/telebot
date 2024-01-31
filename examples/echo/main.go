package main

import (
	tele "go.mamad.dev/telebot"
	"os"
)

func main() {
	tg := tele.New(tele.BotSettings{
		Token: os.Getenv("TELEGRAM_TOKEN"),
	})

	tg.Handle("/start", func(ctx tele.Context) error {
		return ctx.Send("Hello Sir, I'm Echo OldBot, Please send me something to echo.")
	})

	tg.Handle(tele.OnText, func(ctx tele.Context) error {
		return ctx.Send(ctx.Text())
	})

	tg.Start()
}
