package telebot

import (
	"testing"
)

func TestAllUpdateTypes(t *testing.T) {
	testCases := []struct {
		updateType   UpdateType
		updateSetter func(u *Update)
	}{
		{UpdateTypeMessage, func(u *Update) { u.Message = &Message{} }},
		{UpdateTypeEditedMessage, func(u *Update) { u.EditedMessage = &Message{} }},
		{UpdateTypeChannelPost, func(u *Update) { u.ChannelPost = &Message{} }},
		{UpdateTypeEditedChannelPost, func(u *Update) { u.EditedChannelPost = &Message{} }},
		{UpdateTypeMessageReaction, func(u *Update) { u.MessageReaction = &MessageReaction{} }},
		{UpdateTypeMessageReactionCount, func(u *Update) { u.MessageReactionCount = &MessageReactionCountUpdated{} }},
		{UpdateTypeInlineQuery, func(u *Update) { u.Query = &Query{} }},
		{UpdateTypeChosenInlineResult, func(u *Update) { u.InlineResult = &InlineResult{} }},
		{UpdateTypeCallbackQuery, func(u *Update) { u.Callback = &Callback{} }},
		{UpdateTypeShippingQuery, func(u *Update) { u.ShippingQuery = &ShippingQuery{} }},
		{UpdateTypePreCheckoutQuery, func(u *Update) { u.PreCheckoutQuery = &PreCheckoutQuery{} }},
		{UpdateTypePoll, func(u *Update) { u.Poll = &Poll{} }},
		{UpdateTypePollAnswer, func(u *Update) { u.PollAnswer = &PollAnswer{} }},
		{UpdateTypeMyChatMember, func(u *Update) { u.MyChatMember = &ChatMemberUpdate{} }},
		{UpdateTypeChatMember, func(u *Update) { u.ChatMember = &ChatMemberUpdate{} }},
		{UpdateTypeChatJoinRequest, func(u *Update) { u.ChatJoinRequest = &ChatJoinRequest{} }},
		{UpdateTypeChatBoost, func(u *Update) { u.ChatBoost = &ChatBoost{} }},
		{UpdateTypeChatBoostRemoved, func(u *Update) { u.ChatBoostRemoved = &ChatBoostRemoved{} }},
		{UpdateTypeUnknown, func(u *Update) {}},
	}

	for _, testCase := range testCases {
		_ = t.Run(testCase.updateType.String(), func(t *testing.T) {
			var u Update
			testCase.updateSetter(&u)
			if u.Type() != testCase.updateType {
				t.Fatalf("wrong update type for %s", testCase.updateType.String())
			}
		})
	}
}
