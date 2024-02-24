package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Online_SendChatAction(t *testing.T) {
	t.Run("Send a ChatAction With all Possible Tags", func(t *testing.T) {
		err := tg.SendChatAction(whereTo, ChatTyping)
		require.NoError(t, err)
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			_ = tg.SendChatAction(whereTo, ChatTyping, false, true)
		})
	})
}
