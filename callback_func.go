package telebot

import (
	"net/url"
	"time"
)

func (b *bot) AnswerCallbackQuery(callback *Callback, opts ...any) error {
	params := answerCallbackQuery{
		ID: callback.ID,
	}

	for _, opt := range opts {
		switch o := opt.(type) {
		case string:
			params.Text = o
		case bool:
			params.Show = o
		case time.Duration:
			params.Cache = int(o.Seconds())
		case *url.URL:
			params.URL = o.String()
		}
	}

	_, err := b.sendMethodRequest(methodAnswerCallbackQuery, params)
	return err
}
