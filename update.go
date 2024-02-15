package telebot

import (
	"encoding/json"
	"fmt"
)

type Updates []Update

// Update represents an incoming update.
type Update struct {
	// ID is the update's a unique identifier.
	// Update identifiers start from a certain positive number and increase sequentially.
	// This identifier becomes handy if you're using webhooks, allowing you to ignore repeated updates or restore the correct update sequence if they get out of order.
	// If there are no new updates for at least a week, then the identifier of the next update will be chosen randomly instead of sequentially.
	ID int `json:"update_id"`

	// Message is the new incoming message of any kind – text, photo, sticker, etc. (optional).
	Message *Message `json:"message,omitempty"`

	// EditedMessage is the new version of a message that is known to the bot and was edited (optional).
	EditedMessage *Message `json:"edited_message,omitempty"`

	// ChannelPost is the new incoming channel post of any kind – text, photo, sticker, etc. (optional).
	ChannelPost *Message `json:"channel_post,omitempty"`

	// EditedChannelPost is the new version of a channel post that is known to the bot and was edited (optional).
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`

	// Reaction is the reaction to a message changed by a user (optional).
	Reaction *MessageReaction `json:"message_reaction,omitempty"`

	// ReactionCount is the reactions to a message with anonymous reactions that were changed (optional).
	ReactionCount *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`

	// InlineQuery is the new incoming inline query (optional).
	Query *InlineQuery `json:"inline_query,omitempty"`

	// InlineResult is the result of an inline query that was chosen by a user and sent to their chat partner (optional).
	InlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`

	// CallbackQuery is the new incoming callback query (optional).
	CallbackQuery *Callback `json:"callback_query,omitempty"`

	// ShippingQuery is the new incoming shipping query (optional, only for invoices with flexible price).
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`

	// PreCheckoutQuery is the new incoming pre-checkout query. Contains full information about checkout (optional).
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`

	// Poll is the new poll state. Bots receive only updates about manually stopped polls and polls sent by the bot (optional).
	Poll *Poll `json:"poll,omitempty"`

	// PollAnswer is the user changed their answer in a non-anonymous poll. Bots receive new votes only in polls sent by the bot itself (optional).
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`

	// MyChatMember is the bot's chat member status updated in a chat (optional, for private chats, this update is received only when the bot is blocked or unblocked by the user).
	MyChatMember *ChatMemberUpdated `json:"my_chat_member,omitempty"`

	// ChatMember is a chat member's status updated in a chat (optional, the bot must be an administrator in the chat and must explicitly specify "chat_member" in the list of allowed_updates to receive these updates).
	ChatMember *ChatMemberUpdated `json:"chat_member,omitempty"`

	// ChatJoinRequest is a request to join the chat that has been sent (optional, the bot must have the can_invite_users administrator right in the chat to receive these updates).
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request,omitempty"`

	// ChatBoost is a chat boost added or changed (optional, the bot must be an administrator in the chat to receive these updates).
	ChatBoost *ChatBoostUpdated `json:"chat_boost,omitempty"`

	// RemovedChatBoost is a boost removed from a chat (optional, the bot must be an administrator in the chat to receive these updates).
	RemovedChatBoost *ChatBoostRemoved `json:"removed_chat_boost,omitempty"`
}

func (u *Update) MarshalJSON() ([]byte, error) {
	return json.Marshal(u)
}

func (u *Update) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, u)
}

func (u *Update) String() string {
	indented, _ := json.MarshalIndent(u, "", "  ")
	return fmt.Sprintf("Update{ID: %d, Type: %s}\n%s\n", u.ID, u.Type().String(), indented)
}

func (u *Update) Verify() error {
	if u.ID == 0 {
		return fmt.Errorf("telebot: Update ID is empty")
	}

	return nil
}
