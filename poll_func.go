package telebot

import (
	"encoding/json"
	"fmt"
	"time"
)

func (b *bot) SendPoll(to Recipient, question string, answerOptions []string, options ...any) (*AccessibleMessage, error) {
	if len(answerOptions) < 2 || len(answerOptions) > 10 {
		panic("telebot: answerOptions must be between 2 and 10")
	}
	for _, option := range answerOptions {
		if len(option) < 1 || len(option) > 100 {
			panic("telebot: answerOptions must be between 1 and 100 characters")
		}
	}
	params := sendPollRequest{
		ChatID:   to.Recipient(),
		Question: question,
		Options:  answerOptions,
	}

	for _, option := range options {
		switch v := option.(type) {
		case *MessageThreadID:
			params.ThreadID = v
		case ReplyMarkup:
			params.ReplyMarkup = v
		case *ReplyParameters:
			params.ReplyParameters = v
		case PollType:
			params.Type = &v
		case *PollType:
			params.Type = v
		case int:
			params.CorrectOptionID = toPtr(v)
		case string:
			params.Explanation = toPtr(v)
		case []Entity:
			params.ExplanationEntities = v
		case ParseMode:
			params.ExplanationParseMode = &v
		case time.Duration:
			if v.Seconds() > 600 || v.Seconds() < 5 {
				panic("telebot: open period must be between 5 and 600 seconds")
			}
			params.OpenPeriod = toPtr(int(v.Seconds()))

		case Option:
			switch v {
			case Silent:
				params.DisableNotification = toPtr(true)
			case Protected:
				params.Protected = toPtr(true)
			case IsAnonymousPoll:
				params.IsAnonymous = toPtr(true)
			case AllowMultipleAnswers:
				params.AllowMultipleAnswers = toPtr(true)
			case IsClosedPoll:
				params.IsClosed = toPtr(true)
			}

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in SendPoll.")
		}
	}

	var resp struct {
		Result *AccessibleMessage
	}

	req, err := b.sendMethodRequest(methodSendPoll, params)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
