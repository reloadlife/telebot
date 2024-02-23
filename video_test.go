package telebot

import (
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"
	"time"
)

func Test_Online_SendVideo(t *testing.T) {
	tg := GetBot()

	chatID, hasChatID := os.LookupEnv("CHAT_ID")

	if !hasChatID {
		panic("CHAT_ID is not set")
	}

	chat, err := strconv.ParseInt(chatID, 10, 64)
	require.NoError(t, err)
	whereTo := &Chat{ID: chat}

	VideoFileFromURL := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/video.mp4")
	VideoFileFromFile := FromDisk("./.github/video.mp4")
	VideoFileFromFile.fileName = "video.mp4"

	require.True(t, VideoFileFromFile.OnDisk())

	var msg *AccessibleMessage

	t.Run("SendFromURL", func(t *testing.T) {
		msg, err = tg.SendVideo(whereTo, VideoFileFromURL, "Sample Caption.")
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, chat)
		require.Equal(t, *msg.Caption, "Sample Caption.")
	})

	t.Run("SendFromDiscFile", func(t *testing.T) {
		msg, err = tg.SendVideo(whereTo, VideoFileFromFile, "Sample Caption. this one is from Disc.")
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, chat)
		require.Equal(t, *msg.Caption, "Sample Caption. this one is from Disc.")
	})

	fileID := FromFileID(msg.Video.FileID)

	t.Run("ResendFromFileID", func(t *testing.T) {
		msg, err = tg.SendVideo(whereTo, fileID, "Sample Caption. this one is from FileID.", DisableContentTypeDetection)
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

		msg, err = tg.SendVideo(whereTo, fileID,
			"**Caption\\.** with Thumbnail",
			toPtr("**Caption\\.** with Thumbnail"),
			toPtr(ParseModeMarkdownV2),
			markup,
			Silent,
			Protected,
			65*time.Second,
			Width(1080),
			Height(1080),
			Streamable,
			HasSpoiler,
			&thumb,
		)
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, chat)
		require.Equal(t, *msg.Caption, "Caption. with Thumbnail")
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			msg, err = tg.SendVideo(whereTo, VideoFileFromURL, "Some Caption!", false, true)
		})
	})

}
