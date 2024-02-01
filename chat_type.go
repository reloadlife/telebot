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

type banChatMemberRequest struct {
	ChatID         any    `json:"chat_id"`
	UserID         int64  `json:"user_id"`
	UntilDate      *int64 `json:"until_date,omitempty"`
	RevokeMessages *bool  `json:"revoke_messages,omitempty"`
}

type unbanChatMemberRequest struct {
	ChatID       any   `json:"chat_id"`
	UserID       int64 `json:"user_id"`
	OnlyIfBanned bool  `json:"only_if_banned,omitempty"`
}

type restrictChatMemberRequest struct {
	ChatID                        any             `json:"chat_id"`
	UserID                        int64           `json:"user_id"`
	Permissions                   ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions bool            `json:"use_independent_chat_permissions,omitempty"`
	UntilDate                     int64           `json:"until_date,omitempty"`
}

type promoteChatMemberRequest struct {
	ChatID              any   `json:"chat_id"`
	UserID              int64 `json:"user_id"`
	IsAnonymous         *bool `json:"is_anonymous,omitempty"`
	CanManageChat       *bool `json:"can_manage_chat,omitempty"`
	CanDeleteMessages   *bool `json:"can_delete_messages,omitempty"`
	CanManageVideoChats *bool `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers  *bool `json:"can_restrict_members,omitempty"`
	CanPromoteMembers   *bool `json:"can_promote_members,omitempty"`
	CanChangeInfo       *bool `json:"can_change_info,omitempty"`
	CanInviteUsers      *bool `json:"can_invite_users,omitempty"`
	CanPostMessages     *bool `json:"can_post_messages,omitempty"`
	CanEditMessages     *bool `json:"can_edit_messages,omitempty"`
	CanPinMessages      *bool `json:"can_pin_messages,omitempty"`
	CanPostStories      *bool `json:"can_post_stories,omitempty"`
	CanEditStories      *bool `json:"can_edit_stories,omitempty"`
	CanDeleteStories    *bool `json:"can_delete_stories,omitempty"`
	CanManageTopics     *bool `json:"can_manage_topics,omitempty"`
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

type setChatAdministratorCustomTitleRequest struct {
	ChatID      any    `json:"chat_id"`
	UserID      int64  `json:"user_id"`
	CustomTitle string `json:"custom_title"`
}

type banChatSenderChatRequest struct {
	ChatID       any   `json:"chat_id"`
	SenderChatID int64 `json:"sender_chat_id"`
}

type unbanChatSenderChatRequest struct {
	ChatID       any   `json:"chat_id"`
	SenderChatID int64 `json:"sender_chat_id"`
}

type setChatPermissionsRequest struct {
	ChatID                        any             `json:"chat_id"`
	Permissions                   ChatPermissions `json:"permissions"`
	UseIndependentChatPermissions *bool           `json:"use_independent_chat_permissions,omitempty"`
}

type exportChatInviteLinkRequest struct {
	ChatID any `json:"chat_id"`
}

type createChatInviteLinkRequest struct {
	ChatID             any    `json:"chat_id"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int64  `json:"expire_date,omitempty"`
	MemberLimit        int    `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

type editChatInviteLinkRequest struct {
	ChatID             any    `json:"chat_id"`
	InviteLink         string `json:"invite_link"`
	Name               string `json:"name,omitempty"`
	ExpireDate         int64  `json:"expire_date,omitempty"`
	MemberLimit        int    `json:"member_limit,omitempty"`
	CreatesJoinRequest bool   `json:"creates_join_request,omitempty"`
}

type revokeChatInviteLinkRequest struct {
	ChatID     any    `json:"chat_id"`
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

type approveChatJoinRequestParams struct {
	ChatID any   `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

type declineChatJoinRequestParams struct {
	ChatID any   `json:"chat_id"`
	UserID int64 `json:"user_id"`
}

type setChatPhotoParams struct {
	ChatID any       `json:"chat_id"`
	Photo  InputFile `json:"photo"`
}

type deleteChatPhotoParams struct {
	ChatID any `json:"chat_id"`
}

type setChatTitleParams struct {
	ChatID any    `json:"chat_id"`
	Title  string `json:"title"`
}

type setChatDescriptionParams struct {
	ChatID      any    `json:"chat_id"`
	Description string `json:"description"`
}
