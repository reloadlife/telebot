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

func (v *VideoNote) ReflectType() string { return fmt.Sprintf("%T", v) }
func (v *VideoNote) Type() string        { return "VideoNote" }

func (v *VideoNote) File() *File {
	f := FromFileID(v.FileID)
	f.fileName = fmt.Sprintf("video_note_%s.mp4", v.FileUniqueID)
	f.FileSize = v.FileSize
	f.UniqueID = v.FileUniqueID
	return &f
}

func (v *VideoNote) Send(b Bot, to Recipient, options ...any) (Message, error) {
	return b.SendVideoNote(to, *v.File(), options...)
}
