package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Online_SendContact(t *testing.T) {
	contact := Contact{
		PhoneNumber: "+100000001234",
		FirstName:   "TeleBot !",
		LastName:    nil,
		VCard:       toPtr("vCARD !"),
	}
	t.Run("Send a Contact Message With all Possible Tags", func(t *testing.T) {
		markup := NewMarkup()
		markup.Inline()
		markup.AddRow(NewInlineKeyboardButton("Contact !!", "hey"))
		msg, err := tg.SendContact(whereTo, contact,
			markup,
			Silent,
			Protected,
		)
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, intChatID)
		require.Equal(t, msg.Contact.PhoneNumber, contact.PhoneNumber)
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			_, _ = tg.SendContact(whereTo, contact, "Some Caption!", false, true)
		})
	})
}
