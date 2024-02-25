package telebot

import (
	"encoding/json"
	"fmt"
)

func (b *bot) SendGame(to Recipient, gameShortName string, options ...any) (*AccessibleMessage, error) {
	params := sendGameRequest{
		ChatID:        to.Recipient(),
		GameShortName: gameShortName,
	}

	for _, option := range options {
		switch v := option.(type) {
		case *MessageThreadID:
			params.ThreadID = v
		case ReplyMarkup:
			params.ReplyMarkup = v
		case *ReplyParameters:
			params.ReplyParameters = v

		case Option:
			switch v {
			case Silent:
				params.DisableNotification = toPtr(true)
			case Protected:
				params.Protected = toPtr(true)
			}

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in SendGame.")
		}
	}

	var resp struct {
		Result *AccessibleMessage
	}

	req, err := b.sendMethodRequest(methodSendGame, params)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

func (b *bot) SetGameScore(msg Message, userID Userable, score uint, options ...any) (*AccessibleMessage, error) {
	chat, msgID := msg.MessageSig()
	params := setGameScore{
		ChatID:    chat,
		MessageID: msgID,
		UserID:    userID,
		Score:     score,
	}

	for _, option := range options {
		switch v := option.(type) {
		case Option:
			switch v {
			case ForceGameScoreUpdate:
				params.Force = toPtr(true)
			case DisableGameEditMessage:
				params.DisableEditMessage = toPtr(true)
			}

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in SetGameScore.")
		}
	}

	var resp struct {
		Result *AccessibleMessage
	}

	req, err := b.sendMethodRequest(methodSetGameScore, params)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

func (b *bot) SetGameScoreInline(inlineMessageID string, userID Userable, score uint, options ...any) error {
	params := setGameScore{
		InlineMessageID: inlineMessageID,
		UserID:          userID,
		Score:           score,
	}

	for _, option := range options {
		switch v := option.(type) {
		case Option:
			switch v {
			case ForceGameScoreUpdate:
				params.Force = toPtr(true)
			case DisableGameEditMessage:
				params.DisableEditMessage = toPtr(true)
			}

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in SetGameScore.")
		}
	}

	_, err := b.sendMethodRequest(methodSetGameScore, params)
	return err
}

func (b *bot) GetGameHighScores(user Userable, options ...any) ([]GameHighScore, error) {
	params := getGameHighScores{
		UserID: user,
	}

	for _, option := range options {
		switch v := option.(type) {
		case Message:
			chat, msgID := v.MessageSig()
			params.ChatID = chat
			params.MessageID = msgID
		case string:
			params.InlineMessageID = v
		}
	}

	var resp struct {
		Result []GameHighScore `json:"result"`
	}

	req, err := b.sendMethodRequest(methodGetGameHighScores, params)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}
	return resp.Result, nil
}
