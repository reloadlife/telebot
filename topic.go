package telebot

type ForumTopic struct {
	ID                int64       `json:"message_thread_id"`
	Name              string      `json:"name"`
	IconColor         IconColor   `json:"icon_color"`
	IconCustomEmojiID CustomEmoji `json:"icon_custom_emoji_id,omitempty"`
}
