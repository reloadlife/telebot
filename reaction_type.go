package telebot

import "fmt"

// setMessageReactionRequest is the request body for the methodSetMessageReaction (Bot.SetMessageReaction) method.
type setMessageReactionRequest struct {
	ChatID    any            `json:"chat_id"`
	MessageID int            `json:"message_id"`
	Reaction  []ReactionType `json:"reaction,omitempty"`
	IsBig     bool           `json:"is_big,omitempty"`
}

// ReactionType represents the type of a reaction.
type ReactionType struct {
	ReactionType  ReactionTypeType `json:"type"`
	Emoji         Emoji            `json:"emoji,omitempty"`
	CustomEmojiID string           `json:"custom_emoji_id,omitempty"`
}

func (c *ReactionType) Type() string {
	if c.ReactionType == "" {
		return "unknown"
	}
	return string(c.ReactionType)
}
func (c *ReactionType) ReflectType() string { return fmt.Sprintf("%T", c) }

// ReactionCount represents the count of a specific reaction on a message.
type ReactionCount struct {
	// ReactionType is the type of reaction.
	ReactionType ReactionType `json:"type"`

	// Count is the count of reactions of the specified type.
	Count int `json:"count"`
}

func (c *ReactionCount) Type() string {
	if c.ReactionType.ReactionType == "" {
		return "unknown"
	}
	return string(c.ReactionType.ReactionType)
}
func (c *ReactionCount) ReflectType() string { return fmt.Sprintf("%T", c) }
