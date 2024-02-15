package tests

import (
	"errors"
	"github.com/stretchr/testify/require"
	tele "go.mamad.dev/telebot"
	"testing"
)

func TestMethods(t *testing.T) {
	tg := tele.New(tele.BotSettings{
		OfflineMode: true,
		Token:       "12345678:ABCDEFG",
	})

	err := errors.New("this would be a real error")

	err = tg.Close()
	require.NoError(t, err)

	err = tg.Logout()
	require.NoError(t, err)

	_, err = tg.GetMe()
	require.NoError(t, err)
}
