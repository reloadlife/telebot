package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_Online_SendVoice(t *testing.T) {
	VoiceFromURL := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/sample.ogg")
	VoiceFromFile := FromDisk("./.github/sample.ogg")
	VoiceFromFile.fileName = "sample.ogg"
	require.True(t, VoiceFromFile.OnDisk())

	t.Run("Upload an Voice file using a URL with all possible Tags", func(t *testing.T) {
		markup := NewMarkup()
		markup.Inline()
		markup.AddRow(NewInlineKeyboardButton("test!", "hey"))

		msg, err := tg.SendVoice(whereTo, VoiceFromURL,
			"**Caption\\.**",
			toPtr("**Caption\\.**"),
			toPtr(ParseModeMarkdownV2),
			markup,
			Silent,
			Protected,
			1*time.Second,
		)
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, intChatID)
		require.Equal(t, *msg.Caption, "Caption.")
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			_, _ = tg.SendVoice(whereTo, VoiceFromURL, "Some Caption!", false, true)
		})
	})

}
