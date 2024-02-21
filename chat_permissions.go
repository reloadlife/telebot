package telebot

import "fmt"

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

func (c *ChatPermissions) Type() string        { return "chat_permissions" }
func (c *ChatPermissions) ReflectType() string { return fmt.Sprintf("%T", c) }
