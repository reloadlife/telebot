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

	t.Run("Upload an Audio file using a URL with all possible Tags", func(t *testing.T) {
		markup := NewMarkup()
		markup.Inline()
		markup.AddRow(NewInlineKeyboardButton("test!", "hey"))
		thumb := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/thumb.jpeg")

		msg, err := tg.SendAudio(whereTo, AudioFromURL,
			"**Caption\\.**",
			toPtr("**Caption\\.**"),
			toPtr(ParseModeMarkdownV2),
			markup,
			Silent,
			Protected,
			106*time.Second,
			Performer("TeleBot"),
			Title("Test Audio"),
			&thumb,
		)
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, intChatID)
		require.Equal(t, *msg.Caption, "Caption.")
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			_, _ = tg.SendAudio(whereTo, AudioFromURL, "Some Caption!", false, true)
		})
	})

}
