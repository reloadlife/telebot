package telebot

func GetBot() Bot {
	return New(BotSettings{
		Token: readFromEnv("TELEBOT_SECRET"),
	})
}
