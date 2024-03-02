package tutil

import "strings"

func EscapeMarkdown(text string) string {
	// EscapeMarkdownV2 escapes special characters in the provided text and returns the escaped text.
	// The following characters are escaped:
	// - `_` (underscore)
	// - `*` (asterisk)
	// - `[` (left square bracket)
	// - `]` (right square bracket)
	// - `(` (left parenthesis)
	// - `)` (right parenthesis)
	// - `~` (tilde)
	// - `` ` `` (backtick)
	// - `>` (greater than)
	// - `#` (hash)
	// - `+` (plus)
	// - `-` (minus)
	// - `=` (equals)
	// - `|` (pipe)
	// - `{` (left curly brace)
	// - `}` (right curly brace)
	// - `.` (dot)
	// - `!` (exclamation mark)
	//
	// Note that the backslash character `\` is not escaped because it is used to escape characters in Markdown.
	//
	// Example
	//
	// 	text := "Hello, _world_!"
	// 	escapedText := tutil.EscapeMarkdownV2(text)
	// 	fmt.Println(escapedText)
	//
	// Output:
	//
	// 	Hello, \_world\_!
	//
	// Note that the backslash character `\` is used to escape the underscore characters `_` in the input text.

	escapeMap := map[string]string{
		"_": "\\_",
		"*": "\\*",
		"[": "\\[",
		"]": "\\]",
		"(": "\\(",
		")": "\\)",
		"~": "\\~",
		"`": "\\`",
		">": "\\>",
		"#": "\\#",
		"+": "\\+",
		"-": "\\-",
		"=": "\\=",
		"|": "\\|",
		"{": "\\{",
		"}": "\\}",
		".": "\\.",
		"!": "\\!",
	}

	for key, value := range escapeMap {
		text = strings.ReplaceAll(text, key, value)
	}

	return text
}
