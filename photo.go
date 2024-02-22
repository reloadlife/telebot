package telebot

import "fmt"

type PhotoSizes []PhotoSize

type PhotoSize struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	FileSize     int64  `json:"file_size,omitempty"`
}

func (c *PhotoSize) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *PhotoSize) Type() string        { return "PhotoSize" }

func (c *PhotoSize) File() *File {
	f := FromFileID(c.FileID)
	f.fileName = fmt.Sprintf("photo_%s.jpg", c.FileUniqueID)
	f.FileSize = c.FileSize
	f.UniqueID = c.FileUniqueID
	return &f
}
