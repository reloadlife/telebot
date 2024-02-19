package telebot

import (
	"encoding/json"
	"fmt"
)

// MessageReaction represents a change of a reaction on a message performed by a user.
type MessageReaction struct {
	// ID is the unique identifier of the message inside the chat.
	ID int `json:"message_id"`

	// Chat is the chat containing the message the user reacted to.
	Chat Chat `json:"chat"`

	// User is the user that changed the reaction. It is optional if the user isn't anonymous.
	User *User `json:"user,omitempty"`

	// ActorChat is the chat on behalf of which the reaction was changed, if the user is anonymous.
	ActorChat *Chat `json:"actor_chat,omitempty"`

	// Date is the date of the change in Unix time.
	Date int64 `json:"date"`

	// OldReaction is the previous list of reaction types that were set by the user.
	OldReaction []ReactionType `json:"old_reaction"`

	// NewReaction is the new list of reaction types that have been set by the user.
	NewReaction []ReactionType `json:"new_reaction"`
}

// MarshalJSON to be JSON serializable, but only include non-empty fields.
func (r *MessageReaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(r)
}

// UnmarshalJSON to be JSON unserializable
func (r *MessageReaction) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, r)
}

func (r *MessageReaction) String() string {
	indented, _ := json.MarshalIndent(r, "", "  ")
	return fmt.Sprintf("MessageReaction{ID: %d}\n%s\n", r.ID, indented)
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

// MarshalJSON to be JSON serializable, but only include non-empty fields.
func (r *MessageReactionCountUpdated) MarshalJSON() ([]byte, error) {
	return json.Marshal(r)
}

// UnmarshalJSON to be JSON unserializable
func (r *MessageReactionCountUpdated) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, r)
}

func (r *MessageReactionCountUpdated) String() string {
	indented, _ := json.MarshalIndent(r, "", "  ")
	return fmt.Sprintf("MessageReactionCountUpdated{ID: %d}\n%s\n", r.ID, indented)
}
