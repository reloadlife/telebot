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
	markup := NewMarkup()
	markup.Inline()
	markup.AddRow(NewInlineKeyboardButton("Document test", "hey"))

	msg, err = tg.SendDocument(whereTo, DocumentFileFromURL,
		"__Document__ `File Test`",
		ParseModeMarkdownV2,
		markup,
		Silent,
		Protected,
		&thumb,
	)
	require.NoError(t, err)
	require.NotNil(t, msg)

	require.Equal(t, msg.Chat.ID, intChatID)
	require.Equal(t, *msg.Caption, "Document File Test")
	require.Panics(t, func() {
		_, _ = tg.SendDocument(whereTo, DocumentFileFromURL, "Some Caption!", false, true)
	})
}
