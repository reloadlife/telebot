package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Online_SendLocation(t *testing.T) {
	loc := Location{
		Longitude:            85,
		Latitude:             69,
		HorizontalAccuracy:   42,
		LivePeriod:           185,
		Heading:              0,
		ProximityAlertRadius: 100,
	}
	t.Run("Send a Location Message With all Possible Tags", func(t *testing.T) {
		markup := NewMarkup()
		markup.Inline()
		markup.AddRow(NewInlineKeyboardButton("Location !!", "hey"))
		msg, err := tg.SendLocation(whereTo, loc,
			markup,
			Silent,
			Protected,
		)
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, intChatID)
		require.Equal(t, int(msg.Location.Longitude), int(loc.Longitude))
		require.Equal(t, int(msg.Location.Latitude), int(loc.Latitude))
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			_, _ = tg.SendLocation(whereTo, loc, "Some Caption!", false, true)
		})
	})
}
