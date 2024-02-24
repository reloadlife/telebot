package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_Online_SendVideoNote(t *testing.T) {
	VideoNoteFromURL := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/video_note.mp4")
	VideoNoteFromFile := FromDisk("./.github/video_note.mp4")
	VideoNoteFromFile.fileName = "video_note.mp4"
	require.True(t, VideoNoteFromFile.OnDisk())

	t.Run("Upload an VideoNote file using a URL with all possible Tags", func(t *testing.T) {
		markup := NewMarkup()
		markup.Inline()
		markup.AddRow(NewInlineKeyboardButton("video note test", "hey"))

		msg, err := tg.SendVideoNote(whereTo, VideoNoteFromURL,
			markup,
			Silent,
			Protected,
			2*time.Second,
			Width(400),
		)
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, intChatID)
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			_, _ = tg.SendVideoNote(whereTo, VideoNoteFromURL, "Some Caption!")
		})
	})

}
