package telebot

type MessageOriginType string

const (
	MessageOriginUser       MessageOriginType = "user"
	MessageOriginHiddenUser MessageOriginType = "hidden_user"
	MessageOriginChat       MessageOriginType = "chat"
	MessageOriginChannel    MessageOriginType = "channel"
)

type MessageOrigin struct {
	Type MessageOriginType `json:"type"`
	Date int               `json:"date"`

	SenderUser      *User   `json:"sender_user,omitempty"`
	SenderUserName  *string `json:"sender_user_name,omitempty"`
	SenderChat      *Chat   `json:"sender_chat,omitempty"`
	AuthorSignature *string `json:"author_signature,omitempty"`
	MessageId       int     `json:"message_id"`
}
