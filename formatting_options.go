package telebot

import "fmt"

type ParseMode string

const (
	ParseModeDefault    ParseMode = ""
	ParseModeMarkdown   ParseMode = "Markdown"
	ParseModeMarkdownV2 ParseMode = "MarkdownV2"
	ParseModeHTML       ParseMode = "HTML"
)

type EntityType string

const (
	EntityTypeMention       EntityType = "mention"
	EntityTypeHashtag       EntityType = "hashtag"
	EntityTypeCashTag       EntityType = "cashtag"
	EntityTypeBotCommand    EntityType = "bot_command"
	EntityTypeURL           EntityType = "url"
	EntityTypeEmail         EntityType = "email"
	EntityTypePhoneNumber   EntityType = "phone_number"
	EntityTypeBold          EntityType = "bold"
	EntityTypeItalic        EntityType = "italic"
	EntityTypeUnderline     EntityType = "underline"
	EntityTypeStrikethrough EntityType = "strikethrough"
	EntityTypeSpoiler       EntityType = "spoiler"
	EntityTypeBlockquote    EntityType = "blockquote"
	EntityTypeCode          EntityType = "code"
	EntityTypePre           EntityType = "pre"
	EntityTypeTextLink      EntityType = "text_link"
	EntityTypeTextMention   EntityType = "text_mention"
	EntityTypeCustomEmoji   EntityType = "custom_emoji"
)

type Entity struct {
	EntityType    EntityType   `json:"type"`
	Offset        int          `json:"offset"`
	Length        int          `json:"length"`
	URL           string       `json:"url,omitempty"`
	User          *User        `json:"user,omitempty"`
	Language      *string      `json:"language,omitempty"`
	CustomEmojiID *CustomEmoji `json:"custom_emoji_id,omitempty"`
}

func (c *Entity) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *Entity) Type() string {
	if c.EntityType == "" {
		return "unknown"
	}
	return string(c.EntityType)
}

type LinkPreviewOptions struct {
	IsDisabled       *bool   `json:"is_disabled,omitempty"`
	URL              *string `json:"url,omitempty"`
	PreferSmallMedia *bool   `json:"prefer_small_media,omitempty"`
	PreferLargeMedia *bool   `json:"prefer_large_media,omitempty"`
	ShowAboveText    *bool   `json:"show_above_text,omitempty"`
}

func (c *LinkPreviewOptions) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *LinkPreviewOptions) Type() string        { return "LinkPreviewOptions" }
