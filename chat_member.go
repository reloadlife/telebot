package telebot

import "fmt"

// ChatMemberUpdated represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	// Chat is the chat the user belongs to.
	Chat Chat `json:"chat"`

	// From is the performer of the action, which resulted in the change.
	From User `json:"from"`

	// Date is the date the change was done in Unix time.
	Date int64 `json:"date"`

	// OldChatMember is the previous information about the chat member.
	OldChatMember ChatMember `json:"old_chat_member"`

	// NewChatMember is the new information about the chat member.
	NewChatMember ChatMember `json:"new_chat_member"`

	// InviteLink is the chat invite link, which was used by the user to join the chat (optional, for joining by invite link events only).
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`

	// ViaChatFolderInviteLink is true if the user joined the chat via a chat folder invite link (optional).
	ViaChatFolderInviteLink bool `json:"via_chat_folder_invite_link,omitempty"`
}

func (c *ChatMemberUpdated) ReflectType() string {
	return fmt.Sprintf("%T", c)
}

func (c *ChatMemberUpdated) Type() string {
	return "ChatMemberUpdated"
}
