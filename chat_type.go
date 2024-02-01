package telebot

import (
	"fmt"
)

// ChatPermissions describes actions that a non-administrator user is allowed to take in a chat.
type ChatPermissions struct {
	// CanSendMessages indicates whether the user is allowed to send text messages, contacts, giveaways, giveaway winners, invoices, locations, and venues.
	CanSendMessages *bool `json:"can_send_messages,omitempty"`

	// CanSendAudios indicates whether the user is allowed to send audios.
	CanSendAudios *bool `json:"can_send_audios,omitempty"`

	// CanSendDocuments indicates whether the user is allowed to send documents.
	CanSendDocuments *bool `json:"can_send_documents,omitempty"`

	// CanSendPhotos indicates whether the user is allowed to send photos.
	CanSendPhotos *bool `json:"can_send_photos,omitempty"`

	// CanSendVideos indicates whether the user is allowed to send videos.
	CanSendVideos *bool `json:"can_send_videos,omitempty"`

	// CanSendVideoNotes indicates whether the user is allowed to send video notes.
	CanSendVideoNotes *bool `json:"can_send_video_notes,omitempty"`

	// CanSendVoiceNotes indicates whether the user is allowed to send voice notes.
	CanSendVoiceNotes *bool `json:"can_send_voice_notes,omitempty"`

	// CanSendPolls indicates whether the user is allowed to send polls.
	CanSendPolls *bool `json:"can_send_polls,omitempty"`

	// CanSendOtherMessages indicates whether the user is allowed to send animations, games, stickers, and use inline bots.
	CanSendOtherMessages *bool `json:"can_send_other_messages,omitempty"`

	// CanAddWebPagePreviews indicates whether the user is allowed to add web page previews to their messages.
	CanAddWebPagePreviews *bool `json:"can_add_web_page_previews,omitempty"`

	// CanChangeInfo indicates whether the user is allowed to change the chat title, photo, and other settings. Ignored in public supergroups.
	CanChangeInfo *bool `json:"can_change_info,omitempty"`

	// CanInviteUsers indicates whether the user is allowed to invite new users to the chat.
	CanInviteUsers *bool `json:"can_invite_users,omitempty"`

	// CanPinMessages indicates whether the user is allowed to pin messages. Ignored in public supergroups.
	CanPinMessages *bool `json:"can_pin_messages,omitempty"`

	// CanManageTopics indicates whether the user is allowed to create forum topics. If omitted, defaults to the value of can_pin_messages.
	CanManageTopics *bool `json:"can_manage_topics,omitempty"`
}

// BanChatMemberRequest represents the parameters for the banChatMember method.
type banChatMemberRequest struct {
	// ChatID is the unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername).
	ChatID interface{} `json:"chat_id"`

	// UserID is the unique identifier of the target user.
	UserID int64 `json:"user_id"`

	// UntilDate is the date when the user will be unbanned; Unix time.
	// If the user is banned for more than 366 days or less than 30 seconds from the current time, they are considered to be banned forever.
	// Applied for supergroups and channels only.
	UntilDate *int64 `json:"until_date,omitempty"`

	// RevokeMessages indicates whether to delete all messages from the chat for the user that is being removed.
	// If false, the user will be able to see messages in the group that were sent before the user was removed.
	// Always true for supergroups and channels.
	RevokeMessages *bool `json:"revoke_messages,omitempty"`
}

// UnbanChatMemberRequest represents the parameters for the unbanChatMember method.
type unbanChatMemberRequest struct {
	// ChatID is the unique identifier for the target group or username of the target supergroup or channel (in the format @channelusername).
	ChatID interface{} `json:"chat_id"`

	// UserID is the unique identifier of the target user.
	UserID int64 `json:"user_id"`

	// OnlyIfBanned indicates whether to do nothing if the user is not banned.
	OnlyIfBanned bool `json:"only_if_banned,omitempty"`
}

// RestrictChatMemberRequest represents the parameters for the restrictChatMember method.
type restrictChatMemberRequest struct {
	// ChatID is the unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatID interface{} `json:"chat_id"`

	// UserID is the unique identifier of the target user.
	UserID int64 `json:"user_id"`

	// Permissions is a JSON-serialized object for new user permissions.
	Permissions ChatPermissions `json:"permissions"`

	// UseIndependentChatPermissions indicates whether chat permissions are set independently.
	// Pass true if chat permissions are set independently.
	// Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages,
	// can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions;
	// the can_send_polls permission will imply the can_send_messages permission.
	UseIndependentChatPermissions bool `json:"use_independent_chat_permissions,omitempty"`

	// UntilDate is the date when restrictions will be lifted for the user; Unix time.
	// If the user is restricted for more than 366 days or less than 30 seconds from the current time, they are considered to be restricted forever.
	UntilDate int64 `json:"until_date,omitempty"`
}

// PromoteChatMemberRequest represents the parameters for the promoteChatMember method.
type promoteChatMemberRequest struct {
	// ChatID is the unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID interface{} `json:"chat_id"`

	// UserID is the unique identifier of the target user.
	UserID int64 `json:"user_id"`

	// IsAnonymous indicates whether the administrator's presence in the chat is hidden.
	IsAnonymous *bool `json:"is_anonymous,omitempty"`

	// CanManageChat indicates whether the administrator can access the chat event log, boost list in channels,
	// see channel members, report spam messages, see anonymous administrators in supergroups and ignore slow mode.
	// Implied by any other administrator privilege.
	CanManageChat *bool `json:"can_manage_chat,omitempty"`

	// CanDeleteMessages indicates whether the administrator can delete messages of other users.
	CanDeleteMessages *bool `json:"can_delete_messages,omitempty"`

	// CanManageVideoChats indicates whether the administrator can manage video chats.
	CanManageVideoChats *bool `json:"can_manage_video_chats,omitempty"`

	// CanRestrictMembers indicates whether the administrator can restrict, ban or unban chat members, or access supergroup statistics.
	CanRestrictMembers *bool `json:"can_restrict_members,omitempty"`

	// CanPromoteMembers indicates whether the administrator can add new administrators with a subset of their own privileges
	// or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by him).
	CanPromoteMembers *bool `json:"can_promote_members,omitempty"`

	// CanChangeInfo indicates whether the administrator can change chat title, photo and other settings.
	CanChangeInfo *bool `json:"can_change_info,omitempty"`

	// CanInviteUsers indicates whether the administrator can invite new users to the chat.
	CanInviteUsers *bool `json:"can_invite_users,omitempty"`

	// CanPostMessages indicates whether the administrator can post messages in the channel, or access channel statistics (channels only).
	CanPostMessages *bool `json:"can_post_messages,omitempty"`

	// CanEditMessages indicates whether the administrator can edit messages of other users and can pin messages (channels only).
	CanEditMessages *bool `json:"can_edit_messages,omitempty"`

	// CanPinMessages indicates whether the administrator can pin messages (supergroups only).
	CanPinMessages *bool `json:"can_pin_messages,omitempty"`

	// CanPostStories indicates whether the administrator can post stories in the channel (channels only).
	CanPostStories *bool `json:"can_post_stories,omitempty"`

	// CanEditStories indicates whether the administrator can edit stories posted by other users (channels only).
	CanEditStories *bool `json:"can_edit_stories,omitempty"`

	// CanDeleteStories indicates whether the administrator can delete stories posted by other users (channels only).
	CanDeleteStories *bool `json:"can_delete_stories,omitempty"`

	// CanManageTopics indicates whether the user is allowed to create, rename, close, and reopen forum topics (supergroups only).
	CanManageTopics *bool `json:"can_manage_topics,omitempty"`
}

// ChatMemberPermission represents the chat member role permissions.
type ChatMemberPermission int

const (
	// IsAnonymous indicates whether the administrator's presence in the chat is hidden.
	IsAnonymous ChatMemberPermission = iota

	// CanManageChat indicates whether the administrator can access the chat event log,
	// boost list in channels, see channel members, report spam messages,
	// see anonymous administrators in supergroups, and ignore slow mode.
	CanManageChat

	// CanDeleteMessages indicates whether the administrator can delete messages of other users.
	CanDeleteMessages

	// CanManageVideoChats indicates whether the administrator can manage video chats.
	CanManageVideoChats

	// CanRestrictMembers indicates whether the administrator can restrict, ban, or unban chat members,
	// or access supergroup statistics.
	CanRestrictMembers

	// CanPromoteMembers indicates whether the administrator can add new administrators with a subset of their own privileges
	// or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by him).
	CanPromoteMembers

	// CanChangeInfo indicates whether the administrator can change chat title, photo, and other settings.
	CanChangeInfo

	// CanInviteUsers indicates whether the administrator can invite new users to the chat.
	CanInviteUsers

	// CanPostMessages indicates whether the administrator can post messages in the channel
	// or access channel statistics (channels only).
	CanPostMessages

	// CanEditMessages indicates whether the administrator can edit messages of other users
	// and can pin messages (channels only).
	CanEditMessages

	// CanPinMessages indicates whether the administrator can pin messages (supergroups only).
	CanPinMessages

	// CanPostStories indicates whether the administrator can post stories in the channel (channels only).
	CanPostStories

	// CanEditStories indicates whether the administrator can edit stories posted by other users (channels only).
	CanEditStories

	// CanDeleteStories indicates whether the administrator can delete stories posted by other users (channels only).
	CanDeleteStories

	// CanManageTopics indicates whether the user is allowed to create, rename, close, and reopen forum topics (supergroups only).
	CanManageTopics
)

func (p ChatMemberPermission) String() string {
	switch p {
	case IsAnonymous:
		return "IsAnonymous"
	case CanManageChat:
		return "CanManageChat"
	case CanDeleteMessages:
		return "CanDeleteMessages"
	case CanManageVideoChats:
		return "CanManageVideoChats"
	case CanRestrictMembers:
		return "CanRestrictMembers"
	case CanPromoteMembers:
		return "CanPromoteMembers"
	case CanChangeInfo:
		return "CanChangeInfo"
	case CanInviteUsers:
		return "CanInviteUsers"
	case CanPostMessages:
		return "CanPostMessages"
	case CanEditMessages:
		return "CanEditMessages"
	case CanPinMessages:
		return "CanPinMessages"
	case CanPostStories:
		return "CanPostStories"
	case CanEditStories:
		return "CanEditStories"
	case CanDeleteStories:
		return "CanDeleteStories"
	case CanManageTopics:
		return "CanManageTopics"
	default:
		return fmt.Sprintf("Unknown ChatMemberPermission: %d", p)
	}
}

// setChatAdministratorCustomTitleRequest represents the parameters for the setChatAdministratorCustomTitle method.
type setChatAdministratorCustomTitleRequest struct {
	// ChatID is the unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatID interface{} `json:"chat_id"`

	// UserID is the unique identifier of the target user.
	UserID int64 `json:"user_id"`

	// CustomTitle is the new custom title for the administrator; 0-16 characters, emoji is not allowed.
	CustomTitle string `json:"custom_title"`
}

// banChatSenderChatRequest represents the parameters for the banChatSenderChat method.
type banChatSenderChatRequest struct {
	// ChatID is the unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID interface{} `json:"chat_id"`

	// SenderChatID is the unique identifier of the target sender chat.
	SenderChatID int64 `json:"sender_chat_id"`
}

// unbanChatSenderChatRequest represents the parameters for the unbanChatSenderChat method.
type unbanChatSenderChatRequest struct {
	// ChatID is the unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID interface{} `json:"chat_id"`

	// SenderChatID is the unique identifier of the target sender chat.
	SenderChatID int64 `json:"sender_chat_id"`
}

// setChatPermissionsRequest represents the parameters for the setChatPermissions method.
type setChatPermissionsRequest struct {
	// ChatID is the unique identifier for the target chat or username of the target supergroup (in the format @supergroupusername).
	ChatID interface{} `json:"chat_id"`

	// Permissions is a JSON-serialized object for new default chat permissions.
	Permissions ChatPermissions `json:"permissions"`

	// UseIndependentChatPermissions indicates whether chat permissions are set independently.
	// Otherwise, the can_send_other_messages and can_add_web_page_previews permissions will imply the can_send_messages,
	// can_send_audios, can_send_documents, can_send_photos, can_send_videos, can_send_video_notes, and can_send_voice_notes permissions;
	// the can_send_polls permission will imply the can_send_messages permission.
	UseIndependentChatPermissions *bool `json:"use_independent_chat_permissions,omitempty"`
}

// exportChatInviteLinkRequest represents the parameters for the exportChatInviteLink method.
type exportChatInviteLinkRequest struct {
	// ChatID is the unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID interface{} `json:"chat_id"`
}

// createChatInviteLinkRequest represents the parameters for the createChatInviteLink method.
type createChatInviteLinkRequest struct {
	// ChatID is the unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID interface{} `json:"chat_id"`

	// Name is the invite link name; 0-32 characters.
	Name string `json:"name,omitempty"`

	// ExpireDate is the point in time (Unix timestamp) when the link will expire.
	ExpireDate int64 `json:"expire_date,omitempty"`

	// MemberLimit is the maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999.
	MemberLimit int `json:"member_limit,omitempty"`

	// CreatesJoinRequest indicates if users joining the chat via the link need to be approved by chat administrators. If true, MemberLimit can't be specified.
	CreatesJoinRequest bool `json:"creates_join_request,omitempty"`
}

// editChatInviteLinkRequest represents the parameters for the editChatInviteLink method.
type editChatInviteLinkRequest struct {
	// ChatID is the unique identifier for the target chat or username of the target channel (in the format @channelusername).
	ChatID interface{} `json:"chat_id"`

	// InviteLink is the invite link to edit.
	InviteLink string `json:"invite_link"`

	// Name is the invite link name; 0-32 characters.
	Name string `json:"name,omitempty"`

	// ExpireDate is the point in time (Unix timestamp) when the link will expire.
	ExpireDate int64 `json:"expire_date,omitempty"`

	// MemberLimit is the maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999.
	MemberLimit int `json:"member_limit,omitempty"`

	// CreatesJoinRequest indicates if users joining the chat via the link need to be approved by chat administrators. If true, MemberLimit can't be specified.
	CreatesJoinRequest bool `json:"creates_join_request,omitempty"`
}

// revokeChatInviteLinkRequest represents the parameters for the revokeChatInviteLink method.
type revokeChatInviteLinkRequest struct {
	// ChatID is the unique identifier of the target chat or username of the target channel (in the format @channelusername).
	ChatID interface{} `json:"chat_id"`

	// InviteLink is the invite link to revoke.
	InviteLink string `json:"invite_link"`
}

// ChatInviteLink represents an invite link for a chat.
type ChatInviteLink struct {
	// InviteLink is the invite link.
	InviteLink string `json:"invite_link"`

	// Creator is the creator of the link.
	Creator *User `json:"creator,omitempty"`

	// CreatesJoinRequest indicates if users joining the chat via the link need to be approved by chat administrators.
	CreatesJoinRequest bool `json:"creates_join_request,omitempty"`

	// IsPrimary indicates if the link is primary.
	IsPrimary bool `json:"is_primary,omitempty"`

	// IsRevoked indicates if the link is revoked.
	IsRevoked bool `json:"is_revoked,omitempty"`

	// Name is the invite link name.
	Name string `json:"name,omitempty"`

	// ExpireDate is the point in time (Unix timestamp) when the link will expire or has been expired.
	ExpireDate int64 `json:"expire_date,omitempty"`

	// MemberLimit is the maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999.
	MemberLimit int `json:"member_limit,omitempty"`

	// PendingJoinRequestCount is the number of pending join requests created using this link.
	PendingJoinRequestCount int `json:"pending_join_request_count,omitempty"`
}
