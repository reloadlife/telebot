package telebot

import (
	"os"
	"strconv"
)

func GetBot() Bot {
	token := readFromEnv("TELEBOT_SECRET")
	if token == "" {
		panic("TELEBOT_SECRET is not set")
	}
	return New(BotSettings{
		Token: token,
	})
}

var (
	tg        = GetBot() // so we initiate it only once in tests. (for less spam on GetMe)
	whereTo   Recipient  // also we can do this one once ?
	intChatID int64
	msg       *AccessibleMessage
	err       error

	fileID string

	thumb = FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/thumb.jpeg")

	replyParam = &ReplyParameters{
		ChatID:                   intChatID,
		MessageID:                1,
		AllowSendingWithoutReply: toPtr(true),
	}
)

func init() {
	chatID, hasChatID := os.LookupEnv("CHAT_ID")
	if !hasChatID {
		panic("CHAT_ID is not set")
	}
	chat, err := strconv.ParseInt(chatID, 10, 64)
	if err != nil {
		panic(err)
	}
	whereTo = &Chat{ID: chat}
	intChatID = chat
	replyParam.ChatID = intChatID
}
