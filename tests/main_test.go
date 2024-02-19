package tests

import (
	"github.com/stretchr/testify/require"
	tele "go.mamad.dev/telebot"
	"testing"
	"time"
)

func TestMethods(t *testing.T) {
	tg := tele.New(tele.BotSettings{
		OfflineMode: true,
		Token:       "12345678:ABCDEFG",
	})

	_ = t.Run("OfflineMode/TelegramAPI/Close", func(t *testing.T) {
		err := tg.Close()
		require.NoError(t, err)
	})

	_ = t.Run("OfflineMode/TelegramAPI/Logout", func(t *testing.T) {
		err := tg.Logout()
		require.NoError(t, err)
	})

	_ = t.Run("OfflineMode/TelegramAPI/GetUpdates", func(t *testing.T) {
		updates, err := tg.GetUpdates(0, 0, time.Second)
		require.NoError(t, err)
		require.Empty(t, updates) // no updates in offline mode
	})

	_ = t.Run("OfflineMode/TelegramAPI/GetMe", func(t *testing.T) {
		user, err := tg.GetMe()
		require.NoError(t, err)
		require.Equal(t, "TestBot", *user.Username)
		require.True(t, user.IsBot)
	})

	_ = t.Run("OfflineMode/TelegramAPI/SendMessage", func(t *testing.T) {
		_, err := tg.SendMessage(&tele.Chat{
			ID: 123,
		}, "Hello, world!")
		require.NoError(t, err)
	})

}
