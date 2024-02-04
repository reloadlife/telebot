package telebot

type createForumTopicRequest struct {
	ChatID            any         `json:"chat_id"`
	Name              string      `json:"name"`
	IconColor         IconColor   `json:"icon_color,omitempty"`
	IconCustomEmojiID CustomEmoji `json:"icon_custom_emoji_id,omitempty"`
}

type editForumTopicRequest struct {
	ChatID            any         `json:"chat_id"`
	MessageThreadID   int64       `json:"message_thread_id,omitempty"`
	Name              string      `json:"name,omitempty"`
	IconCustomEmojiID CustomEmoji `json:"icon_custom_emoji_id,omitempty"`
}

type forumTopicRequest struct {
	ChatID          any   `json:"chat_id"`
	MessageThreadID int64 `json:"message_thread_id,omitempty"`
}
