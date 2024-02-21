package telebot

import "fmt"

// ChatPhoto object represents a chat photo.
type ChatPhoto struct {
	// File identifiers of small (160x160) chat photo
	SmallFileID   string `json:"small_file_id"`
	SmallUniqueID string `json:"small_file_unique_id"`

	// File identifiers of big (640x640) chat photo
	BigFileID   string `json:"big_file_id"`
	BigUniqueID string `json:"big_file_unique_id"`
}

func (c *ChatPhoto) Type() string        { return "chat_photo" }
func (c *ChatPhoto) ReflectType() string { return fmt.Sprintf("%T", c) }
