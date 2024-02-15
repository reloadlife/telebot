package tests

import (
	"github.com/stretchr/testify/require"
	tele "go.mamad.dev/telebot"
	"testing"
)

func TestMethods(t *testing.T) {
	tg := tele.New(tele.BotSettings{
		OfflineMode: true,
	})

	_, err := tg.GetMe()
	require.NoError(t, err)
}
