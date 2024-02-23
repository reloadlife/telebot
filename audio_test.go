package telebot

import (
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"
	"time"
)

func Test_Online_SendAudio(t *testing.T) {
	tg := GetBot()

	chatID, hasChatID := os.LookupEnv("CHAT_ID")

	if !hasChatID {
		panic("CHAT_ID is not set")
	}

	chat, err := strconv.ParseInt(chatID, 10, 64)
	require.NoError(t, err)
	whereTo := &Chat{ID: chat}

	AudioFromURL := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/audio.m4a")
	AudioFromFile := FromDisk("./.github/audio.m4a")
	AudioFromFile.fileName = "audio.m4a"

	require.True(t, AudioFromFile.OnDisk())

	var msg *AccessibleMessage

	t.Run("SendFromURL", func(t *testing.T) {
		msg, err = tg.SendAudio(whereTo, AudioFromURL, "Sample Caption.")
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, chat)
		require.Equal(t, *msg.Caption, "Sample Caption.")
	})

	t.Run("SendFromDiscFile", func(t *testing.T) {
		msg, err = tg.SendAudio(whereTo, AudioFromFile, "Sample Caption. this one is from Disc.")
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, chat)
		require.Equal(t, *msg.Caption, "Sample Caption. this one is from Disc.")
	})

	fileID := FromFileID(msg.Audio.FileID)

	t.Run("ResendFromFileID", func(t *testing.T) {
		msg, err = tg.SendAudio(whereTo, fileID, "Sample Caption. this one is from FileID.")
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, chat)
		require.Equal(t, *msg.Caption, "Sample Caption. this one is from FileID.")
	})

	t.Run("WithAllTags", func(t *testing.T) {
		markup := NewMarkup()
		markup.Inline()
		markup.AddRow(NewInlineKeyboardButton("test!", "hey"))
		thumb := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/thumb.jpeg")

		msg, err = tg.SendAudio(whereTo, fileID,
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

		require.Equal(t, msg.Chat.ID, chat)
		require.Equal(t, *msg.Caption, "Caption.")
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			msg, err = tg.SendAudio(whereTo, AudioFromURL, "Some Caption!", false, true)
		})
	})

}
