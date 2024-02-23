package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_Online_SendAnimation(t *testing.T) {
	AnimationFileFromURL := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/cat.mp4")
	AnimationFileFromFile := FromDisk("./.github/video.mp4")
	AnimationFileFromFile.fileName = "video.mp4"
	require.True(t, AnimationFileFromFile.OnDisk())
	t.Run("WithAllTags", func(t *testing.T) {
		markup := NewMarkup()
		markup.Inline()
		markup.AddRow(NewInlineKeyboardButton("Caaaatttt !!! ca🐈🐈!", "hey"))
		thumb := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/thumb.jpeg")
		msg, err := tg.SendAnimation(whereTo, AnimationFileFromURL,
			"**Gif Caption\\.** with Thumbnail and MarkDownV2",
			toPtr("**Gif Caption\\.** with Thumbnail and MarkDownV2 as a pointer"),
			toPtr(ParseModeMarkdownV2),
			markup,
			Silent,
			Protected,
			7*time.Second,
			Width(848),
			Height(848),
			HasSpoiler,
			&thumb,
		)
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, intChatID)
		require.Equal(t, *msg.Caption, "Gif Caption. with Thumbnail and MarkDownV2 as a pointer")
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			_, _ = tg.SendAnimation(whereTo, AnimationFileFromURL, "Some Caption!", false, true)
		})
	})

}
