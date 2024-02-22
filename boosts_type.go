package telebot

import "fmt"

type getUserChatBoosts struct {
	ChatId any `json:"chat_id"`
	UserId any `json:"user_id"`
}

type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"`
}

func (c *UserChatBoosts) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *UserChatBoosts) Type() string        { return "UserChatBoosts" }

type ChatBoost struct {
	BoostID   string          `json:"boost_id"`
	Date      int             `json:"date"`
	ExpiresAt int             `json:"expires_at"`
	Source    ChatBoostSource `json:"source"`
}

func (c *ChatBoost) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *ChatBoost) Type() string        { return "ChatBoost" }

type ChatBoostSource struct {
	Source            BoostSource `json:"source"`
	User              *User       `json:"user,omitempty"`
	GiveawayMessageID int         `json:"giveaway_message_id,omitempty"`
	IsUnclaimed       bool        `json:"is_unclaimed,omitempty"`
}

func (c *ChatBoostSource) ReflectType() string { return fmt.Sprintf("%T", c) }
func (c *ChatBoostSource) Type() string        { return "ChatBoostSource" }
