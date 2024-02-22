package telebot

import (
	"fmt"
	"strings"
)

type Sticker struct {
	FileID           string        `json:"file_id"`
	FileUniqueID     string        `json:"file_unique_id"`
	StickerType      string        `json:"type"`
	Width            int           `json:"width"`
	Height           int           `json:"height"`
	IsAnimated       bool          `json:"is_animated"`
	IsVideo          bool          `json:"is_video"`
	Thumbnail        *PhotoSize    `json:"thumbnail,omitempty"`
	Emoji            string        `json:"emoji,omitempty"`
	SetName          string        `json:"set_name,omitempty"`
	PremiumAnimation *File         `json:"premium_animation,omitempty"`
	MaskPosition     *MaskPosition `json:"mask_position,omitempty"`
	CustomEmojiID    *string       `json:"custom_emoji_id,omitempty"`
	NeedsRepainting  *bool         `json:"needs_repainting,omitempty"`
	FileSize         int64         `json:"file_size,omitempty"`
}

func (c *Sticker) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *Sticker) Type() string        { return "Sticker" }

func (c *Sticker) File() *File {
	f := FromFileID(c.FileID)
	f.fileName = fmt.Sprintf("sticker_%s_%s.webp", strings.ReplaceAll(" ", "_", c.SetName), c.FileUniqueID)
	f.FileSize = c.FileSize
	f.UniqueID = c.FileUniqueID
	return &f
}

type MaskPosition struct {
	Point  string  `json:"point"`
	XShift float64 `json:"x_shift"`
	YShift float64 `json:"y_shift"`
	Scale  float64 `json:"scale"`
}

func (c *MaskPosition) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *MaskPosition) Type() string        { return "MaskPosition" }
