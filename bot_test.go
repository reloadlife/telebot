package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBot_GetMe(t *testing.T) {
	b, err := tg.GetMe()
	require.NoError(t, err)

	require.True(t, b.IsBot)
	require.NotEmpty(t, b.Username)

	require.Equal(t, *b.Username, "TeleBotUnitTestBot")
}
