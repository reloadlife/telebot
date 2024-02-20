package telebot

import (
	"github.com/stretchr/testify/require"
	"os"
	"strconv"
	"testing"
	"time"
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
		miMsg := &MaybeInaccessibleMessage{
			InaccessibleMessage: nil,
			AccessibleMessage:   nil,
		}

		require.NotNil(t, miMsg)

		require.NotPanics(t, func() {
			require.Equal(t, miMsg.MessageType(), MaybeInaccessibleMessageType)
		})

		require.False(t, miMsg.IsAccessible())

		t.Run("/Verify", func(t *testing.T) {
			require.Error(t, miMsg.Verify())
		})

		t.Run("AccessibleMessage/Fill", func(t *testing.T) {
			miMsg.AccessibleMessage = &AccessibleMessage{
				ID:   12,
				Date: time.Now().Unix(),
			}
			require.NotNil(t, miMsg.AccessibleMessage)

			require.NotPanics(t, func() {
				require.Equal(t, miMsg.MessageType(), MaybeInaccessibleMessageType)
			})

			require.NotPanics(t, func() {
				require.True(t, miMsg.IsAccessible())
			})
		})

		t.Run("MaybeInaccessibleMessage/MarshalJSON", func(t *testing.T) {
			miMsg = &MaybeInaccessibleMessage{}
			miMsg.AccessibleMessage = &AccessibleMessage{
				ID:   12,
				Date: time.Now().Unix(),
				From: &User{
					ID: 1234,
				},
			}
			miMsg.InaccessibleMessage = nil
			t.Logf("isAccessible: %v", miMsg.IsAccessible())
			t.Logf("miMsg: %v", miMsg)
			in, err := miMsg.MarshalJSON()
			require.NoError(t, err)
			t.Logf("in: %s", in)
		})

		t.Run("MaybeInaccessibleMessage/UnmarshalJSON", func(t *testing.T) {
			miMsg = &MaybeInaccessibleMessage{}

			in := []byte(`{"id":12,"date":1620000000,"from":{"id":1234}}`)
			err := miMsg.UnmarshalJSON(in)
			require.NoError(t, err)

			t.Logf("IsAccessible: %v", miMsg.IsAccessible())

			require.True(t, miMsg.IsAccessible())
			t.Logf("miMsg: %v", miMsg)
		})
	})

	aMsg := &AccessibleMessage{}
	require.NotNil(t, aMsg)

	iMsg := &InaccessibleMessage{}
	require.NotNil(t, iMsg)
}
