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

func (b *bot) GetCustomEmojiStickers(customEmojiIDs ...CustomEmoji) ([]Sticker, error) {
	params := struct {
		CustomEmojiIDs []CustomEmoji `json:"custom_emoji_ids"`
	}{
		CustomEmojiIDs: customEmojiIDs,
	}

	var resp struct {
		Result []Sticker
	}

	req, err := b.sendMethodRequest(methodGetCustomEmojiStickers, params)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(req, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}

func (b *bot) SetStickerEmojiList(sticker string, emojiList []StickerEmoji) error {
	if len(emojiList) < 1 || len(emojiList) > 20 {
		panic("telebot: emojis count must be between 1 and 50")
	}

	params := struct {
		Sticker string         `json:"sticker"`
		Emojis  []StickerEmoji `json:"emojis"`
	}{
		Sticker: sticker,
		Emojis:  emojiList,
	}

	_, err := b.sendMethodRequest(methodSetStickerEmojiList, params)
	return err
}

func (b *bot) SetStickerKeywords(sticker string, keywords []string) error {
	if len(keywords) < 1 || len(keywords) > 20 {
		panic("telebot: keywords count must be between 1 and 50")
	}

	params := struct {
		Sticker  string   `json:"sticker"`
		Keywords []string `json:"keywords"`
	}{
		Sticker:  sticker,
		Keywords: keywords,
	}

	_, err := b.sendMethodRequest(methodSetStickerKeywords, params)
	return err
}

func (b *bot) SetStickerMaskPosition(sticker string, maskPosition ...MaskPosition) error {
	params := struct {
		Sticker      string       `json:"sticker"`
		MaskPosition MaskPosition `json:"mask_position,omitempty"`
	}{
		Sticker: sticker,
	}

	if len(maskPosition) > 0 {
		params.MaskPosition = maskPosition[0]
	}

	_, err := b.sendMethodRequest(methodSetStickerMaskPosition, params)
	return err
}

func (b *bot) SetStickerSetTitle(sticker, title string) error {
	params := struct {
		Sticker string `json:"sticker"`
		Title   string `json:"title"`
	}{
		Sticker: sticker,
		Title:   title,
	}

	_, err := b.sendMethodRequest(methodSetStickerSetTitle, params)
	return err
}

func (b *bot) SetStickerSetThumbnail(name string, user Userable, thumbnail File) error {
	params := struct {
		Name      string `json:"name"`
		UserID    any    `json:"user_id"`
		Thumbnail *File  `json:"thumbnail"`
	}{
		Name:      name,
		UserID:    user,
		Thumbnail: &thumbnail,
	}

	_, err := b.sendMethodRequest(methodSetStickerSetThumbnail, params)
	return err
}

func (b *bot) SetCustomEmojiStickerSetThumbnail(name string, CustomEmojiID ...string) error {
	params := struct {
		Name          string `json:"name"`
		CustomEmojiID string `json:"custom_emoji_id"`
	}{
		Name: name,
	}

	if len(CustomEmojiID) > 0 {
		params.CustomEmojiID = CustomEmojiID[0]
	}

	_, err := b.sendMethodRequest(methodSetCustomEmojiStickerSetThumbnail, params)
	return err
}

func (b *bot) DeleteStickerSet(name string) error {
	params := struct {
		Name string `json:"name"`
	}{
		Name: name,
	}

	_, err := b.sendMethodRequest(methodDeleteStickerSet, params)
	return err
}
