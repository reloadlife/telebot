package telebot

import (
	"encoding/json"
	"fmt"
)

type ForumTopicCreated struct {
	Name              string      `json:"name"`
	IconColor         IconColor   `json:"icon_color"`
	IconCustomEmojiID CustomEmoji `json:"icon_custom_emoji_id,omitempty"`
}

type ForumTopicEdited struct {
	Name          string      `json:"name"`
	CustomEmojiID CustomEmoji `json:"custom_emoji_id,omitempty"`
}

type ForumTopicClosed struct {
}

type ForumTopicReopened struct {
}

type ForumTopic struct {
	ID                int64       `json:"message_thread_id"`
	Name              string      `json:"name"`
	IconColor         IconColor   `json:"icon_color"`
	IconCustomEmojiID CustomEmoji `json:"icon_custom_emoji_id,omitempty"`
}

type GeneralForumTopicHidden struct {
}

type GeneralForumTopicUnhidden struct {
}

func (b *OldBot) CreateTopic(where Recipient, name string, opts ...interface{}) (*ForumTopic, error) {
	params := map[string]string{
		"chat_id": where.Recipient(),
		"name":    name,
	}

	options := extractOptions(opts)
	if options.IconCustomEmojiID != "" {
		params["icon_custom_emoji_id"] = string(options.IconCustomEmojiID)
	}

	if options.IconColor != 0 {
		params["icon_color"] = fmt.Sprintf("%d", options.IconColor)
	}

	data, err := b.Raw("createForumTopic", params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result ForumTopic
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return &resp.Result, nil
}

func (b *OldBot) EditTopic(where Recipient, topicID int64, name string, opts ...interface{}) error {
	params := map[string]string{
		"chat_id":           where.Recipient(),
		"message_thread_id": fmt.Sprintf("%d", topicID),
		"name":              name,
	}

	options := extractOptions(opts)
	if options.IconCustomEmojiID != "" {
		params["icon_custom_emoji_id"] = string(options.IconCustomEmojiID)
	}
	_, err := b.Raw("editForumTopic", params)
	return err
}

func (b *OldBot) CloseTopic(where Recipient, topicID int64) error {
	params := map[string]string{
		"chat_id":           where.Recipient(),
		"message_thread_id": fmt.Sprintf("%d", topicID),
	}
	_, err := b.Raw("closeForumTopic", params)
	return err
}

func (b *OldBot) ReopenTopic(where Recipient, topicID int64) error {
	params := map[string]string{
		"chat_id":           where.Recipient(),
		"message_thread_id": fmt.Sprintf("%d", topicID),
	}
	_, err := b.Raw("reopenForumTopic", params)
	return err
}

func (b *OldBot) DeleteTopic(where Recipient, topicID int64) error {
	params := map[string]string{
		"chat_id":           where.Recipient(),
		"message_thread_id": fmt.Sprintf("%d", topicID),
	}
	_, err := b.Raw("deleteForumTopic", params)
	return err
}

func (b *OldBot) UnpinAllForumTopicMessages(where Recipient, topicID int64) error {
	params := map[string]string{
		"chat_id":           where.Recipient(),
		"message_thread_id": fmt.Sprintf("%d", topicID),
	}
	_, err := b.Raw("unpinAllForumTopicMessages", params)
	return err
}

func (b *OldBot) GetForumTopicIconStickers() ([]Sticker, error) {
	data, err := b.Raw("getForumTopicIconStickers", nil)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result []Sticker
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, err
	}
	return resp.Result, nil
}
