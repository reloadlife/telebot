package telebot

import "fmt"

// Animation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
type Animation struct {
	// FileID is the identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is the unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Width is the video width as defined by the sender.
	Width int `json:"width"`

	// Height is the video height as defined by the sender.
	Height int `json:"height"`

	// Duration is the duration of the video in seconds as defined by the sender.
	Duration int `json:"duration"`

	// Thumbnail is the animation thumbnail as defined by the sender (optional).
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// FileName is the original animation filename as defined by the sender (optional).
	FileName string `json:"file_name,omitempty"`

	// MimeType is the MIME type of the file as defined by the sender (optional).
	MimeType string `json:"mime_type,omitempty"`

	// FileSize is the file size in bytes. It can be bigger than 2^31, and some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value (optional).
	FileSize int64 `json:"file_size,omitempty"`
}

func (a *Animation) ReflectType() string { return fmt.Sprintf("%T", a) }
func (a *Animation) Type() string        { return "Animation" }

func (a *Animation) File() *File {
	f := FromFileID(a.FileID)
	f.UniqueID = a.FileUniqueID
	f.FileSize = a.FileSize
	f.fileName = a.FileName
	return &f
}

func (a *Animation) Send(b Bot, to Recipient, options ...any) (Message, error) {
	return b.SendAnimation(to, *a.File(), options...)
}
