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
	markup := NewMarkup()
	markup.Inline()
	markup.AddRow(NewInlineKeyboardButton("test!", "hey"))

	msg, err = tg.SendPhoto(whereTo, pictureFromFile,
		"**Sample Caption\\.**",
		HasSpoiler,
		ParseModeMarkdownV2,
		Silent,
		Protected,
		replyParam,
		markup)

	require.NoError(t, err)
	require.NotNil(t, msg)

	replyParam.MessageID = msg.ID
	fileID = msg.Photo.HighRes().FileID

	require.Panics(t, func() {
		_, _ = tg.SendPhoto(whereTo, pictureFromURL, "Some Caption!", false, true)
	})
}
