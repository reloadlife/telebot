package telebot

// Audio represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	// FileID is the identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is the unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// Duration is the duration of the audio in seconds as defined by the sender.
	Duration int `json:"duration"`

	// Performer is the performer of the audio as defined by the sender or by audio tags (optional).
	Performer string `json:"performer,omitempty"`

	// Title is the title of the audio as defined by the sender or by audio tags (optional).
	Title string `json:"title,omitempty"`

	// FileName is the original filename as defined by the sender (optional).
	FileName string `json:"file_name,omitempty"`

	// MimeType is the MIME type of the file as defined by the sender (optional).
	MimeType string `json:"mime_type,omitempty"`

	// FileSize is the file size in bytes. It can be bigger than 2^31, and some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a signed 64-bit integer or double-precision float type are safe for storing this value (optional).
	FileSize int `json:"file_size,omitempty"`

	// Thumbnail is the thumbnail of the album cover to which the music file belongs (optional).
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}
