package telebot

import (
	"encoding/json"
	"fmt"
	httpc "go.mamad.dev/telebot/http"
	"io"
	"strconv"
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

func (b *bot) UploadStickerFile(user Userable, sticker File, format StickerFormat) (*File, error) {
	params := uploadStickerFileParams{
		UserID:        user,
		Sticker:       &sticker,
		StickerFormat: format,
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

// CreateNewStickerSet
// todo: this function needs to be tests via uploading files (it works with URL and fileID)
// if it didn't work, either provide a fix or Open an issue at https://github.com/reloadlife/telebot/issues/new
func (b *bot) CreateNewStickerSet(user Userable, name, title string, stickers []InputSticker, format StickerFormat, options ...any) error {
	if len(stickers) < 1 || len(stickers) > 50 {
		panic("telebot: stickers count must be between 1 and 50")
	}

	files := make([]httpc.File, 0, len(stickers))

	for i, s := range stickers {
		switch {
		case s.Sticker.GetFileReader() != nil:
			s.Repr = "attach://" + strconv.Itoa(i)
			r, _ := io.ReadAll(s.Sticker.GetFileReader())
			files = append(files, httpc.File{
				Name:     strconv.Itoa(i),
				FileName: s.Sticker.GetFileName(),
				DATA:     r,
			})
		}
	}

	params := createNewStickerSetParams{
		UserID:        user,
		Name:          name,
		Title:         title,
		Stickers:      stickers,
		StickerFormat: format,
	}

	for _, option := range options {
		switch v := option.(type) {
		case StickerType:
			params.StickerType = v

		case Option:
			switch v {
			case StickerNeedsRepainting:
				params.NeedsRepainting = true
			}

		default:
			panic("telebot: unknown option type " + fmt.Sprintf("%T", v) + " in CreateNewStickerSet.")
		}
	}

	_, err := b.sendMethodRequest(methodCreateNewStickerSet, params)
	return err
}

func (b *bot) AddStickerToSet(user Userable, name string, sticker InputSticker) error {
	params := addStickerToSetParams{
		UserID:  user,
		Name:    name,
		Sticker: sticker,
	}

	_, err := b.sendMethodRequest(methodAddStickerToSet, params)
	return err
}

func (b *bot) SetStickerPositionInSet(sticker string, position int) error {
	params := struct {
		Sticker  string `json:"sticker"`
		Position int    `json:"position"`
	}{
		Sticker:  sticker,
		Position: position,
	}

	_, err := b.sendMethodRequest(methodSetStickerPositionInSet, params)
	return err
}

func (b *bot) DeleteStickerFromSet(sticker string) error {
	params := struct {
		Sticker string `json:"sticker"`
	}{
		Sticker: sticker,
	}

	_, err := b.sendMethodRequest(methodDeleteStickerFromSet, params)
	return err
}
