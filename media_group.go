package telebot

import "fmt"

type InputMedia struct {
	Repr      string         `json:"-"`
	MediaType InputMediaType `json:"type"`
	Media     *File          `json:"media"`

	Caption   *string   `json:"caption,omitempty"`
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	Entities  []Entity  `json:"caption_entities,omitempty"`

	HasSpoiler *bool `json:"has_spoiler,omitempty"`

	Thumb *File `json:"thumbnail,omitempty" file:"1"`

	Width    *int `json:"width,omitempty"`
	Height   *int `json:"height,omitempty"`
	Duration *int `json:"duration,omitempty"`

	Streamable *bool `json:"supports_streaming,omitempty"`

	Preformer *string `json:"performer,omitempty"`
	Title     *string `json:"title,omitempty"`

	DisableContentTypeDetection *bool `json:"disable_content_type_detection,omitempty"`
}

func (m *InputMedia) FileRepresent() string { return m.Repr }

func (m *InputMedia) ReflectType() string { return fmt.Sprintf("%T", m) }
func (m *InputMedia) Type() string {
	if m.MediaType == "" {
		return Unknown
	}
	return string(m.MediaType)
}
