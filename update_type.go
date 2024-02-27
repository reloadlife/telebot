package telebot

// String returns the string representation of the UpdateType.
func (ut UpdateType) String() string {
	if ut == "" {
		return Unknown
	}
	return string(ut)
}

// UpdateType returns the type of the Update.
func (u *Update) UpdateType() UpdateType {
	switch {
	case u.Message != nil:
		return UpdateTypeMessage

	case u.EditedMessage != nil:
		return UpdateTypeEditedMessage

	case u.ChannelPost != nil:
		return UpdateTypeChannelPost

	case u.EditedChannelPost != nil:
		return UpdateTypeEditedChannelPost

	case u.Reaction != nil:
		return UpdateTypeMessageReaction

	case u.ReactionCount != nil:
		return UpdateTypeMessageReactionCount

	case u.Query != nil:
		return UpdateTypeInlineQuery

	case u.InlineResult != nil:
		return UpdateTypeChosenInlineResult

	case u.CallbackQuery != nil:
		return UpdateTypeCallbackQuery

	case u.ShippingQuery != nil:
		return UpdateTypeShippingQuery

	case u.PreCheckoutQuery != nil:
		return UpdateTypePreCheckoutQuery

	case u.Poll != nil:
		return UpdateTypePoll

	case u.PollAnswer != nil:
		return UpdateTypePollAnswer

	case u.MyChatMember != nil:
		return UpdateTypeMyChatMember

	case u.ChatMember != nil:
		return UpdateTypeChatMember

	case u.ChatJoinRequest != nil:
		return UpdateTypeChatJoinRequest

	case u.ChatBoost != nil:
		return UpdateTypeChatBoost

	case u.RemovedChatBoost != nil:
		return UpdateTypeChatBoostRemoved

	default:
		return UpdateTypeUnknown
	}
}

func (ut UpdateType) IsValid() bool {
	return ut == UpdateTypeMessage ||
		ut == UpdateTypeEditedMessage ||
		ut == UpdateTypeChannelPost ||
		ut == UpdateTypeEditedChannelPost ||
		ut == UpdateTypeMessageReaction ||
		ut == UpdateTypeMessageReactionCount ||
		ut == UpdateTypeInlineQuery ||
		ut == UpdateTypeChosenInlineResult ||
		ut == UpdateTypeCallbackQuery ||
		ut == UpdateTypeShippingQuery ||
		ut == UpdateTypePreCheckoutQuery ||
		ut == UpdateTypePoll ||
		ut == UpdateTypePollAnswer ||
		ut == UpdateTypeMyChatMember ||
		ut == UpdateTypeChatMember ||
		ut == UpdateTypeChatJoinRequest ||
		ut == UpdateTypeChatBoost ||
		ut == UpdateTypeChatBoostRemoved
}
