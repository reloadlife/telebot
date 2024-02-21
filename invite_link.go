package telebot

import "fmt"

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

func (c *ChatInviteLink) ReflectType() string {
	return fmt.Sprintf("%T", c)
}

func (c *ChatInviteLink) Type() string {
	return "ChatInviteLink"
}
