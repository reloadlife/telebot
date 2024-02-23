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

	t.Run("SendMediaGroup URL with all possible Tags", func(t *testing.T) {
		msg, err := tg.SendMediaGroup(whereTo, MediaList,
			Silent,
			Protected,
		)
		require.NoError(t, err)
		require.NotNil(t, msg)
	})

	t.Run("WithBadOptions", func(t *testing.T) {
		require.Panics(t, func() {
			_, _ = tg.SendMediaGroup(whereTo, []InputMedia{
				{
					MediaType: InputMediaPhoto,
					Media:     &PhotoFromFromURL,
					Caption:   toPtr("Caption"),
				}})
		})
	})

}
