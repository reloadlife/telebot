package telebot

import "fmt"

// Voice represents a voice note.
type Voice struct {
	// FileID is the identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file, supposed to be the same over time and for different bots.
	// It can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Duration is the duration of the audio in seconds as defined by the sender.
	Duration int `json:"duration"`

	// MimeType is an optional field representing the MIME type of the file as defined by the sender.
	MimeType *string `json:"mime_type,omitempty"`

	// FileSize is an optional field representing the file size in bytes.
	// It can be bigger than 2^31, and some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type is safe for storing this value.
	FileSize int64 `json:"file_size,omitempty"`
}

func (v *Voice) ReflectType() string { return fmt.Sprintf("%T", v) }
func (v *Voice) Type() string        { return "Voice" }

func (v *Voice) File() *File {
	f := FromFileID(v.FileID)
	f.fileName = fmt.Sprintf("voice_%s.ogg", v.FileUniqueID)
	f.FileSize = v.FileSize
	f.UniqueID = v.FileUniqueID
	return &f
}

func (v *Voice) Send(b Bot, to Recipient, options ...any) (Message, error) {
	return b.SendVoice(to, *v.File(), options...)
}
