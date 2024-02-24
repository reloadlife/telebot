package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_Online_SendAudio(t *testing.T) {
	AudioFromURL := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/audio.m4a")
	AudioFromFile := FromDisk("./.github/audio.m4a")
	AudioFromFile.fileName = "audio.m4a"
	require.True(t, AudioFromFile.OnDisk())

	markup := NewMarkup()
	markup.Inline()
	markup.AddRow(NewInlineKeyboardButton("Audio TEST", "hey"))

	msg, err = tg.SendAudio(whereTo, AudioFromURL,
		"**Caption\\.**",
		ParseModeMarkdownV2,
		markup,
		Silent,
		Protected,
		106*time.Second,
		Performer("TeleBot"),
		Title("Test Audio"),
		&thumb,
		replyParam,
	)
	require.NoError(t, err)
	require.NotNil(t, msg)

	require.Equal(t, msg.Chat.ID, intChatID)
	require.Equal(t, *msg.Caption, "Caption.")

	replyParam.MessageID = msg.ID

	require.Panics(t, func() {
		_, _ = tg.SendAudio(whereTo, AudioFromURL, "Some Caption!", false, true)
	})

}
