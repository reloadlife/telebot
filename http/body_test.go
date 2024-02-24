package httpc

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBody(t *testing.T) {
	firstBody := NewBody(map[string]any{
		"hello": "hii",
		"world": "earth",
	})

	encoded, contentType := firstBody.Encode()
	require.Equal(t, bytes.NewReader([]byte(`{"hello":"hii","world":"earth"}`)), encoded)
	require.Equal(t, "application/json", contentType)

	firstBody.Add("hello", "world")

	encoded, contentType = firstBody.Encode()
	require.Equal(t, bytes.NewReader([]byte(`{"hello":"world","world":"earth"}`)), encoded)
	require.Equal(t, "application/json", contentType)

	firstBody.Delete("hello")

	encoded, contentType = firstBody.Encode()
	require.Equal(t, bytes.NewReader([]byte(`{"world":"earth"}`)), encoded)
	require.Equal(t, "application/json", contentType)

	data := firstBody.Get("world")

	require.Equal(t, "earth", data)

	firstBody.Add("File", []byte("hello world"))

	encoded, contentType = firstBody.Encode()

	require.NotEqual(t, "application/json", contentType)
}
