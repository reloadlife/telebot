package telebot

import "fmt"

type MessageOrigin struct {
	ID         int               `json:"message_id"`
	OriginType MessageOriginType `json:"type"`
	Date       int               `json:"date"`

	SenderUser      *User   `json:"sender_user,omitempty"`
	SenderUserName  *string `json:"sender_user_name,omitempty"`
	SenderChat      *Chat   `json:"sender_chat,omitempty"`
	AuthorSignature *string `json:"author_signature,omitempty"`
}

func (m *MessageOrigin) Type() string {
	if m.OriginType == "" {
		return "unknown"
	}
	return string(m.OriginType)
}

func (m *MessageOrigin) ReflectType() string {
	return fmt.Sprintf("%T", m)
}
