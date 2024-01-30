package telebot

import (
	"encoding/json"
	"fmt"
)

// Update represents an incoming update in a Telegram bot.
// At most, one of the optional parameters can be present in any given update.
type Update struct {
	ID int `json:"update_id"`

	// New incoming message of any kind - text, photo, sticker, etc.
	Message *Message `json:"message,omitempty"`

	// New version of a message that is known to the bot and was edited.
	EditedMessage *Message `json:"edited_message,omitempty"`

	// New incoming channel post of any kind - text, photo, sticker, etc.
	ChannelPost *Message `json:"channel_post,omitempty"`

	// New version of a channel post that is known to the bot and was edited.
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`

	// A reaction to a message was changed by a user.
	MessageReaction *MessageReaction `json:"message_reaction,omitempty"`

	// Reactions to a message with anonymous reactions were changed.
	MessageReactionCount *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`

	// New incoming inline query.
	Query *Query `json:"inline_query,omitempty"`

	// The result of an inline query that was chosen by a user and sent to their chat partner.
	InlineResult *InlineResult `json:"chosen_inline_result,omitempty"`

	// New incoming callback query.
	Callback *Callback `json:"callback_query,omitempty"`

	// New incoming shipping query. Only for invoices with flexible price.
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`

	// New incoming pre-checkout query. Contains full information about a checkout.
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`

	// New poll state. Bots receive only updates about manually stopped polls and polls, which are sent by the bot.
	Poll *Poll `json:"poll,omitempty"`

	// A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`

	// The bot's chat member status was updated in a chat. For private chats, this update is received only when the bot is blocked or unblocked by the user.
	MyChatMember *ChatMemberUpdate `json:"my_chat_member,omitempty"`

	// A chat member's status was updated in a chat. The bot must be an administrator in the chat and must explicitly specify "chat_member" in the list of allowed_updates to receive these updates.
	ChatMember *ChatMemberUpdate `json:"chat_member,omitempty"`

	// A request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request,omitempty"`

	// A chat boost was added or changed. The bot must be an administrator in the chat to receive these updates.
	ChatBoost *ChatBoost `json:"chat_boost,omitempty"`

	// A boost was removed from a chat. The bot must be an administrator in the chat to receive these updates.
	ChatBoostRemoved *ChatBoostRemoved `json:"removed_chat_boost,omitempty"`
}

// MarshalJSON to be JSON serializable, but only include non-empty fields.
func (u *Update) MarshalJSON() ([]byte, error) {
	return json.Marshal(u)
}

// UnmarshalJSON to be JSON unserializable
func (u *Update) UnmarshalJSON(b []byte) error {
	return json.Unmarshal(b, u)
}

func (u *Update) String() string {
	indented, _ := json.MarshalIndent(u, "", "  ")
	return fmt.Sprintf("Update{ID: %d, Type: %s}\n%s\n", u.ID, u.Type().String(), indented)
}

// UpdateType represents the type of update.
type UpdateType int

// Update types.
const (
	UpdateTypeMessage UpdateType = iota
	UpdateTypeEditedMessage
	UpdateTypeChannelPost
	UpdateTypeEditedChannelPost
	UpdateTypeMessageReaction
	UpdateTypeMessageReactionCount
	UpdateTypeInlineQuery
	UpdateTypeChosenInlineResult
	UpdateTypeCallbackQuery
	UpdateTypeShippingQuery
	UpdateTypePreCheckoutQuery
	UpdateTypePoll
	UpdateTypePollAnswer
	UpdateTypeMyChatMember
	UpdateTypeChatMember
	UpdateTypeChatJoinRequest
	UpdateTypeChatBoost
	UpdateTypeChatBoostRemoved
	UpdateTypeUnknown
)

// String returns the string representation of the UpdateType.
func (ut UpdateType) String() string {
	switch ut {
	case UpdateTypeMessage:
		return "message"
	case UpdateTypeEditedMessage:
		return "edited_message"
	case UpdateTypeChannelPost:
		return "channel_post"
	case UpdateTypeEditedChannelPost:
		return "edited_channel_post"
	case UpdateTypeMessageReaction:
		return "message_reaction"
	case UpdateTypeMessageReactionCount:
		return "message_reaction_count"
	case UpdateTypeInlineQuery:
		return "inline_query"
	case UpdateTypeChosenInlineResult:
		return "chosen_inline_result"
	case UpdateTypeCallbackQuery:
		return "callback_query"
	case UpdateTypeShippingQuery:
		return "shipping_query"
	case UpdateTypePreCheckoutQuery:
		return "pre_checkout_query"
	case UpdateTypePoll:
		return "poll"
	case UpdateTypePollAnswer:
		return "poll_answer"
	case UpdateTypeMyChatMember:
		return "my_chat_member"
	case UpdateTypeChatMember:
		return "chat_member"
	case UpdateTypeChatJoinRequest:
		return "chat_join_request"
	case UpdateTypeChatBoost:
		return "chat_boost"
	case UpdateTypeChatBoostRemoved:
		return "removed_chat_boost"
	default:
		return "unknown"
	}
}

// Type returns the type of the Update.
func (u *Update) Type() UpdateType {
	switch true {
	case u.Message != nil:
		return UpdateTypeMessage

	case u.EditedMessage != nil:
		return UpdateTypeEditedMessage

	case u.ChannelPost != nil:
		return UpdateTypeChannelPost

	case u.EditedChannelPost != nil:
		return UpdateTypeEditedChannelPost

	case u.MessageReaction != nil:
		return UpdateTypeMessageReaction

	case u.MessageReactionCount != nil:
		return UpdateTypeMessageReactionCount

	case u.Query != nil:
		return UpdateTypeInlineQuery

	case u.InlineResult != nil:
		return UpdateTypeChosenInlineResult

	case u.Callback != nil:
		return UpdateTypeCallbackQuery

	case u.ShippingQuery != nil:
		return UpdateTypeShippingQuery

	case u.PreCheckoutQuery != nil:
		return UpdateTypePreCheckoutQuery

	case u.Poll != nil:
		return UpdateTypePoll

	case u.PollAnswer != nil:
		return UpdateTypePollAnswer

	case u.MyChatMember != nil:
		return UpdateTypeMyChatMember

	case u.ChatMember != nil:
		return UpdateTypeChatMember

	case u.ChatJoinRequest != nil:
		return UpdateTypeChatJoinRequest

	case u.ChatBoost != nil:
		return UpdateTypeChatBoost

	case u.ChatBoostRemoved != nil:
		return UpdateTypeChatBoostRemoved

	default:
		return UpdateTypeUnknown

	}
}
