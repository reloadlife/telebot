package telebot

import "fmt"

// Audio represents an audio file to be treated as music by the Telegram clients.
// <a href="https://core.telegram.org/bots/api#audio">/bots/api#audio</a>
// Wiki: <a href="https://github.com/reloadlife/telebot/wiki/Audio">/wiki/Audio</a>
type Audio struct {
	// FileID is the identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is the unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Duration is the duration of the audio in seconds as defined by the sender.
	Duration int `json:"duration"`

	// Performer is the performer of the audio as defined by the sender or by audio tags.
	Performer string `json:"performer,omitempty"`

	// Title is the title of the audio as defined by the sender or by audio tags.
	Title string `json:"title,omitempty"`

	// FileName is the original filename as defined by the sender.
	FileName string `json:"file_name,omitempty"`

	// MimeType is the MIME type of the file as defined by the sender.
	MimeType string `json:"mime_type,omitempty"`

	// FileSize is the file size in bytes.
	FileSize int64 `json:"file_size,omitempty"`

	// Thumbnail is the thumbnail of the album cover to which the music file belongs.
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}

// ReflectType returns the type of this object
func (a *Audio) ReflectType() string { return fmt.Sprintf("%T", a) }

// Type returns the type of this object
func (a *Audio) Type() string { return "Audio" }

// File returns a File object from this Audio
func (a *Audio) File() *File {
	f := FromFileID(a.FileID)
	f.FileSize = a.FileSize
	f.UniqueID = a.FileUniqueID
	if a.FileName == "" {
		// todo: check if it's MP3 or M4A from MIME
		a.FileName = fmt.Sprintf("audio_%s.m4a", a.FileID)
	}
	f.fileName = a.FileName
	return &f
}

// Send sends this audio to the recipient
func (a *Audio) Send(b Bot, to Recipient, options ...any) (Message, error) {
	return b.SendAudio(to, *a.File(), options...)
}
