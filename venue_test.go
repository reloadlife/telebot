package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Online_SendVenue(t *testing.T) {
	ven := Venue{
		Location: Location{
			Longitude:            85,
			Latitude:             69,
			HorizontalAccuracy:   42,
			LivePeriod:           185,
			Heading:              0,
			ProximityAlertRadius: 100,
		},
		Title:           "Lovely Place",
		Address:         "Boyotoofool",
		FoursquareID:    "",
		FoursquareType:  "",
		GooglePlaceID:   "",
		GooglePlaceType: "",
	}
	t.Run("Send a Venue Message With all Possible Tags", func(t *testing.T) {
		markup := NewMarkup()
		markup.Inline()
		markup.AddRow(NewInlineKeyboardButton("Location !!", "hey"))
		msg, err := tg.SendVenue(whereTo, ven,
			markup,
			Silent,
			Protected,
		)
		require.NoError(t, err)
		require.NotNil(t, msg)

		require.Equal(t, msg.Chat.ID, intChatID)
		require.Equal(t, msg.Venue.Title, ven.Title)
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			_, _ = tg.SendVenue(whereTo, ven, "Some Caption!", false, true)
		})
	})
}
