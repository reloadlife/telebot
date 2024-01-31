package telebot

import (
	"go.mamad.dev/telebot/.old"
	"testing"
)

func TestAllUpdateTypes(t *testing.T) {
	testCases := []struct {
		updateType   UpdateType
		updateSetter func(u *Update)
	}{
		{UpdateTypeMessage, func(u *Update) { u.Message = &_old.Message{} }},
		{UpdateTypeEditedMessage, func(u *Update) { u.EditedMessage = &_old.Message{} }},
		{UpdateTypeChannelPost, func(u *Update) { u.ChannelPost = &_old.Message{} }},
		{UpdateTypeEditedChannelPost, func(u *Update) { u.EditedChannelPost = &_old.Message{} }},
		{UpdateTypeMessageReaction, func(u *Update) { u.MessageReaction = &_old.MessageReaction{} }},
		{UpdateTypeMessageReactionCount, func(u *Update) { u.MessageReactionCount = &_old.MessageReactionCountUpdated{} }},
		{UpdateTypeInlineQuery, func(u *Update) { u.Query = &_old.Query{} }},
		{UpdateTypeChosenInlineResult, func(u *Update) { u.InlineResult = &_old.InlineResult{} }},
		{UpdateTypeCallbackQuery, func(u *Update) { u.Callback = &Callback{} }},
		{UpdateTypeShippingQuery, func(u *Update) { u.ShippingQuery = &_old.ShippingQuery{} }},
		{UpdateTypePreCheckoutQuery, func(u *Update) { u.PreCheckoutQuery = &_old.PreCheckoutQuery{} }},
		{UpdateTypePoll, func(u *Update) { u.Poll = &_old.Poll{} }},
		{UpdateTypePollAnswer, func(u *Update) { u.PollAnswer = &_old.PollAnswer{} }},
		{UpdateTypeMyChatMember, func(u *Update) { u.MyChatMember = &_old.ChatMemberUpdate{} }},
		{UpdateTypeChatMember, func(u *Update) { u.ChatMember = &_old.ChatMemberUpdate{} }},
		{UpdateTypeChatJoinRequest, func(u *Update) { u.ChatJoinRequest = &_old.ChatJoinRequest{} }},
		{UpdateTypeChatBoost, func(u *Update) { u.ChatBoost = &_old.ChatBoost{} }},
		{UpdateTypeChatBoostRemoved, func(u *Update) { u.ChatBoostRemoved = &_old.ChatBoostRemoved{} }},
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
