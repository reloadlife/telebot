package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Online_SendDice(t *testing.T) {
	t.Run("Send a Dice Message With all Possible Tags", func(t *testing.T) {
		markup := NewMarkup()
		markup.Inline()
		markup.AddRow(NewInlineKeyboardButton("Dice Bro Dice !!", "hey"))
		msg, err := tg.SendDice(whereTo,
			DiceEmojiSlot,
			markup,
			Silent,
			Protected,
		)
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, intChatID)
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			_, _ = tg.SendDice(whereTo, "Cant!", false, true)
		})
	})
}
