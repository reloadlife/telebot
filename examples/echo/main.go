package echo

import (
	"go.mamad.dev/telebot"
	"os"
)

func main() {
	tg, err := telebot.NewBot(telebot.Settings{
		Token: os.Getenv("TELEGRAM_TOKEN"),
	})

	if err != nil {
		panic(err)
	}

	tg.Handle("/start", func(ctx telebot.Context) error {
		return ctx.Send("Hello Sir, I'm Echo Bot, Please send me something to echo.")
	})

	tg.Handle(telebot.OnText, func(ctx telebot.Context) error {
		return ctx.Send(ctx.Text())
	})

	tg.Start()
}
