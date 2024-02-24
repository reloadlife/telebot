package telebot

import "fmt"

// Document represents a general file (as opposed to photos, voice messages, and audio files).
type Document struct {
	// FileID is the identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is the unique identifier for this file, which is supposed to be the same over time and for different bots.
	// It can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Thumbnail is the document thumbnail as defined by the sender (optional).
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	// FileName is the original filename as defined by the sender (optional).
	FileName string `json:"file_name,omitempty"`

	// MimeType is the MIME type of the file as defined by the sender (optional).
	MimeType string `json:"mime_type,omitempty"`

	// FileSize is the file size in bytes. It can be bigger than 2^31, and some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type is safe for storing this value (optional).
	FileSize int64 `json:"file_size,omitempty"`
}

func (d *Document) ReflectType() string { return fmt.Sprintf("%T", d) }
func (d *Document) Type() string        { return "Document" }

func (d *Document) File() *File {
	f := FromFileID(d.FileID)
	f.fileName = d.FileName
	f.FileSize = d.FileSize
	f.UniqueID = d.FileUniqueID
	return &f
}

func (d *Document) Send(b Bot, to Recipient, options ...any) (Message, error) {
	return b.SendDocument(to, *d.File(), options...)
}
