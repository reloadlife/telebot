package telebot

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
