package telebot

import (
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"
)

func Test_Online_SendMessage(t *testing.T) {
	tg := GetBot()

	chatID, hasChatID := os.LookupEnv("CHAT_ID")

	if !hasChatID {
		panic("CHAT_ID is not set")
	}

	chat, err := strconv.ParseInt(chatID, 10, 64)

	require.NoError(t, err)

	whereTo := &Chat{ID: chat}

	b, err := tg.SendMessage(whereTo, "Hello, World!")
	require.NoError(t, err)

	require.NotNil(t, b)

	require.Equal(t, b.Chat.ID, chat)
	require.Equal(t, *b.Text, "Hello, World!") // nullable stuff are pointer. So, we need to dereference it.
}

func TestMessage(t *testing.T) {
	t.Run("MaybeInaccessibleMessage", func(t *testing.T) {
		miMsg := &MaybeInaccessibleMessage{}
		require.NotNil(t, miMsg)

		require.NotPanics(t, func() {
			require.Equal(t, miMsg.MessageType(), MaybeInaccessibleMessageType)
		})

		require.Panics(t, func() {
			_ = miMsg.IsAccessible()
		})

		t.Run("InaccessibleMessage/Verify", func(t *testing.T) {
			require.NotPanics(t, func() {
				_ = miMsg.Verify()
			})
		})

		t.Run("AccessibleMessage/Fill", func(t *testing.T) {
			// fill in the MaybeInaccessibleMessage fields
			miMsg.AccessibleMessage = &AccessibleMessage{
				ID: 12,
			}
			require.NotNil(t, miMsg.AccessibleMessage)

			require.NotPanics(t, func() {
				require.Equal(t, miMsg.MessageType(), MaybeInaccessibleMessageType)
			})

			require.NotPanics(t, func() {
				require.True(t, miMsg.IsAccessible())
			})
		})
	})
	aMsg := &AccessibleMessage{}
	require.NotNil(t, aMsg)

	iMsg := &InaccessibleMessage{}
	require.NotNil(t, iMsg)
}
