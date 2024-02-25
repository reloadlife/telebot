package telebot

import (
	"fmt"
	"time"
)

func (b *bot) AnswerInlineQuery(queryID string, results QueryResults, options ...any) error {
	if len(results) > 50 {
		return fmt.Errorf("telebot: no more than 50 results are allowed in a single query. Got %d", len(results))
	}
	params := answerInlineQueryRequest{
		ID:      queryID,
		Results: results,
	}

	for _, option := range options {
		switch v := option.(type) {
		case time.Duration:
			params.CacheTime = int(v.Seconds())
		case string:
			params.NextOffset = v

		case InlineQueryButton:
			params.Button = v

		case Option:
			switch v {
			case InlineQueryIsPersonal:
				params.IsPersonal = true
			}

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in SendVoice.")
		}
	}

	_, err := b.sendMethodRequest(methodAnswerInlineQuery, params)
	return err
}
