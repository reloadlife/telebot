package telebot

type answerInlineQueryRequest struct {
	ID      string       `json:"inline_query_id"`
	Results QueryResults `json:"results"`

	CacheTime  int               `json:"cache_time,omitempty"`
	IsPersonal bool              `json:"is_personal,omitempty"`
	NextOffset string            `json:"next_offset,omitempty"`
	Button     InlineQueryButton `json:"button,omitempty"`
}
