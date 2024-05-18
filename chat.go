package telebot

import (
	"fmt"
)

// Chat represents a chat.
type Chat struct {
	// ID is the unique identifier for this chat.
	// This number may have more than 32 significant bits, and some programming languages
	// may have difficulty/silent defects in interpreting it. It has at most 52 significant bits,
	// so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	ID int64 `json:"id"`

	// Type is the type of chat, which can be either "private", "group", "supergroup" or "channel".
	ChatType ChatType `json:"type"`

	// Title is the title for supergroups, channels, and group chats.
	Title *string `json:"title,omitempty"`

	// Username is the username for private chats, supergroups, and channels if available.
	Username *string `json:"username,omitempty"`

	// FirstName is the first name of the other party in a private chat.
	FirstName *string `json:"first_name,omitempty"`

	// LastName is the last name of the other party in a private chat.
	LastName *string `json:"last_name,omitempty"`

	// IsForum indicates whether the supergroup chat is a forum (has topics enabled).
	IsForum *bool `json:"is_forum,omitempty"`

	// Photo is the chat photo. Returned only in getChat.
	Photo *ChatPhoto `json:"photo,omitempty"`

	// ActiveUsernames is the list of all active chat usernames for private chats, supergroups, and channels.
	// Returned only in getChat.
	ActiveUsernames []string `json:"active_usernames,omitempty"`

	// AvailableReactions is the list of available reactions allowed in the chat.
	// If omitted, then all emoji reactions are allowed. Returned only in getChat.
	AvailableReactions []ReactionType `json:"available_reactions,omitempty"`

	// AccentColorID is the identifier of the accent color for the chat name and backgrounds of the chat photo,
	// reply header, and link preview. Returned only in getChat. Always returned in getChat.
	// todo: make Custom enum
	AccentColorID *int `json:"accent_color_id,omitempty"`

	// BackgroundCustomEmojiID is the custom emoji identifier of emoji chosen by the chat for the reply header and link preview background.
	// Returned only in getChat.
	BackgroundCustomEmojiID *string `json:"background_custom_emoji_id,omitempty"`

	// ProfileAccentColorID is the identifier of the accent color for the chat's profile background.
	// Returned only in getChat.
	ProfileAccentColorID *int `json:"profile_accent_color_id,omitempty"`

	// ProfileBackgroundCustomEmojiID is the custom emoji identifier of the emoji chosen by the chat for its profile background.
	// Returned only in getChat.
	ProfileBackgroundCustomEmojiID *string `json:"profile_background_custom_emoji_id,omitempty"`

	// EmojiStatusCustomEmojiID is the custom emoji identifier of the emoji status of the chat or the other party in a private chat.
	// Returned only in getChat.
	EmojiStatusCustomEmojiID *string `json:"emoji_status_custom_emoji_id,omitempty"`

	// EmojiStatusExpirationDate is the expiration date of the emoji status of the chat or the other party in a private chat, in Unix time, if any.
	// Returned only in getChat.
	EmojiStatusExpirationDate *int `json:"emoji_status_expiration_date,omitempty"`

	// Bio is the bio of the other party in a private chat. Returned only in getChat.
	Bio *string `json:"bio,omitempty"`

	// HasPrivateForwards indicates whether the privacy settings of the other party in the private chat allows
	// to use tg://user?id=<user_id> links only in chats with the user. Returned only in getChat.
	HasPrivateForwards *bool `json:"has_private_forwards,omitempty"`

	// HasRestrictedVoiceAndVideoMessages indicates whether the privacy settings of the other party restrict
	// sending voice and video note messages in the private chat. Returned only in getChat.
	HasRestrictedVoiceAndVideoMessages *bool `json:"has_restricted_voice_and_video_messages,omitempty"`

	// JoinToSendMessages indicates whether users need to join the supergroup before they can send messages.
	// Returned only in getChat.
	JoinToSendMessages *bool `json:"join_to_send_messages,omitempty"`

	// JoinByRequest indicates whether all users directly joining the supergroup need to be approved by supergroup administrators.
	// Returned only in getChat.
	JoinByRequest *bool `json:"join_by_request,omitempty"`

	// Description is the description for groups, supergroups, and channel chats.
	// Returned only in getChat.
	Description *string `json:"description,omitempty"`

	// InviteLink is the primary invite link for groups, supergroups, and channel chats.
	// Returned only in getChat.
	InviteLink *string `json:"invite_link,omitempty"`

	// PinnedMessage is the most recent pinned message (by sending date).
	// Returned only in getChat.
	PinnedMessage *AccessibleMessage `json:"pinned_message,omitempty"`

	// Permissions is the default chat member permissions for groups and supergroups.
	// Returned only in getChat.
	Permissions *ChatPermissions `json:"permissions,omitempty"`

	// SlowModeDelay is the minimum allowed delay between consecutive messages sent by each unprivileged user; in seconds.
	// Returned only in getChat.
	SlowModeDelay *int `json:"slow_mode_delay,omitempty"`

	// MessageAutoDeleteTime is the time after which all messages sent to the chat will be automatically deleted; in seconds.
	// Returned only in getChat.
	MessageAutoDeleteTime *int `json:"message_auto_delete_time,omitempty"`

	// HasAggressiveAntiSpamEnabled indicates whether aggressive anti-spam checks are enabled in the supergroup.
	// The field is only available to chat administrators. Returned only in getChat.
	HasAggressiveAntiSpamEnabled *bool `json:"has_aggressive_anti_spam_enabled,omitempty"`

	// HasHiddenMembers indicates whether non-administrators can only get the list of bots and administrators in the chat.
	// Returned only in getChat.
	HasHiddenMembers *bool `json:"has_hidden_members,omitempty"`

	// HasProtectedContent indicates whether messages from the chat can't be forwarded to other chats.
	// Returned only in getChat.
	HasProtectedContent *bool `json:"has_protected_content,omitempty"`

	// HasVisibleHistory indicates whether new chat members will have access to old messages;
	// available only to chat administrators. Returned only in getChat.
	HasVisibleHistory *bool `json:"has_visible_history,omitempty"`

	// StickerSetName is the name of the group sticker set. Returned only in getChat.
	StickerSetName *string `json:"sticker_set_name,omitempty"`

	// CanSetStickerSet indicates whether the bot can change the group sticker set. Returned only in getChat.
	CanSetStickerSet *bool `json:"can_set_sticker_set,omitempty"`

	// LinkedChatID is the unique identifier for the linked chat, i.e., the discussion group identifier for a channel and vice versa;
	// for supergroups and channel chats. This identifier may be greater than 32 bits, and some programming languages
	// may have difficulty/silent defects in interpreting it. But it is smaller than 52 bits, so a signed 64-bit integer or double-precision float type are safe for storing this identifier.
	// Returned only in getChat.
	LinkedChatID *int64 `json:"linked_chat_id,omitempty"`

	// Location is the location to which the supergroup is connected.
	// Returned only in getChat.
	Location *ChatLocation `json:"location,omitempty"`

	// UnRestrictBoostCount For supergroups, the minimum number of boosts that a
	// non-administrator user needs to add in order to ignore slow mode and chat
	// permissions. Returned only in getChat.
	UnRestrictBoostCount *int `json:"unrestrict_boost_count,omitempty"`

	// CustomEmojiStickerSetName For supergroups,
	// the name of the group's custom emoji sticker set.
	// Custom emoji from this set can be used by all users and bots in the group.
	// Returned only in getChat.
	CustomEmojiStickerSetName *string `json:"custom_emoji_sticker_set_name,omitempty"`

	PersonalChat     *Chat                 `json:"personal_chat,omitempty"`
	BusinessIntro    *BusinessIntro        `json:"business_intro,omitempty"`
	BusinessLocation *BusinessLocation     `json:"business_location,omitempty"`
	BusinessHours    *BusinessOpeningHours `json:"business_opening_hours,omitempty"`

	BirthDate *BirthDate `json:"birthdate,omitempty"`
}

func (c *Chat) ReflectType() string {
	return fmt.Sprintf("%T", c)
}

func (c *Chat) Type() string {
	if c.ChatType == "" {
		return Unknown
	}
	return string(c.ChatType)
}
