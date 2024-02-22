package telebot

import (
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"
)

func Test_Online_SendPhoto(t *testing.T) {
	tg := GetBot()

	chatID, hasChatID := os.LookupEnv("CHAT_ID")

	if !hasChatID {
		panic("CHAT_ID is not set")
	}

	chat, err := strconv.ParseInt(chatID, 10, 64)
	require.NoError(t, err)
	whereTo := &Chat{ID: chat}

	pictureFromURL := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/telegramlogo.png")
	pictureFromFile := FromDisk("./.github/telegramlogo.png")
	pictureFromFile.fileName = "telegramlogo.png"

	require.True(t, pictureFromFile.OnDisk())

	var msg *AccessibleMessage

	t.Run("SendFromURL", func(t *testing.T) {
		msg, err = tg.SendPhoto(whereTo, pictureFromURL, "Sample Caption.")
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, chat)
		require.Equal(t, *msg.Caption, "Sample Caption.")
	})

	t.Run("SendFromDiscFile", func(t *testing.T) {
		msg, err = tg.SendPhoto(whereTo, pictureFromFile, "Sample Caption. this one is from Disc.")
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, chat)
		require.Equal(t, *msg.Caption, "Sample Caption. this one is from Disc.")

		t.Run("ResendFromFileID", func(t *testing.T) {
			fromFileID := FromFileID(msg.Photo[0].FileID)
			msg, err = tg.SendPhoto(whereTo, fromFileID, "Sample Caption. this one is from FileID.")
			require.NoError(t, err)
			require.NotNil(t, msg)

			require.Equal(t, msg.Chat.ID, chat)
			require.Equal(t, *msg.Caption, "Sample Caption. this one is from FileID.")

			t.Run("WithSpoiler", func(t *testing.T) {
				msg, err = tg.SendPhoto(whereTo, fromFileID, "Sample Caption. HasSpoiler = 1", HasSpoiler)
				require.NoError(t, err)
				require.NotNil(t, msg)

				require.Equal(t, msg.Chat.ID, chat)
				require.Equal(t, *msg.Caption, "Sample Caption. HasSpoiler = 1")
			})
		})
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			msg, err = tg.SendPhoto(whereTo, pictureFromURL, "Some Caption!", false, true)
		})
	})

}
