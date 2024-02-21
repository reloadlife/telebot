package telebot

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestButton(t *testing.T) {
	row := Row{
		&InlineKeyboardButton{Text: "Hello, World!", CallbackData: toPtr("hello-world")},
		&InlineKeyboardButton{Text: "Hello, World! 2", CallbackData: toPtr("hello-world-2")},
	}

	byt, err := json.Marshal(row)
	require.NoError(t, err)
	t.Logf("JSON: %s", byt)

	b := &Row{}

	err = json.Unmarshal(byt, b)
	require.NoError(t, err)
	t.Logf("Unmarshaled: %#v", b)

}

func TestMarkup(t *testing.T) {
	m := NewMarkup()

	m.Inline()

	b, err := json.Marshal(m)
	require.NoError(t, err)
	t.Logf("JSON: %s", b)

	m.AddRow(NewInlineKeyboardButton("Hello, World!", "hi"))

	m.AddRows(Row{
		NewInlineKeyboardButton("Hello, World! 2", "hi 2"),
		NewInlineKeyboardButton("Hello, World! 3", "hi 3"),
	})

	m.AddRow(
		NewInlineKeyboardButton("1", "6"),
		NewInlineKeyboardButton("2", "5"),
		NewInlineKeyboardButton("3", "4"),
	)

	b, err = json.Marshal(m)
	require.NoError(t, err)
	t.Logf("JSON: %s", b)

	m = &baseMarkUp{}
	err = m.UnmarshalJSON(b)
	require.NoError(t, err)
	t.Logf("Unmarshaled: %#v", m)
}
