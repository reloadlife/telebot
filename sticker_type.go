package telebot

import "fmt"

// StickerSet represents a sticker set.
type StickerSet struct {
	// Name represents the sticker set name.
	Name string `json:"name"`

	// Title represents the sticker set title.
	Title string `json:"title"`

	// StickerType represents the type of stickers in the set, currently one of "regular", "mask", "custom_emoji".
	StickerType string `json:"sticker_type"`

	// IsAnimated is true if the sticker set contains animated stickers.
	IsAnimated bool `json:"is_animated"`

	// IsVideo is true if the sticker set contains video stickers.
	IsVideo bool `json:"is_video"`

	// Stickers is a list of all set stickers.
	Stickers []Sticker `json:"stickers"`

	// Thumbnail is an optional sticker set thumbnail in the .WEBP, .TGS, or .WEBM format.
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}

func (c *StickerSet) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *StickerSet) Type() string        { return "StickerSet" }

type InputSticker struct {
	Sticker      File           `json:"sticker"`
	EmojiList    []StickerEmoji `json:"emoji_list,omitempty"`
	MaskPosition *MaskPosition  `json:"mask_position,omitempty"`
	Keywords     []string       `json:"keywords,omitempty"`
}

func (c *InputSticker) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *InputSticker) Type() string        { return "InputSticker" }
