package telebot

import (
	"encoding/json"
	"errors"
)

type QueryResults []QueryResult

type InlineQueryResultType string

const (
	InlineQueryResultCachedAudio    InlineQueryResultType = "audio"
	InlineQueryResultCachedDocument InlineQueryResultType = "document"
	InlineQueryResultCachedGif      InlineQueryResultType = "gif"
	InlineQueryResultCachedMpeg4Gif InlineQueryResultType = "mpeg4_gif"
	InlineQueryResultCachedPhoto    InlineQueryResultType = "photo"
	InlineQueryResultCachedSticker  InlineQueryResultType = "sticker"
	InlineQueryResultCachedVideo    InlineQueryResultType = "video"
	InlineQueryResultCachedVoice    InlineQueryResultType = "voice"
	InlineQueryResultArticle        InlineQueryResultType = "article"
	InlineQueryResultAudio          InlineQueryResultType = "audio"
	InlineQueryResultContact        InlineQueryResultType = "contact"
	InlineQueryResultGame           InlineQueryResultType = "game"
	InlineQueryResultDocument       InlineQueryResultType = "document"
	InlineQueryResultGif            InlineQueryResultType = "gif"
	InlineQueryResultLocation       InlineQueryResultType = "location"
	InlineQueryResultMpeg4Gif       InlineQueryResultType = "mpeg4_gif"
	InlineQueryResultPhoto          InlineQueryResultType = "photo"
	InlineQueryResultVenue          InlineQueryResultType = "venue"
	InlineQueryResultVideo          InlineQueryResultType = "video"
	InlineQueryResultVoice          InlineQueryResultType = "voice"
)

type QueryResult struct {
	// Type of the result
	Type InlineQueryResultType `json:"type"`

	// Unique identifier for this result, 1-64 Bytes
	ID string `json:"id"`

	// Title of the result (For those results that require a Title)
	Title *string `json:"title,omitempty"`

	// MessageContent Content of the message to be sent
	MessageContent InputMessageContent `json:"input_message_content,omitempty"`

	// ReplyMarkup Optional. Inline keyboard attached to the message.
	ReplyMarkup *ReplyMarkup `json:"reply_markup,omitempty"`

	// URL of the result
	URL *string `json:"url,omitempty"`

	// HideURL True if you don't want the URL to be shown in the message.
	HideURL *bool `json:"hide_url,omitempty"`

	// Short description of the result
	Description *string `json:"description,omitempty"`

	// ThumbnailURL of the thumbnail for the result
	ThumbnailURL *string `json:"thumbnail_url,omitempty"`

	// ThumbnailWidth width
	ThumbnailWidth *int `json:"thumbnail_width,omitempty"`

	// ThumbnailHeight height
	ThumbnailHeight *int `json:"thumbnail_height,omitempty"`
}

var (
	ErrBadInlineQueryResultType = errors.New("telebot: bad inline query result type")
)

func (i *QueryResult) MarshalJSON() ([]byte, error) {
	if err := i.Verify(); err != nil {
		return nil, err
	}

	return json.Marshal(i)
}

// Verify will verify the QueryResult struct and cleans it up before sending it to Telegram.
func (i *QueryResult) Verify() error {
	switch i.Type {

	case InlineQueryResultArticle:

	default:
		return ErrBadInlineQueryResultType
	}

	return nil
}
