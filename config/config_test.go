package config

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var (
	sampleToken = "123456:ABC-DEF1234ghIkl-zyx57W2v1u123ew11"
)

func TestNewConfigFromFile(t *testing.T) {
	var c Config

	require.Panics(t, func() {
		c = NewConfigFromFile("example-config.yaml")
	})

	t.Setenv("TELEGRAM_SECRET", sampleToken)

	require.NotPanics(t, func() {
		c = NewConfigFromFile("example-config.yaml")
	})

	require.NotNil(t, c)

	s := c.GetSettings()

	require.Equal(t, sampleToken, s.GetToken())
	require.Equal(t, "https://api.telegram.org/", s.GetURL())
	require.Len(t, s.GetAllowedUpdates(), 16)

	require.Equal(t, "fa", c.GetDefaultLocale())

	require.Panics(t, func() {
		c.L("fa", "hellloasidojoausbdyfigosdyuf")
	})

	require.Equal(t, c.L("en", "commandStart"), "Start")

	require.NotPanics(t, func() {
		c.GetBot().GetDefaultAdminRights(true)
		c.GetBot().GetDefaultAdminRights(false)
	})

	require.NotPanics(t, func() {
		eng := c.GetBot().GetCommands("en")

		for cmd, d := range eng {
			require.NotEmpty(t, cmd)
			require.NotEmpty(t, d)

			require.False(t, strings.HasPrefix(d, "locale:"))
		}
	})

	handles := c.GetHandles()

	t.Logf("locales: %v", c.GetLocales())

	for _, h := range handles {
		require.NotEmpty(t, h.GetCommand(c.GetLocales()...))
		han := h.GetHandler()
		key := han.GetKeyboard("en")
		require.False(t, strings.HasPrefix(han.GetText("en"), "locale:"))
		j, _ := key.MarshalJSON()
		t.Logf("Commands: %v, Keyboard: %s, Text: %s", h.GetCommand(c.GetLocales()...), j, han.GetText("en"))
	}

	t.Logf("%s", c.L("en", "hello_text", "hello! hello hello "))

}
