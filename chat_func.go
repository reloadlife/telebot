package telebot

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

func (b *bot) Ban(chatID Recipient, userID int64, untilDate *int64, revokeMessages *bool) error {
	params := banChatMemberRequest{
		ChatID: chatID,
		UserID: userID,
	}

	if untilDate != nil {
		params.UntilDate = untilDate
	}

	if revokeMessages != nil {
		params.RevokeMessages = revokeMessages
	}

	_, err := b.sendMethodRequest(methodBanChatMember, params)
	return err
}

func (b *bot) Unban(chatID Recipient, userID int64, onlyIfBanned *bool) error {
	params := unbanChatMemberRequest{
		ChatID:       chatID,
		UserID:       userID,
		OnlyIfBanned: true, // to prevent users getting kicked when UnBan function is called on user that is not already banned.
	}

	if onlyIfBanned != nil {
		params.OnlyIfBanned = *onlyIfBanned
	}

	_, err := b.sendMethodRequest(methodUnbanChatMember, params)
	return err
}
