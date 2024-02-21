package telebot

type Rights struct {
	IsAnonymous         bool  `json:"is_anonymous"`
	CanManageChat       bool  `json:"can_manage_chat"`
	CanDeleteMessages   bool  `json:"can_delete_messages"`
	CanManageVideoChats bool  `json:"can_manage_video_chats"`
	CanRestrictMembers  bool  `json:"can_restrict_members"`
	CanPromoteMembers   bool  `json:"can_promote_members"`
	CanChangeInfo       bool  `json:"can_change_info"`
	CanInviteUsers      bool  `json:"can_invite_users"`
	CanPostMessages     *bool `json:"can_post_messages,omitempty"`
	CanEditMessages     *bool `json:"can_edit_messages,omitempty"`
	CanPinMessages      *bool `json:"can_pin_messages,omitempty"`
	CanPostStories      *bool `json:"can_post_stories,omitempty"`
	CanEditStories      *bool `json:"can_edit_stories,omitempty"`
	CanDeleteStories    *bool `json:"can_delete_stories,omitempty"`
	CanManageTopics     *bool `json:"can_manage_topics,omitempty"`
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
	ChatID any  `json:"chat_id"`
	Photo  File `json:"photo"`
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

type pinChatMessageRequest struct {
	ChatID              any  `json:"chat_id"`
	MessageID           int  `json:"message_id"`
	DisableNotification bool `json:"disable_notification,omitempty"`
}

type unpinChatMessageRequest struct {
	ChatID    any `json:"chat_id"`
	MessageID int `json:"message_id,omitempty"`
}

type unpinAllChatMessagesRequest struct {
	ChatID any `json:"chat_id"`
}

type leaveChatRequest struct {
	ChatID any `json:"chat_id"`
}

type getChatRequest struct {
	ChatID interface{} `json:"chat_id"`
}

type getChatAdministratorsRequest struct {
	ChatID interface{} `json:"chat_id"`
}

type getChatMemberCountRequest struct {
	ChatID interface{} `json:"chat_id"`
}

type getChatMemberRequest struct {
	ChatID interface{} `json:"chat_id"`
	UserID int64       `json:"user_id"`
}

type setChatStickerSetRequest struct {
	ChatID         any    `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

type deleteChatStickerSetRequest struct {
	ChatID any `json:"chat_id"`
}
