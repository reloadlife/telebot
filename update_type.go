package telebot

// String returns the string representation of the UpdateType.
func (ut UpdateType) String() string {
	switch ut {
	case UpdateTypeMessage:
		return "message"
	case UpdateTypeEditedMessage:
		return "edited_message"
	case UpdateTypeChannelPost:
		return "channel_post"
	case UpdateTypeEditedChannelPost:
		return "edited_channel_post"
	case UpdateTypeMessageReaction:
		return "message_reaction"
	case UpdateTypeMessageReactionCount:
		return "message_reaction_count"
	case UpdateTypeInlineQuery:
		return "inline_query"
	case UpdateTypeChosenInlineResult:
		return "chosen_inline_result"
	case UpdateTypeCallbackQuery:
		return "callback_query"
	case UpdateTypeShippingQuery:
		return "shipping_query"
	case UpdateTypePreCheckoutQuery:
		return "pre_checkout_query"
	case UpdateTypePoll:
		return "poll"
	case UpdateTypePollAnswer:
		return "poll_answer"
	case UpdateTypeMyChatMember:
		return "my_chat_member"
	case UpdateTypeChatMember:
		return "chat_member"
	case UpdateTypeChatJoinRequest:
		return "chat_join_request"
	case UpdateTypeChatBoost:
		return "chat_boost"
	case UpdateTypeChatBoostRemoved:
		return "removed_chat_boost"
	default:
		return "unknown"
	}
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
