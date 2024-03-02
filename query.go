package telebot

import "fmt"

type QueryResults []QueryResult

type QueryResult struct {
	// Type of the result
	QueryType InlineQueryResultType `json:"type"`

	// Unique identifier for this result, 1-64 Bytes
	ID string `json:"id"`

	// Title of the result (For those results that require a Title)
	Title *string `json:"title,omitempty"`

	// MessageContent Content of the message to be sent
	MessageContent InputMessageContent `json:"input_message_content,omitempty"`

	// ReplyMarkup Optional. Inline keyboard attached to the message.
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`

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

func (q *QueryResult) ReflectType() string { return fmt.Sprintf("%T", q) }
func (q *QueryResult) Type() string {
	if q.QueryType == "" {
		return Unknown
	}

	return string(q.QueryType)
}

type InlineQueryButton struct {
	Text       string      `json:"text"`
	WebApp     *WebAppInfo `json:"url,omitempty"`
	StartParam *string     `json:"start_parameter,omitempty"`
}
