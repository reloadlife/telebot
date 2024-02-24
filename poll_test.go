package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_Online_SendPoll(t *testing.T) {
	t.Run("Send a Poll Message With all Possible Tags", func(t *testing.T) {
		markup := NewMarkup()
		markup.Inline()
		markup.AddRow(NewInlineKeyboardButton("Markup on poll !!", "hey"))
		msg, err := tg.SendPoll(whereTo, "do you like TeleBot ?!",
			[]string{"yes!", "very much!", "YES but in caps!"},
			IsClosedPoll,
			IsAnonymousPoll,
			AllowMultipleAnswers,
			0,
			PollTypeQuiz,
			"**We all love TeleBot**",
			time.Second*50,
			ParseModeMarkdownV2,
			markup,
			Silent,
			Protected,
		)
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, intChatID)
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			_, _ = tg.SendPoll(whereTo, "love?", []string{"Some Caption!"}, false, true)
		})
	})
}
