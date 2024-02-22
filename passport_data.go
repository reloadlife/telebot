package telebot

import "fmt"

// PassportData describes Telegram Passport data shared with the bot by the user.
type PassportData struct {
	// Data is an array with information about documents and other Telegram Passport elements that were shared with the bot.
	Data []EncryptedPassportElement `json:"data"`

	// Credentials encrypted credentials required to decrypt the data.
	Credentials EncryptedCredentials `json:"credentials"`
}

func (c *PassportData) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *PassportData) Type() string        { return "PassportData" }

// PassportFile represents a file uploaded to Telegram Passport. Currently, all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	// FileID is the identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`

	// FileUniqueID is a unique identifier for this file, supposed to be the same over time and for different bots.
	// It can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	// FileSize is the file size in bytes.
	FileSize int64 `json:"file_size"`

	// FileDate is the Unix time when the file was uploaded.
	FileDate int `json:"file_date"`
}

func (c *PassportFile) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *PassportFile) Type() string        { return "PassportFile" }

func (c *PassportFile) File() *File {
	f := FromFileID(c.FileID)
	f.UniqueID = c.FileUniqueID
	f.FileSize = c.FileSize
	f.fileName = fmt.Sprintf("passport_data_%s.jpg", c.FileUniqueID)
	return &f
}

// EncryptedPassportElement describes documents or other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {
	// ElementType is the element type, one of the predefined types.
	ElementType string `json:"type"`

	// Data is optional. Base64-encoded encrypted Telegram Passport element data provided by the user, available for certain types.
	// Can be decrypted and verified using the accompanying EncryptedCredentials.
	Data string `json:"data,omitempty"`

	// PhoneNumber is optional. User's verified phone number, available only for "phone_number" type.
	PhoneNumber string `json:"phone_number,omitempty"`

	// Email is optional. User's verified email address, available only for "email" type.
	Email string `json:"email,omitempty"`

	// Files is an array of encrypted files with documents provided by the user, available for certain types.
	// Files can be decrypted and verified using the accompanying EncryptedCredentials.
	Files []PassportFile `json:"files,omitempty"`

	// FrontSide is an encrypted file with the front side of the document, provided by the user.
	// Available for certain types. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	FrontSide *PassportFile `json:"front_side,omitempty"`

	// ReverseSide is an encrypted file with the reverse side of the document, provided by the user.
	// Available for certain types. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	ReverseSide *PassportFile `json:"reverse_side,omitempty"`

	// Selfie is an encrypted file with the selfie of the user holding a document, provided by the user.
	// Available for certain types. The file can be decrypted and verified using the accompanying EncryptedCredentials.
	Selfie *PassportFile `json:"selfie,omitempty"`

	// Translation is an array of encrypted files with translated versions of documents provided by the user.
	// Available if requested for certain types. Files can be decrypted and verified using the accompanying EncryptedCredentials.
	Translation []PassportFile `json:"translation,omitempty"`

	// Hash is the Base64-encoded element hash for using in PassportElementErrorUnspecified.
	Hash string `json:"hash"`
}

func (c *EncryptedPassportElement) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *EncryptedPassportElement) Type() string        { return "EncryptedPassportElement" }

// EncryptedCredentials describes data required for decrypting and authenticating EncryptedPassportElement.
type EncryptedCredentials struct {
	// Data is Base64-encoded encrypted JSON-serialized data with unique user's payload, data hashes, and secrets required for EncryptedPassportElement decryption and authentication.
	Data string `json:"data"`

	// Hash is Base64-encoded data hash for data authentication.
	Hash string `json:"hash"`

	// Secret is Base64-encoded secret, encrypted with the bot's public RSA key, required for data decryption.
	Secret string `json:"secret"`
}

func (c *EncryptedCredentials) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *EncryptedCredentials) Type() string        { return "EncryptedCredentials" }

type PassportElementError struct {
	// todo:
}

func (c *PassportElementError) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *PassportElementError) Type() string        { return "PassportElementError" }
