package telebot

import "encoding/json"

func (b *bot) GetForumTopicIconStickers() ([]Sticker, error) {
	data, err := b.sendMethodRequest(methodGetForumTopicIconStickers, nil)
	if err != nil {
		return nil, err
	}

	var result struct {
		Result []Sticker `json:"result"`
	}
	err = json.Unmarshal(data, &result)
	return result.Result, err
}

func (b *bot) CreateForumTopic(chatID Recipient, name string, options ...any) (*ForumTopic, error) {
	params := createForumTopicRequest{
		ChatID: chatID,
		Name:   name,
	}

	for _, option := range options {
		switch v := option.(type) {
		case IconColor:
			params.IconColor = v
		case CustomEmoji:
			params.IconCustomEmojiID = v
		default:
			panic("telebot: unsupported flag-option")
		}
	}

	var result struct {
		Result ForumTopic `json:"result"`
	}
	data, err := b.sendMethodRequest(methodCreateForumTopic, params)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &result)
	return &result.Result, err
}

func (b *bot) EditForumTopic(chatID Recipient, topicID int64, options ...any) error {
	params := editForumTopicRequest{
		ChatID:          chatID,
		MessageThreadID: topicID,
	}

	for _, option := range options {
		switch v := option.(type) {
		case string:
			params.Name = v
		case CustomEmoji:
			params.IconCustomEmojiID = v
		default:
			panic("telebot: unsupported flag-option")
		}
	}

	_, err := b.sendMethodRequest(methodEditForumTopic, params)
	return err
}

func (b *bot) CloseForumTopic(chatID Recipient, topicID int64) error {
	params := forumTopicRequest{
		ChatID:          chatID,
		MessageThreadID: topicID,
	}

	_, err := b.sendMethodRequest(methodCloseForumTopic, params)
	return err
}

func (b *bot) ReopenForumTopic(chatID Recipient, topicID int64) error {
	params := forumTopicRequest{
		ChatID:          chatID,
		MessageThreadID: topicID,
	}

	_, err := b.sendMethodRequest(methodReopenForumTopic, params)
	return err
}

func (b *bot) DeleteForumTopic(chatID Recipient, topicID int64) error {
	params := forumTopicRequest{
		ChatID:          chatID,
		MessageThreadID: topicID,
	}

	_, err := b.sendMethodRequest(methodDeleteForumTopic, params)
	return err
}

func (b *bot) UnpinAllForumTopicMessages(chatID Recipient, topicID int64) error {
	params := forumTopicRequest{
		ChatID:          chatID,
		MessageThreadID: topicID,
	}

	_, err := b.sendMethodRequest(methodUnpinAllForumTopicMessages, params)
	return err
}

func (b *bot) EditGeneralForumTopic(chatID Recipient, name string) error {
	params := editForumTopicRequest{
		ChatID: chatID,
		Name:   name,
	}

	_, err := b.sendMethodRequest(methodEditGeneralForumTopic, params)
	return err
}

func (b *bot) CloseGeneralForumTopic(chatID Recipient) error {
	params := forumTopicRequest{
		ChatID: chatID,
	}

	_, err := b.sendMethodRequest(methodCloseGeneralForumTopic, params)
	return err
}

func (b *bot) ReopenGeneralForumTopic(chatID Recipient) error {
	params := forumTopicRequest{
		ChatID: chatID,
	}

	_, err := b.sendMethodRequest(methodReopenGeneralForumTopic, params)
	return err
}

func (b *bot) HideGeneralForumTopic(chatID Recipient) error {
	params := forumTopicRequest{
		ChatID: chatID,
	}

	_, err := b.sendMethodRequest(methodHideGeneralForumTopic, params)
	return err
}

func (b *bot) UnhideGeneralForumTopic(chatID Recipient) error {
	params := forumTopicRequest{
		ChatID: chatID,
	}

	_, err := b.sendMethodRequest(methodUnhideGeneralForumTopic, params)
	return err
}

func (b *bot) UnpinAllGeneralForumTopicMessages(chatID Recipient) error {
	params := forumTopicRequest{
		ChatID: chatID,
	}

	_, err := b.sendMethodRequest(methodUnpinAllGeneralForumTopicMessages, params)
	return err
}
