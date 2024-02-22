package telebot

import "fmt"

type ForumTopic struct {
	ID                int64       `json:"message_thread_id"`
	Name              string      `json:"name"`
	IconColor         IconColor   `json:"icon_color"`
	IconCustomEmojiID CustomEmoji `json:"icon_custom_emoji_id,omitempty"`
}

func (c *ForumTopic) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *ForumTopic) Type() string        { return "ForumTopic" }
