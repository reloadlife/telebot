package tests

import (
	"fmt"
	"github.com/stretchr/testify/require"
	tele "go.mamad.dev/telebot"
	"strings"
	"testing"
)

var (
	messageTypes = []any{
		&tele.AccessibleMessage{},
		&tele.MaybeInaccessibleMessage{},
		&tele.InaccessibleMessage{},
	}

	updateTypes = []any{

		&tele.MessageReaction{},
		&tele.MessageReactionCountUpdated{},
		&tele.InlineQuery{},
		&tele.ChosenInlineResult{},
		&tele.Callback{},
		&tele.ShippingQuery{},
		&tele.PreCheckoutQuery{},
		&tele.Poll{},
		&tele.PollAnswer{},
		&tele.ChatMemberUpdated{},
		&tele.ChatJoinRequest{},
		&tele.ChatBoostUpdated{},
		&tele.ChatBoostRemoved{},
	}

	replyMarkupTypes = []any{
		&tele.InlineKeyboardMarkup{},
	}

	recipients = []any{
		&tele.User{},
		&tele.Chat{},
	}

	listOfTypes = []any{
		&tele.Update{},

		&tele.ChatPhoto{},
		&tele.ChatLocation{},
		&tele.ChatPermissions{},

		&tele.ReactionType{},

		&tele.ChatMember{},
		&tele.ChatMemberUpdated{},

		&tele.ChatInviteLink{},

		&tele.Rights{},

		// AccessibleMessage types
		&tele.MessageOrigin{},
		&tele.ExternalReplyInfo{},
		&tele.TextQuote{},
		&tele.Entity{},
		&tele.LinkPreviewOptions{},

		&tele.Animation{},
		&tele.Audio{},
		&tele.Document{},
		&tele.PhotoSize{},
		&tele.Sticker{},
		&tele.Story{},
		&tele.Video{},
		&tele.VideoNote{},
		&tele.Voice{},
		&tele.Contact{},
		&tele.Dice{},
		&tele.Game{},
		&tele.Poll{},

		&tele.PollOption{},
		&tele.PollAnswer{},

		&tele.Venue{},
		&tele.Location{},

		&tele.AutoDeleteTimerChanged{},

		&tele.Invoice{},
		&tele.SuccessfulPayment{},

		&tele.OrderInfo{},
		&tele.ShippingAddress{},

		&tele.LabeledPrice{},

		&tele.ProximityAlertTriggered{},
		&tele.UsersShared{},
		&tele.ChatShared{},
		&tele.PassportData{},

		&tele.EncryptedPassportElement{},
		&tele.EncryptedCredentials{},
		&tele.PassportFile{},

		&tele.WriteAccessAllowed{},
		&tele.ProximityAlertTriggered{},
		&tele.ForumTopicCreated{},
		&tele.ForumTopicEdited{},
		&tele.ForumTopicReopened{},
		&tele.ForumTopicClosed{},
		&tele.GeneralForumTopicHidden{},
		&tele.GeneralForumTopicUnhidden{},

		&tele.GiveawayCreated{},
		&tele.GiveawayCompleted{},
		&tele.GiveawayWinners{},
		&tele.Giveaway{},

		&tele.VideoChatScheduled{},
		&tele.VideoChatEnded{},
		&tele.VideoChatParticipantsInvited{},
		&tele.VideoChatStarted{},

		&tele.WebAppData{},

		&tele.WebAppInfo{},
	}
)

func init() {
	listOfTypes = append(listOfTypes, updateTypes...)
	listOfTypes = append(listOfTypes, replyMarkupTypes...)
	listOfTypes = append(listOfTypes, recipients...)
	listOfTypes = append(listOfTypes, messageTypes...)
}
func Test_All_Types_Covered(t *testing.T) {
	packageName := "go.mamad.dev/telebot"
	allTypes, err := getAllTypes(packageName)

	// cleanup types

	allTypes = removeDuplicates(allTypes)

	for i, ty := range allTypes {
		// remove those without telebot.
		if !strings.HasPrefix(ty.String(), packageName) {
			allTypes = append(allTypes[:i], allTypes[i+1:]...)
		}
	}

	require.NoError(t, err)

	//require.Equal(t, len(allTypes), len(types), "not all types are covered")
	listOfStringTypes := make([]string, len(listOfTypes))
	for i, typ := range listOfTypes {
		listOfStringTypes[i] = "go.mamad.dev/" + strings.ReplaceAll(fmt.Sprintf("%T", typ), "*", "")
	}
	for _, typ := range allTypes {
		_ = t.Run(fmt.Sprintf("TestTypes/%s", typ), func(t *testing.T) {
			require.Contains(t, listOfStringTypes, fmt.Sprintf("%s", typ), "type %s is not covered", typ)
		})
	}
}
func Test_TT_Type_Compatibility(t *testing.T) {
	for _, typ := range listOfTypes {
		_ = t.Run(fmt.Sprintf("TestTypes/%T", typ), func(t *testing.T) {
			require.Implements(t, (*tele.TType)(nil), typ, "type %T does not implement tele.TType", typ)
			TT, ok := typ.(tele.TType)
			require.True(t, ok, "type %T does not implement tele.TType", typ)
			_ = t.Run(fmt.Sprintf("TestTypes/%T/Type", typ), func(t *testing.T) {
				require.NotPanics(t, func() { _ = TT.Type() })
				require.NotEmptyf(t, TT.Type(), "type %T returned an empty string", typ)
			})
			_ = t.Run(fmt.Sprintf("TestTypes/%T/ReflectType", typ), func(t *testing.T) {
				require.NotPanics(t, func() { _ = TT.ReflectType() })
				require.Equal(t, fmt.Sprintf("%T", typ), TT.ReflectType())
			})
			_ = t.Run(fmt.Sprintf("TestTypes/%T/Verify", typ), func(t *testing.T) {
				require.NotPanics(t, func() { _ = TT.Verify() })
			})
			_ = t.Run(fmt.Sprintf("TestTypes/%T/String", typ), func(t *testing.T) {
				require.NotPanics(t, func() { _ = TT.String() })
			})
			_ = t.Run(fmt.Sprintf("TestTypes/%T/MarshalJSON", typ), func(t *testing.T) {
				require.NotPanics(t, func() { _, _ = TT.MarshalJSON() })
			})
			_ = t.Run(fmt.Sprintf("TestTypes/%T/UnmarshalJSON", typ), func(t *testing.T) {
				require.NotPanics(t, func() { _ = TT.UnmarshalJSON([]byte("{}")) })
			})
		})
	}
}
