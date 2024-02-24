package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Online_SendChatAction(t *testing.T) {
	err = tg.SendChatAction(whereTo, ChatTyping)
	require.NoError(t, err)

	require.Panics(t, func() {
		_ = tg.SendChatAction(whereTo, ChatTyping, false, true)
	})
}
