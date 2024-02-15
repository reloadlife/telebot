package telebot

func GetBot() Bot {
	token := readFromEnv("TELEBOT_SECRET")
	if token == "" {
		panic("TELEBOT_SECRET is not set")
	}
	return New(BotSettings{
		Token: token,
	})
}
