package telebot

import "fmt"

// Animation represents an animation file (GIF or H.264/MPEG-4 AVC video without sound).
// <a href="https://core.telegram.org/bots/api#animation">/bots/api#animation</a>
// Wiki: <a href="https://github.com/reloadlife/telebot/wiki/Animation">/wiki/Animation</a>
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

	// Thumbnail is the animation thumbnail as defined by the sender.
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// FileName is the original animation filename as defined by the sender.
	FileName string `json:"file_name,omitempty"`

	// MIME is the MIME type of the file as defined by the sender.
	MIME string `json:"mime_type,omitempty"`

	// Size is the file size in bytes.
	Size int64 `json:"file_size,omitempty"`
}

// ReflectType returns the type of this object
func (a *Animation) ReflectType() string { return fmt.Sprintf("%T", a) }

// Type returns the type of this object
func (a *Animation) Type() string { return "Animation" }

// File returns a pointer to File
func (a *Animation) File() *File {
	f := FromFileID(a.FileID)
	f.UniqueID = a.FileUniqueID
	f.FileSize = a.Size
	if a.FileName == "" {
		a.FileName = fmt.Sprintf("animation_%s.mp4", a.FileID)
	}
	f.fileName = a.FileName
	return &f
}

// Send sends an animation to the recipient with Options.
// for available Options please refer to <a href="https://">Available Animation Options</a>
func (a *Animation) Send(b Bot, to Recipient, options ...any) (Message, error) {
	return b.SendAnimation(to, *a.File(), options...)
}
