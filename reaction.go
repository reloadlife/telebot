package telebot

import (
	"fmt"
)

// MessageReaction represents a change of a reaction on a message performed by a user.
type MessageReaction struct {
	// ID is the unique identifier of the message inside the chat.
	ID int `json:"message_id" verify:"required"`

	// Chat is the chat containing the message the user reacted to.
	Chat Chat `json:"chat" verify:"required"`

	// User is the user that changed the reaction. It is optional if the user isn't anonymous.
	User *User `json:"user,omitempty"`

	// ActorChat is the chat on behalf of which the reaction was changed, if the user is anonymous.
	ActorChat *Chat `json:"actor_chat,omitempty"`

	// Date is the date of the change in Unix time.
	Date int64 `json:"date" verify:"required"`

	// OldReaction is the previous list of reaction types that were set by the user.
	OldReaction []ReactionType `json:"old_reaction" verify:"required"`

	// NewReaction is the new list of reaction types that have been set by the user.
	NewReaction []ReactionType `json:"new_reaction" verify:"required"`
}

func (r *MessageReaction) ReflectType() string {
	return fmt.Sprintf("%T", r)
}

func (r *MessageReaction) Type() string {
	return "MessageReaction"
}

type MessageReactionCountUpdated struct {
	// Chat is the chat containing the message.
	Chat Chat `json:"chat"`

	// MessageID is the unique message identifier inside the chat.
	ID int `json:"message_id"`

	// Date is the date of the change in Unix time.
	Date int64 `json:"date"`

	// Reactions is the list of reactions that are present on the message.
	Reactions []ReactionCount `json:"reactions"`
}

func (r *MessageReactionCountUpdated) ReflectType() string {
	return fmt.Sprintf("%T", r)
}

func (r *MessageReactionCountUpdated) Type() string {
	return "MessageReactionCountUpdated"
}
