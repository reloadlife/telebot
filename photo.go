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

func (p *PhotoSize) ReflectType() string { return fmt.Sprintf("%T", p) }
func (p *PhotoSize) Type() string        { return "PhotoSize" }

func (p *PhotoSize) File() *File {
	f := FromFileID(p.FileID)
	f.fileName = fmt.Sprintf("photo_%s.jpg", p.FileUniqueID)
	f.FileSize = p.FileSize
	f.UniqueID = p.FileUniqueID
	return &f
}

func (p *PhotoSize) Send(b Bot, to Recipient, options ...any) (Message, error) {
	return b.SendPhoto(to, *p.File(), options...)
}
