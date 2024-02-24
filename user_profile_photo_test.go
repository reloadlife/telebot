package telebot

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserProfilePhoto(t *testing.T) {
	profiles, err := tg.GetUserProfilePhotos(400610322, 0, 100)
	require.NoError(t, err)
	require.NotNil(t, profiles)
	photos := profiles.Photos
	t.Logf("Total photos: %d", len(photos))
	for i, photo := range photos {
		t.Logf("Photo %d: %s", i, photo.HighRes().File().URL())
		require.NotNil(t, photo)
		require.NotEmpty(t, photo)
	}
}
