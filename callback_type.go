package telebot

type answerCallbackQuery struct {
	ID    string `json:"callback_query_id"`
	Text  string `json:"text,omitempty"`
	Show  bool   `json:"show_alert,omitempty"`
	URL   string `json:"url,omitempty"`
	Cache int    `json:"cache_time,omitempty"`
}
