package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Online_SendPhoto(t *testing.T) {
	pictureFromURL := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/telegramlogo.png")
	pictureFromFile := FromDisk("./.github/telegramlogo.png")
	pictureFromFile.fileName = "telegramlogo.png"

	require.True(t, pictureFromFile.OnDisk())
	t.Run("Upload a photo from Files on Disc", func(t *testing.T) {
		markup := NewMarkup()
		markup.Inline()
		markup.AddRow(NewInlineKeyboardButton("test!", "hey"))

		_, err := tg.SendPhoto(whereTo, pictureFromFile,
			"**Sample Caption\\.**",
			toPtr("**Sample Caption\\.**"),
			HasSpoiler,
			toPtr(ParseModeMarkdownV2),
			Silent,
			Protected,
			markup)

		require.NoError(t, err)
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			_, _ = tg.SendPhoto(whereTo, pictureFromURL, "Some Caption!", false, true)
		})
	})

}
