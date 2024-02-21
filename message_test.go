package telebot

import (
	"fmt"
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

	var msg *AccessibleMessage

	t.Run("Normal", func(t *testing.T) {
		msg, err = tg.SendMessage(whereTo, "Hello, World!")
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, chat)
		require.Equal(t, *msg.Text, "Hello, World!") // nullable stuff are pointer. So, we need to dereference it.
	})

	t.Run("WithReplyToPreviousMessage", func(t *testing.T) {
		text := fmt.Sprintf("Hello, World, With reply to previous message (%d)!", msg.ID)
		msg, err = tg.SendMessage(whereTo, text, &ReplyParameters{
			ChatID:                   whereTo.Recipient(),
			MessageID:                int(msg.ID),
			AllowSendingWithoutReply: toPtr(false),
		})
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, chat)
		require.Equal(t, *msg.Text, text) // nullable stuff are pointer. So,
		// we need to dereference it.
		require.NotNil(t, msg.ReplyTo)
		require.Equal(t, msg.ReplyTo.ID, msg.ID-1)
	})

	t.Run("WithReplyMarkup", func(t *testing.T) {
		text := fmt.Sprintf("Hello, World, With reply markup and Reply Message (%d)!", msg.ID)
		markup := NewMarkup()
		markup.Inline()
		markup.AddRow(NewInlineKeyboardButton("Hey", "hey"))
		markup.AddRow(NewInlineKeyboardButton("Test From GH", "hey"))

		msg, err = tg.SendMessage(whereTo, text, &ReplyParameters{
			ChatID:                   whereTo.Recipient(),
			MessageID:                int(msg.ID),
			AllowSendingWithoutReply: toPtr(false),
		}, markup)
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.True(t, markup.deepEqual(msg.ReplyMarkup))
	})
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

		t.Run("MarshalJSON", func(t *testing.T) {
			in, err := miMsg.MarshalJSON()
			require.ErrorIs(t, err, ErrNoMessageToMarshal)
			require.Nil(t, in)
		})

		t.Run("String", func(t *testing.T) {
			require.NotPanics(t, func() {
				require.NotEmpty(t, miMsg.String())
			})
		})

		t.Run("Verify", func(t *testing.T) {
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

			t.Logf("miMsg: %#v", miMsg)
			t.Logf("msg: %s", miMsg)

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
