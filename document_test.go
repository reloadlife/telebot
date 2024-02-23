package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Online_SendDocument(t *testing.T) {
	DocumentFileFromURL := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/pdffile.pdf")
	DocumentFileFromFile := FromDisk("./.github/pdffile.pdf")
	DocumentFileFromFile.fileName = "pdffile.pdf"

	require.True(t, DocumentFileFromFile.OnDisk())
	t.Run("Upload a file FROM URL ", func(t *testing.T) {
		markup := NewMarkup()
		markup.Inline()
		markup.AddRow(NewInlineKeyboardButton("test!", "hey"))
		thumb := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/thumb.jpeg")

		msg, err := tg.SendDocument(whereTo, DocumentFileFromURL,
			"**Caption\\.** with Thumbnail",
			toPtr("**Caption\\.** with Thumbnail"),
			toPtr(ParseModeMarkdownV2),
			markup,
			Silent,
			Protected,
			&thumb,
		)
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, intChatID)
		require.Equal(t, *msg.Caption, "Caption. with Thumbnail")
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			_, _ = tg.SendDocument(whereTo, DocumentFileFromURL, "Some Caption!", false, true)
		})
	})

}
