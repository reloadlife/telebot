package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

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
