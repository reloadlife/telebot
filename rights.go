package telebot

import "fmt"

type Rights struct {
	IsAnonymous         bool  `json:"is_anonymous"`
	CanManageChat       bool  `json:"can_manage_chat" yaml:"can_manage_chat"`
	CanDeleteMessages   bool  `json:"can_delete_messages" yaml:"can_delete_messages"`
	CanManageVideoChats bool  `json:"can_manage_video_chats" yaml:"can_manage_video_chats"`
	CanRestrictMembers  bool  `json:"can_restrict_members" yaml:"can_restrict_members"`
	CanPromoteMembers   bool  `json:"can_promote_members" yaml:"can_promote_members"`
	CanChangeInfo       bool  `json:"can_change_info" yaml:"can_change_info"`
	CanInviteUsers      bool  `json:"can_invite_users" yaml:"can_invite_users"`
	CanPostMessages     *bool `json:"can_post_messages,omitempty" yaml:"can_post_messages"`
	CanEditMessages     *bool `json:"can_edit_messages,omitempty" yaml:"can_edit_messages"`
	CanPinMessages      *bool `json:"can_pin_messages,omitempty" yaml:"can_pin_messages"`
	CanPostStories      *bool `json:"can_post_stories,omitempty" yaml:"can_post_stories"`
	CanEditStories      *bool `json:"can_edit_stories,omitempty" yaml:"can_edit_stories"`
	CanDeleteStories    *bool `json:"can_delete_stories,omitempty" yaml:"can_delete_stories"`
	CanManageTopics     *bool `json:"can_manage_topics,omitempty" yaml:"can_manage_topics"`
}

func (r *Rights) ReflectType() string {
	return fmt.Sprintf("%T", r)
}

func (r *Rights) Type() string {
	return "Rights"
}
