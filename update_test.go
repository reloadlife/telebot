package telebot_test

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"go.mamad.dev/telebot"
	"testing"
)

func TestUpdate_MarshalJSON(t *testing.T) {
	update := &telebot.Update{ID: 123}

	bytes, err := json.Marshal(update)
	require.NoError(t, err)
	require.NotEmpty(t, bytes)
	require.Equal(t, `{"update_id":123}`, string(bytes))
}
