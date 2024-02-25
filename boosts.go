package telebot

import "fmt"

// BoostUpdated represents a boost added to a chat or changed.
// <a href="https://core.telegram.org/bots/api#chatboostupdated">/bots/api#chatboostupdated</a>
type BoostUpdated struct {
	// Chat is the chat which was boosted.
	Chat Chat `json:"chat"`

	// Boost is the information about the chat boost.
	Boost ChatBoost `json:"boost"`
}

// ReflectType returns the type of the struct.
func (b *BoostUpdated) ReflectType() string {
	return fmt.Sprintf("%T", b)
}

// Type returns the type of the struct.
func (b *BoostUpdated) Type() string {
	return "BoostUpdated"
}

// BoostRemoved represents a boost removed from a chat.
// <a href="https://core.telegram.org/bots/api#chatboostremoved">/bots/api#chatboostremoved</a>
type BoostRemoved struct {
	// Chat is the chat which was boosted.
	Chat Chat `json:"chat"`

	// BoostID is the unique identifier of the boost.
	BoostID string `json:"boost_id"`

	// RemoveDate is the point in time (Unix timestamp) when the boost was removed.
	RemoveDate int64 `json:"remove_date"`

	// Source is the source of the removed boost.
	Source ChatBoostSource `json:"source"`
}

// ReflectType returns the type of the struct.
func (c *BoostRemoved) ReflectType() string {
	return fmt.Sprintf("%T", c)
}

// Type returns the type of the struct.
func (c *BoostRemoved) Type() string {
	return "BoostRemoved"
}
