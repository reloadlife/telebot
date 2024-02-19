package telebot

import (
	"encoding/json"
	"errors"
)

type InputMediaType string

const (
	InputMediaAnimation InputMediaType = "animation"
	InputMediaDocument  InputMediaType = "document"
	InputMediaAudio     InputMediaType = "audio"
	InputMediaPhoto     InputMediaType = "photo"
	InputMediaVideo     InputMediaType = "video"
)

type InputMedia struct {
	Type  InputMediaType `json:"type"`
	Media File           `json:"media"`

	Caption   *string    `json:"caption,omitempty"`
	ParseMode *ParseMode `json:"parse_mode,omitempty"`
	Entities  []Entity   `json:"caption_entities,omitempty"`

	HasSpoiler *bool `json:"has_spoiler,omitempty"`

	Thumb *File `json:"thumbnail,omitempty"`

	Width    *int `json:"width,omitempty"`
	Height   *int `json:"height,omitempty"`
	Duration *int `json:"duration,omitempty"`

	Streamable *bool `json:"supports_streaming,omitempty"`

	Preformer *string `json:"performer,omitempty"`
	Title     *string `json:"title,omitempty"`

	DisableContentTypeDetection *bool `json:"disable_content_type_detection,omitempty"`
}

var (
	ErrBadMediaType = errors.New("bad media type")
)

func (i *InputMedia) MarshalJSON() ([]byte, error) {
	if err := i.Verify(); err != nil {
		return nil, err
	}

	return json.Marshal(i)
}

// Verify will verify the InputMedia struct and cleans it up before sending it to Telegram.
func (i *InputMedia) Verify() error {
	switch i.Type {
	case InputMediaAnimation:
		i.Thumb = nil
		i.Width = nil
		i.Height = nil
		i.Duration = nil
		i.Streamable = nil
		i.DisableContentTypeDetection = nil

	case InputMediaDocument:
		i.Width = nil
		i.Height = nil

	case InputMediaAudio:
		i.Height = nil
		i.Duration = nil
		i.Preformer = nil
		i.Title = nil
		i.DisableContentTypeDetection = nil

	case InputMediaPhoto:
		i.Thumb = nil
		i.Width = nil
		i.Height = nil
		i.Duration = nil
		i.Streamable = nil
		i.Preformer = nil
		i.Title = nil
		i.DisableContentTypeDetection = nil

	case InputMediaVideo:
		i.Thumb = nil
		i.Width = nil
		i.Height = nil
		i.Duration = nil
		i.Streamable = nil
		i.DisableContentTypeDetection = nil

	default:
		return ErrBadMediaType
	}

	return nil
}
