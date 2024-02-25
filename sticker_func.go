package telebot

import (
	"encoding/json"
	"fmt"
)

func (b *bot) SendSticker(to Recipient, sticker File, options ...any) (*AccessibleMessage, error) {
	params := sendStickerRequest{
		ChatID:  to.Recipient(),
		Sticker: &sticker,
	}

	for _, option := range options {
		switch v := option.(type) {
		case *MessageThreadID:
			params.ThreadID = v
		case StickerEmoji:
			params.Emoji = string(v)
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
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in SendSticker.")
		}
	}

	var resp struct {
		Result *AccessibleMessage
	}

	req, err := b.sendMethodRequest(methodSendSticker, params)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

func (b *bot) GetStickerSet(name string) (*StickerSet, error) {
	params := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	var resp struct {
		Result *StickerSet
	}

	req, err := b.sendMethodRequest(methodGetStickerSet, params)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}

func (b *bot) UploadStickerFile(user Userable, sticker File, Format string) (*File, error) {
	params := uploadStickerFileParams{
		UserID:        user,
		Sticker:       &sticker,
		StickerFormat: StickerFormat(Format),
	}

	var resp struct {
		Result *File
	}

	req, err := b.sendMethodRequest(methodUploadStickerFile, params)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(req, &resp); err != nil {
		return nil, err
	}

	return resp.Result, nil
}
