package telebot

import "fmt"

type sendStickerRequest struct {
	ChatID  any    `json:"chat_id"`
	Sticker *File  `json:"sticker" file:"1"`
	Emoji   string `json:"emoji,omitempty"`

	BusinessID          *BusinessID      `json:"business_connection_id,omitempty"`
	ThreadID            *MessageThreadID `json:"message_thread_id,omitempty"`
	DisableNotification *bool            `json:"disable_notification,omitempty"`
	Protected           *bool            `json:"protect_content,omitempty"`
	ReplyParameters     *ReplyParameters `json:"reply_parameters,omitempty"`
	ReplyMarkup         ReplyMarkup      `json:"reply_markup,omitempty"`
}

type uploadStickerFileParams struct {
	UserID        any           `json:"user_id"`
	Sticker       *File         `json:"sticker" file:"1"`
	StickerFormat StickerFormat `json:"sticker_format,omitempty"`
}

type createNewStickerSetParams struct {
	UserID          any            `json:"user_id"`
	Name            string         `json:"name"`
	Title           string         `json:"title"`
	Stickers        []InputSticker `json:"stickers"`
	StickerType     StickerType    `json:"sticker_type"`
	NeedsRepainting bool           `json:"needs_repainting,omitempty"`
}

type addStickerToSetParams struct {
	UserID  any          `json:"user_id"`
	Name    string       `json:"name"`
	Sticker InputSticker `json:"sticker"`
}

// StickerSet represents a sticker set.
type StickerSet struct {
	// Name represents the sticker set name.
	Name string `json:"name"`

	// Title represents the sticker set title.
	Title string `json:"title"`

	// StickerType represents the type of stickers in the set, currently one of "regular", "mask", "custom_emoji".
	StickerType string `json:"sticker_type"`

	// Stickers is a list of all set stickers.
	Stickers []Sticker `json:"stickers"`

	// Thumbnail is an optional sticker set thumbnail in the .WEBP, .TGS, or .WEBM format.
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}

func (c *StickerSet) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *StickerSet) Type() string        { return "StickerSet" }

type InputSticker struct {
	Sticker      File           `json:"sticker" file:"1"`
	EmojiList    []StickerEmoji `json:"emoji_list,omitempty"`
	MaskPosition *MaskPosition  `json:"mask_position,omitempty"`
	Keywords     []string       `json:"keywords,omitempty"`
	Format       StickerFormat  `json:"format,omitempty"`
	Repr         string         `json:"-"`
}

func (s *InputSticker) FileRepresent() string {
	return s.Repr
}

func (s *InputSticker) ReflectType() string { return fmt.Sprintf("%T", s) }
func (s *InputSticker) Type() string        { return "InputSticker" }
