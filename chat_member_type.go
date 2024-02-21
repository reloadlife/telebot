package telebot

import "fmt"

// ChatMember represents information about a member of a chat.
type ChatMember struct {
	Rights

	Status Status `json:"status"`
	User   User   `json:"user"`

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

	// Additional fields based on the ChatMember type
	CustomTitle           string `json:"custom_title,omitempty"`
	CanBeEdited           bool   `json:"can_be_edited,omitempty"`
	IsMember              bool   `json:"is_member,omitempty"`
	CanSendMessages       bool   `json:"can_send_messages,omitempty"`
	CanSendAudios         bool   `json:"can_send_audios,omitempty"`
	CanSendDocuments      bool   `json:"can_send_documents,omitempty"`
	CanSendPhotos         bool   `json:"can_send_photos,omitempty"`
	CanSendVideos         bool   `json:"can_send_videos,omitempty"`
	CanSendVideoNotes     bool   `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes     bool   `json:"can_send_voice_notes,omitempty"`
	CanSendPolls          bool   `json:"can_send_polls,omitempty"`
	CanSendOtherMessages  bool   `json:"can_send_other_messages,omitempty"`
	CanAddWebPagePreviews bool   `json:"can_add_web_page_previews,omitempty"`
	UntilDate             int    `json:"until_date,omitempty"`
}

func (c *ChatMember) ReflectType() string { return fmt.Sprintf("%T", c) }

func (c *ChatMember) Type() string { return "ChatMember" }
