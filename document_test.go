package telebot

import (
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"
)

func Test_Online_SendDocument(t *testing.T) {
	tg := GetBot()

	chatID, hasChatID := os.LookupEnv("CHAT_ID")

	if !hasChatID {
		panic("CHAT_ID is not set")
	}

	chat, err := strconv.ParseInt(chatID, 10, 64)
	require.NoError(t, err)
	whereTo := &Chat{ID: chat}

	DocumentFileFromURL := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/sample-file.txt")
	DocumentFileFromFile := FromDisk("./.github/sample-file.txt")
	DocumentFileFromFile.fileName = "sample-file.txt"

	require.True(t, DocumentFileFromFile.OnDisk())

	var msg *AccessibleMessage

	t.Run("SendFromURL", func(t *testing.T) {
		msg, err = tg.SendDocument(whereTo, DocumentFileFromURL, "Sample Caption.")
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, chat)
		require.Equal(t, *msg.Caption, "Sample Caption.")
	})

	t.Run("SendFromDiscFile", func(t *testing.T) {
		msg, err = tg.SendDocument(whereTo, DocumentFileFromFile, "Sample Caption. this one is from Disc.")
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, chat)
		require.Equal(t, *msg.Caption, "Sample Caption. this one is from Disc.")
	})

	fileID := FromFileID(msg.Audio.FileID)

	t.Run("ResendFromFileID", func(t *testing.T) {
		msg, err = tg.SendDocument(whereTo, fileID, "Sample Caption. this one is from FileID.", DisableContentTypeDetection)
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, chat)
		require.Equal(t, *msg.Caption, "Sample Caption. this one is from FileID.")
	})

	t.Run("WithAllTags", func(t *testing.T) {
		markup := NewMarkup()
		markup.Inline()
		markup.AddRow(NewInlineKeyboardButton("test!", "hey"))
		thumb := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/telegramlogo.png")

		msg, err = tg.SendDocument(whereTo, fileID,
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

		require.Equal(t, msg.Chat.ID, chat)
		require.Equal(t, *msg.Caption, "Caption.")
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			msg, err = tg.SendDocument(whereTo, DocumentFileFromURL, "Some Caption!", false, true)
		})
	})

}
