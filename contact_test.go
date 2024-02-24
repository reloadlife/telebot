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
	markup := NewMarkup()
	markup.Inline()
	markup.AddRow(NewInlineKeyboardButton("Contact Test", "hey"))
	msg, err = tg.SendContact(whereTo, contact, markup, Silent, Protected, replyParam)

	require.NoError(t, err)
	require.NotNil(t, msg)

	require.Equal(t, msg.Chat.ID, intChatID)
	require.Equal(t, msg.Contact.PhoneNumber, contact.PhoneNumber)

	replyParam.MessageID = msg.ID

	require.Panics(t, func() {
		_, _ = tg.SendContact(whereTo, contact, "Some Caption!", false, true)
	})

}
