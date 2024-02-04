package telebot

type IconColor int

const (
	IconColorBlue      IconColor = 7322096
	IconColorOrange    IconColor = 16766590
	IconColorPurple    IconColor = 13338331
	IconColorGreen     IconColor = 9367192
	IconColorRed       IconColor = 16749490
	IconColorOrangeRed IconColor = 16478047
)

type ForumTopic struct {
	ID                int64       `json:"message_thread_id"`
	Name              string      `json:"name"`
	IconColor         IconColor   `json:"icon_color"`
	IconCustomEmojiID CustomEmoji `json:"icon_custom_emoji_id,omitempty"`
}
