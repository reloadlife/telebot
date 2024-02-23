package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_Online_SendVideo(t *testing.T) {
	VideoFileFromURL := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/video.mp4")
	VideoFileFromFile := FromDisk("./.github/video.mp4")
	VideoFileFromFile.fileName = "video.mp4"
	require.True(t, VideoFileFromFile.OnDisk())
	t.Run("WithAllTags", func(t *testing.T) {
		markup := NewMarkup()
		markup.Inline()
		markup.AddRow(NewInlineKeyboardButton("test!", "hey"))
		thumb := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/thumb.jpeg")
		msg, err := tg.SendVideo(whereTo, VideoFileFromURL,
			"**Caption\\.** with Thumbnail and MarkDownV2",
			toPtr("**Caption\\.** with Thumbnail and MarkDownV2 as a pointer"),
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

		require.Equal(t, msg.Chat.ID, intChatID)
		require.Equal(t, *msg.Caption, "Caption. with Thumbnail and MarkDownV2 as a pointer")
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			_, _ = tg.SendVideo(whereTo, VideoFileFromURL, "Some Caption!", false, true)
		})
	})

}
