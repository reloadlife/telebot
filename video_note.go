package telebot

import "fmt"

// VideoNote objects represent video messages, available in Telegram apps as of v.4.0.
type VideoNote struct {
	// FileID is the identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file, supposed to be the same over time and for different bots.
	// It can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Length represents the video width and height (diameter of the video message) as defined by the sender.
	Length int `json:"length"`

	// Duration is the duration of the video in seconds as defined by the sender.
	Duration int `json:"duration"`

	// Thumbnail is an optional field representing the video thumbnail (PhotoSize).
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// FileSize is an optional field representing the file size in bytes.
	FileSize int64 `json:"file_size,omitempty"`
}

func (c *VideoNote) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *VideoNote) Type() string        { return "VideoNote" }

func (c *VideoNote) File() *File {
	f := FromFileID(c.FileID)
	f.fileName = fmt.Sprintf("video_note_%s.mp4", c.FileUniqueID)
	f.FileSize = c.FileSize
	f.UniqueID = c.FileUniqueID
	return &f
}
