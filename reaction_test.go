package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Online_SendReaction(t *testing.T) {
	t.Run("Send a message then react to it with Banana", func(t *testing.T) {
		err = tg.SetMessageReaction(whereTo, msg.ID, ReactionType{
			ReactionType: ReactionTypeTypeEmoji,
			Emoji:        EmojiBanana,
		})
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, intChatID)
	})
}
