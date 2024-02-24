package telebot

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func Test_Online_SendMessage(t *testing.T) {
	text := "*MarkDownMessage* [link](https://github.com/reloadlife/telebot) `code`"
	markup := NewMarkup()
	markup.Inline()
	markup.AddRow(NewInlineKeyboardButton("Hey", "hey"))
	markup.AddRow(NewInlineKeyboardButton("TeleBot !!", "hey"))

	msg, err = tg.SendMessage(whereTo, text,
		ParseModeMarkdown,
		Protected,
		&LinkPreviewOptions{
			URL:           toPtr("https://go.mamad.dev/telebot"),
			ShowAboveText: toPtr(true),
		})
	require.NoError(t, err)
	require.NotNil(t, msg)

	require.Equal(t, msg.Entities, []Entity{
		{
			EntityType: EntityTypeBold,
			Offset:     0,
			Length:     15,
		},
		{
			EntityType: EntityTypeTextLink,
			Offset:     16,
			Length:     4,
			URL:        "https://github.com/reloadlife/telebot",
		},
		{
			EntityType: EntityTypeCode,
			Offset:     21,
			Length:     4,
		},
	})

	replyParam.MessageID = msg.ID

	t.Run("WithBadOptions", func(t *testing.T) {
		text = fmt.Sprintf("Hello, World, With reply markup and Reply Message (%d)!", msg.ID)
		require.Panics(t, func() {
			msg, err = tg.SendMessage(whereTo, text, false, true)
		})
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
