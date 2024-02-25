package telebot

import (
	"encoding/json"
	"fmt"
)

func (b *bot) EditMessageText(msg Message, text string, options ...any) (*AccessibleMessage, error) {
	chat, msgID := msg.MessageSig()
	params := editMessageTextParams{
		ChatID:    chat,
		MessageID: msgID,
		Text:      text,
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case ParseMode:
			params.ParseMode = &v

		case []Entity:
			params.Entities = v

		case *LinkPreviewOptions:
			params.LinkPreviewOptions = v

		case ReplyMarkup:
			params.ReplyMarkup = v

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in EditMessageText.")
		}
	}

	var result struct {
		Result AccessibleMessage `json:"result"`
	}
	data, err := b.sendMethodRequest(methodEditMessageText, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return &result.Result, err
}

func (b *bot) EditMessageTextInline(inlineMessageID string, text string, options ...any) error {
	params := editMessageTextParams{
		InlineMessageID: inlineMessageID,
		Text:            text,
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case ParseMode:
			params.ParseMode = &v

		case []Entity:
			params.Entities = v

		case *LinkPreviewOptions:
			params.LinkPreviewOptions = v

		case ReplyMarkup:
			params.ReplyMarkup = v

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in EditMessageText (inline).")
		}
	}

	_, err := b.sendMethodRequest(methodEditMessageText, params)
	if err != nil {
		return err
	}
	return err
}

func (b *bot) EditMessageCaption(msg Message, caption string, options ...any) (*AccessibleMessage, error) {
	chat, msgID := msg.MessageSig()
	params := editMessageCaptionParams{
		ChatID:    chat,
		MessageID: msgID,
		Caption:   caption,
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case ParseMode:
			params.ParseMode = &v

		case []Entity:
			params.Entities = v

		case ReplyMarkup:
			params.ReplyMarkup = v

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in EditMessageCaption.")
		}
	}

	var result struct {
		Result AccessibleMessage `json:"result"`
	}
	data, err := b.sendMethodRequest(methodEditMessageCaption, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return &result.Result, err
}

func (b *bot) EditMessageCaptionInline(inlineMessageID string, caption string, options ...any) error {
	params := editMessageCaptionParams{
		InlineMessageID: inlineMessageID,
		Caption:         caption,
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case ParseMode:
			params.ParseMode = &v

		case []Entity:
			params.Entities = v

		case ReplyMarkup:
			params.ReplyMarkup = v

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in EditMessageCaption (inline).")
		}
	}

	_, err := b.sendMethodRequest(methodEditMessageCaption, params)
	if err != nil {
		return err
	}
	return err
}

func (b *bot) EditMessageMedia(msg Message, media InputMedia, options ...any) (*AccessibleMessage, error) {
	chat, msgID := msg.MessageSig()
	err := media.Verify()
	if err != nil {
		panic("telebot: bad InputMedia: " + err.Error())
	}
	params := editMessageMediaParams{
		ChatID:    chat,
		MessageID: msgID,
		Media:     media,
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case ReplyMarkup:
			params.ReplyMarkup = v

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in EditMessageMedia.")
		}
	}

	var result struct {
		Result AccessibleMessage `json:"result"`
	}
	data, err := b.sendMethodRequest(methodEditMessageMedia, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return &result.Result, err
}

func (b *bot) EditMessageMediaInline(inlineMessageID string, media InputMedia, options ...any) error {
	err := media.Verify()
	if err != nil {
		panic("telebot: bad InputMedia: " + err.Error())
	}
	params := editMessageMediaParams{
		InlineMessageID: inlineMessageID,
		Media:           media,
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case ReplyMarkup:
			params.ReplyMarkup = v

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in EditMessageMedia (inline).")
		}
	}

	_, err = b.sendMethodRequest(methodEditMessageMedia, params)
	if err != nil {
		return err
	}
	return err
}

func (b *bot) EditMessageLiveLocation(msg Message, location Location, options ...any) (*AccessibleMessage, error) {
	chat, msgID := msg.MessageSig()
	params := editMessageLiveLocationParams{
		ChatID:               chat,
		MessageID:            msgID,
		Latitude:             location.Latitude,
		Longitude:            location.Longitude,
		HorizontalAccuracy:   location.HorizontalAccuracy,
		Heading:              location.Heading,
		ProximityAlertRadius: location.ProximityAlertRadius,
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case ReplyMarkup:
			params.ReplyMarkup = v

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in EditMessageLiveLocation.")
		}
	}

	var result struct {
		Result AccessibleMessage `json:"result"`
	}
	data, err := b.sendMethodRequest(methodEditMessageLiveLocation, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return &result.Result, err
}

func (b *bot) EditMessageLiveLocationInline(inlineMessageID string, location Location, options ...any) error {
	params := editMessageLiveLocationParams{
		InlineMessageID:      inlineMessageID,
		Latitude:             location.Latitude,
		Longitude:            location.Longitude,
		HorizontalAccuracy:   location.HorizontalAccuracy,
		Heading:              location.Heading,
		ProximityAlertRadius: location.ProximityAlertRadius,
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case ReplyMarkup:
			params.ReplyMarkup = v

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in EditMessageLiveLocation (inline).")
		}
	}

	_, err := b.sendMethodRequest(methodEditMessageLiveLocation, params)
	if err != nil {
		return err
	}
	return err
}

func (b *bot) StopMessageLiveLocation(msg Message, options ...any) (*AccessibleMessage, error) {
	chat, msgID := msg.MessageSig()
	params := editMessageParams{
		ChatID:    chat,
		MessageID: msgID,
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case ReplyMarkup:
			params.ReplyMarkup = v

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in StopMessageLiveLocation.")
		}
	}

	var result struct {
		Result AccessibleMessage `json:"result"`
	}

	data, err := b.sendMethodRequest(methodStopMessageLiveLocation, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return &result.Result, err
}

func (b *bot) StopMessageLiveLocationInline(inlineMessageID string, options ...any) error {
	params := editMessageParams{
		InlineMessageID: inlineMessageID,
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case ReplyMarkup:
			params.ReplyMarkup = v

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in StopMessageLiveLocation (inline).")
		}
	}

	_, err := b.sendMethodRequest(methodStopMessageLiveLocation, params)
	if err != nil {
		return err
	}
	return err
}

func (b *bot) EditMessageReplyMarkup(msg Message, markup ReplyMarkup) (*AccessibleMessage, error) {
	chat, msgID := msg.MessageSig()
	params := editMessageParams{
		ChatID:      chat,
		MessageID:   msgID,
		ReplyMarkup: markup,
	}

	var result struct {
		Result AccessibleMessage `json:"result"`
	}

	data, err := b.sendMethodRequest(methodStopMessageLiveLocation, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return &result.Result, err
}

func (b *bot) EditMessageReplyMarkupInline(inlineMessageID string, markup ReplyMarkup) error {
	params := editMessageParams{
		InlineMessageID: inlineMessageID,
		ReplyMarkup:     markup,
	}

	_, err := b.sendMethodRequest(methodStopMessageLiveLocation, params)
	if err != nil {
		return err
	}
	return err
}

func (b *bot) StopPoll(msg Message, options ...any) (*Poll, error) {
	chat, msgID := msg.MessageSig()
	params := editMessageParams{
		ChatID:    chat,
		MessageID: msgID,
	}

	for _, opt := range options {
		switch v := opt.(type) {
		case ReplyMarkup:
			params.ReplyMarkup = v

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in StopPoll.")
		}
	}

	var result struct {
		Result Poll `json:"result"`
	}

	data, err := b.sendMethodRequest(methodStopPoll, params)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &result)
	return &result.Result, err
}

func (b *bot) DeleteMessage(msg Message, messageIDs ...int64) error {
	chat, msgID := msg.MessageSig()
	params := deleteMessageParams{
		ChatID:     chat,
		MessageIDs: append([]int64{msgID}, messageIDs...),
	}

	_, err := b.sendMethodRequest(methodDeleteMessages, params)
	return err
}
