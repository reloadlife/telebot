package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMethods(t *testing.T) {
	tg := New(BotSettings{
		OfflineMode: true,
		Token:       "12345678:ABCDEFG",
	})

	_ = t.Run("LocalScope/OfflineMode/TelegramAPI/Close", func(t *testing.T) {
		err := tg.Close()
		require.NoError(t, err)
	})

	_ = t.Run("LocalScope/OfflineMode/TelegramAPI/Logout", func(t *testing.T) {
		err := tg.Logout()
		require.NoError(t, err)
	})
}

func TestOnlineMethods(t *testing.T) {
	//tg := GetBot()

	_ = t.Run("LocalScope/Real/TelegramAPI/Close", func(t *testing.T) {
		//err := tg.Close()
		//require.NoError(t, err)
		t.Logf("Close method doesn't require test, cuz if you call it then for 15 minutes, no other method is callable via that API.")
	})

	_ = t.Run("LocalScope/Real/TelegramAPI/Logout", func(t *testing.T) {
		//err := tg.Logout()
		//require.NoError(t, err)
		t.Logf("Logout method doesn't require test, cuz if you call it then for 15 minutes, no other method is callable via that API.")
	})
}
