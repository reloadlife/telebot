package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_Online_SendMediaGroup(t *testing.T) {
	PhotoFromFromURL := FromURL("https://raw.githubusercontent.com/reloadlife/telebot/master/.github/telegramlogo.png")

	MediaList := []InputMedia{
		{
			MediaType: InputMediaPhoto,
			Media:     &PhotoFromFromURL,
			Caption:   toPtr("Caption"),
		},
		{
			MediaType: InputMediaPhoto,
			Media:     &PhotoFromFromURL,
			Caption:   toPtr("Caption"),
		},
		{
			MediaType: InputMediaPhoto,
			Media:     &PhotoFromFromURL,
			Caption:   toPtr("Caption"),
		},
	}

	msgs, err := tg.SendMediaGroup(whereTo, MediaList,
		Silent,
		Protected,
		replyParam,
	)
	require.NoError(t, err)
	require.NotNil(t, msgs)
	require.Len(t, msgs, 3)
	require.Equal(t, *msgs[0].Caption, "Caption")

	require.Panics(t, func() {
		_, _ = tg.SendMediaGroup(whereTo, []InputMedia{
			{
				MediaType: InputMediaPhoto,
				Media:     &PhotoFromFromURL,
				Caption:   toPtr("Caption"),
			}})
	})

}
