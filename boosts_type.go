package telebot

import "fmt"

// UserChatBoosts represents a list of boosts added to a chat by a user.
// <a href="https://core.telegram.org/bots/api#userchatboosts">/bots/api#userchatboosts</a>
type UserChatBoosts struct {
	// Boosts The list of boosts added to the chat by the user
	Boosts []ChatBoost `json:"boosts"`
}

// ReflectType returns the type of the struct
func (boost *UserChatBoosts) ReflectType() string { return fmt.Sprintf("%T", boost) }

// Type returns the type of the struct
func (boost *UserChatBoosts) Type() string { return "UserChatBoosts" }

// ChatBoost contains information about a chat boost.
// <a href="https://core.telegram.org/bots/api#chatboost">/bots/api#chatboost</a>
type ChatBoost struct {
	// ID Unique identifier of the boost
	ID string `json:"boost_id"`
	// AddedDate Point in time (Unix timestamp) when the chat was boosted
	// todo: #6: change to time.Time
	AddedDate int64 `json:"add_date"`
	// ExpiresAt Point in time (Unix timestamp) when the boost will automatically expire,
	// unless the booster's Telegram Premium subscription is prolonged
	// todo: #6: change to time.Time
	ExpiresAt int `json:"expires_at"`
	// Source Information about the source of the boost
	Source ChatBoostSource `json:"source"`
}

// ReflectType returns the type of the struct
func (cb *ChatBoost) ReflectType() string { return fmt.Sprintf("%T", cb) }

// Type returns the type of the struct
func (cb *ChatBoost) Type() string { return "ChatBoost" }

// ChatBoostSource describes the source of a chat boost. It can be one of
// the following: premium (ChatBoostSourcePremium), gift_code (ChatBoostSourceGiftCode), or giveaway (ChatBoostSourceGiveaway).
// <a href="https://core.telegram.org/bots/api#chatboostsource">/bots/api#chatboostsource</a>
type ChatBoostSource struct {
	// Source The source of the boost, one of "premium", "gift_code", or "giveaway"
	Source BoostSource `json:"source"`
	// User The user who added the boost to the chat.
	// It can be null if the source is "giveaway" and the winner has not claimed the boost yet
	User *User `json:"user,omitempty"`
	// GiveawayID Identifier of a message in the chat with the giveaway;
	// the message could have been deleted already. May be 0 if the message isn't sent yet.
	GiveawayID int `json:"giveaway_message_id,omitempty"`
	// IsUnclaimed True, if the giveaway was completed, but there was no user to win the prize
	IsUnclaimed bool `json:"is_unclaimed,omitempty"`
}

// ReflectType returns the type of the struct
func (bs *ChatBoostSource) ReflectType() string { return fmt.Sprintf("%T", bs) }

// Type returns the type of the struct
func (bs *ChatBoostSource) Type() string { return "ChatBoostSource" }
