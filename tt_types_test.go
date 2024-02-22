package telebot

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"go/types"
	"golang.org/x/tools/go/packages"
	"reflect"
	"strings"
	"testing"
)

var (
	messageTypes = []any{
		&AccessibleMessage{},
		&MaybeInaccessibleMessage{},
		&InaccessibleMessage{},
	}
	// TTType ✅

	updateTypes = []any{
		&MessageReaction{},
		&MessageReactionCountUpdated{},
		&InlineQuery{},
		&ChosenInlineResult{},
		&Callback{},
		&ShippingQuery{},
		&PreCheckoutQuery{},
		&Poll{},
		&PollAnswer{},
		&ChatMemberUpdated{},
		&ChatJoinRequest{},
		&ChatBoostUpdated{},
		&ChatBoostRemoved{},
	}
	// TTType ✅

	replyMarkupTypes = []any{
		&InlineKeyboardButton{},
		&InlineKeyboardMarkup{},
		&ReplyKeyboardMarkup{},
		&ReplyKeyboardRemove{},
	}

	recipients = []any{
		&User{},
		&Chat{},
	}
	// TTType ✅

	listOfTypes = []any{
		&Update{},

		&ChatPhoto{},
		&ChatLocation{},
		&ChatPermissions{},

		&ReactionType{},

		&ChatMember{},
		&ChatMemberUpdated{},

		&ChatInviteLink{},

		&Rights{},

		// AccessibleMessage types
		&MessageOrigin{},
		&ExternalReplyInfo{},
		&TextQuote{},
		&Entity{},
		&LinkPreviewOptions{},

		&Animation{},
		&Audio{},
		&Document{},
		&PhotoSize{},
		&Sticker{},
		&Story{},
		&Video{},
		&VideoNote{},
		&Voice{},
		&Contact{},
		&Dice{},
		&Game{},
		&PollOption{},

		&Venue{},
		&Location{},

		&AutoDeleteTimerChanged{},

		&Invoice{},
		&SuccessfulPayment{},

		&OrderInfo{},
		&ShippingAddress{},

		&LabeledPrice{},

		&ProximityAlertTriggered{},
		&UsersShared{},
		&ChatShared{},
		&PassportData{},

		&EncryptedPassportElement{},
		&EncryptedCredentials{},
		&PassportFile{},

		&WriteAccessAllowed{},
		&ProximityAlertTriggered{},
		&ForumTopicCreated{},
		&ForumTopicEdited{},
		&ForumTopicReopened{},
		&ForumTopicClosed{},
		&GeneralForumTopicHidden{},
		&GeneralForumTopicUnhidden{},

		&GiveawayCreated{},
		&GiveawayCompleted{},
		&GiveawayWinners{},
		&Giveaway{},

		&VideoChatScheduled{},
		&VideoChatEnded{},
		&VideoChatParticipantsInvited{},
		&VideoChatStarted{},

		&WebAppData{},

		&WebAppInfo{},

		&BotCommand{},
		&BotCommandScope{},

		&UserProfilePhotos{},
		&UserChatBoosts{},

		&SwitchInlineQueryChosenChat{},
		&StickerSet{},

		&SentWebAppMessage{},
		&ReplyParameters{},
		&QueryResult{},
		&PassportElementError{},
		&MenuButton{},
		&MaskPosition{},

		&KeyboardButton{},
		&KeyboardButtonPollType{},
		&KeyboardButtonRequestUsers{},
		&KeyboardButtonRequestChat{},
		&GameHighScore{},

		&ForumTopic{},

		&File{},

		&ChatBoost{},
		&ChatBoostSource{},

		&ReactionCount{},

		&InputContactMessageContent{},
		&InputLocationMessageContent{},
		&InputTextMessageContent{},
		&InputVenueMessageContent{},
		&InputInvoiceMessageContent{},
		&InputSticker{},
		&InputMedia{},

		&LoginURL{},
		&CallbackGame{},

		&ChatBoostAdded{},
	}

	notStructTypes = []string{
		"BoostSource",
		"Bot",
		"BotCommandScopeType",

		"BotSettings",
		"CallbackEndpoint",
		"ChatMemberPermission",
		"ChatAction",
		"ChatType",
		"Context",
		"CustomEmoji",
		"Option",
		"Emoji",
		"EntityType",
		"Width",
		"Height",
		"Verifiable",

		"Updates",
		"UpdateType",
		"Title",
		"TType",
		"StickerEmoji",
		"Status",

		"ReplyMarkup",
		"Recipient",
		"ReactionTypeType",

		"QueryResults",
		"Poller",
		"Performer",

		"ParseMode",
		"UpdateHandlerOn",

		"MiddlewarePoller",
		"MiddlewareFunc",
		"MessageThreadID",
		"MessageOriginType",

		"Message",

		"LongPoller",
		"InlineQueryResultType",
		"IconColor",

		"Group",
		"GroupError",

		"HandlerFunc",

		"Error",
		"FloodError",

		"InputMessageContent",
		"InputMediaType",

		"PollType",
		"SwitchInlineQueryStringType",
		"SwitchInlineQueryChosenChatStringType",
		"Button",
		"Row",
		"RequestLocation",
		"RequestContact",
		"ForceReplyMarkup",

		"MenuButtonType",
		"Equal",

		"PhotoSizes",
		"Stringer",
		"JSONer",
	}
)

func getAllTypes(packageName string) ([]types.Type, error) {
	// Load the package
	cfg := &packages.Config{
		Mode: packages.NeedTypes,
	}

	pkgs, err := packages.Load(cfg, packageName)
	if err != nil {
		return nil, fmt.Errorf("error loading package %s: %w", packageName, err)
	}

	// Get all types from the package
	var allTypes []types.Type
	for _, pkg := range pkgs {
		for _, name := range pkg.Types.Scope().Names() {
			// Construct the fully qualified type name
			typeFullName := name
			// isPrivate ? ignore it
			if !pkg.Types.Scope().Lookup(name).Exported() {
				//_, _ = fmt.Println("Ignoring private type: ", typeFullName)
				continue
			}

			// Resolve the type by name
			typ, err := getTypeByName(typeFullName, pkgs)
			if err != nil {
				return nil, fmt.Errorf("error resolving type %s: %w", typeFullName, err)
			}

			allTypes = append(allTypes, typ)
		}
	}

	return allTypes, nil
}

// getTypeByName returns the reflect.Type for a fully qualified type name
func getTypeByName(typeName string, pkgs []*packages.Package) (types.Type, error) {
	// Attempt to find the type by name
	for _, pkg := range pkgs {
		if typ := pkg.Types.Scope().Lookup(typeName); typ != nil {
			return typ.Type(), nil
		}
	}

	return nil, fmt.Errorf("type not found: %s", typeName)
}

func removeDuplicates[T any](slice []T) []T {
	keys := make(map[any]bool)
	list := []T{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func isStruct(v interface{}) bool {
	t := reflect.TypeOf(v)

	// Check if the type is a struct and not an interface
	return t.Kind() == reflect.Struct
}

func init() {
	listOfTypes = append(listOfTypes, updateTypes...)
	listOfTypes = append(listOfTypes, replyMarkupTypes...)
	listOfTypes = append(listOfTypes, recipients...)
	listOfTypes = append(listOfTypes, messageTypes...)
}
func Test_All_Types_Covered(t *testing.T) {
	packageName := "go.mamad.dev/telebot"
	allTypes, err := getAllTypes(packageName)
	allTypes = removeDuplicates(allTypes)

	for i, ty := range allTypes {
		if !strings.HasPrefix(ty.String(), packageName) {
			allTypes = append(allTypes[:i], allTypes[i+1:]...)
		}
	}

	require.NoError(t, err)

	listOfStringTypes := make([]string, len(listOfTypes))
	for i, typ := range listOfTypes {
		listOfStringTypes[i] = "go.mamad.dev/" + strings.ReplaceAll(fmt.Sprintf("%T", typ), "*", "")
	}

	for _, not := range notStructTypes {
		listOfStringTypes = append(listOfStringTypes, "go.mamad.dev/telebot."+not)
	}

	for _, typ := range allTypes {
		typName := fmt.Sprintf("TestTypes/%s", typ)
		if strings.Contains(typName, "func") ||
			typName == "TestTypes/*go.mamad.dev/telebot.Error" { // a duplicate type, "Error" is already covered.
			continue
		}
		_ = t.Run(typName, func(t *testing.T) {
			require.Contains(t, listOfStringTypes, fmt.Sprintf("%s", typ), "type %s is not covered", typ)
		})
	}
}

func Test_TT_Type_CompatibilityVerifiable(t *testing.T) {
	for _, typ := range listOfTypes {
		_ = t.Run(fmt.Sprintf("TestTypes/Verifiable/%T", typ), func(t *testing.T) {
			require.Implements(t, (*Verifiable)(nil), typ, "type %T does not implement Verifiable", typ)
			V, ok := typ.(Verifiable)
			require.True(t, ok, "type %T does not implement Verifiable", typ)
			_ = t.Run(fmt.Sprintf("TestTypes/Verifiable/%T/Verify", typ), func(t *testing.T) {
				require.NotPanics(t, func() { _ = V.Verify() })
			})
		})
	}
}
func Test_TT_Type_CompatibilityString(t *testing.T) {
	for _, typ := range listOfTypes {
		_ = t.Run(fmt.Sprintf("TestTypes/String/%T", typ), func(t *testing.T) {
			require.Implements(t, (*Stringer)(nil), typ, "type %T does not implement Stringer", typ)
			V, ok := typ.(Stringer)
			require.True(t, ok, "type %T does not implement Stringer", typ)
			_ = t.Run(fmt.Sprintf("TestTypes/String/%T/String", typ), func(t *testing.T) {
				require.NotPanics(t, func() {
					require.NotEmptyf(t, V.String(), "type %T returned an empty string", typ)
				})
			})
		})
	}
}
func Test_TT_Type_CompatibilityJSON(t *testing.T) {
	for _, typ := range listOfTypes {
		_ = t.Run(fmt.Sprintf("TestTypes/JSONer/%T", typ), func(t *testing.T) {
			require.Implements(t, (*JSONer)(nil), typ, "type %T does not implement JSONer", typ)
			V, ok := typ.(JSONer)
			require.True(t, ok, "type %T does not implement JSONer", typ)
			_ = t.Run(fmt.Sprintf("TestTypes/JSONer/%T/JSONMarshal", typ), func(t *testing.T) {
				require.NotPanics(t, func() {
					b, err := V.MarshalJSON()
					require.NoError(t, err)
					require.NotEmptyf(t, b, "type %T returned an empty string", typ)
				})
			})
			_ = t.Run(fmt.Sprintf("TestTypes/JSONer/%T/JSONUnmarshal", typ), func(t *testing.T) {
				require.NotPanics(t, func() {
					err := V.UnmarshalJSON([]byte("{}"))
					require.NoError(t, err)
				})
			})
		})
	}
}
func Test_TT_Type_CompatibilityTType(t *testing.T) {
	for _, typ := range listOfTypes {
		_ = t.Run(fmt.Sprintf("TestTypes/TType/%T", typ), func(t *testing.T) {
			require.Implements(t, (*TType)(nil), typ, "type %T does not implement TType", typ)
			TT, ok := typ.(TType)
			require.True(t, ok, "type %T does not implement TType", typ)
			_ = t.Run(fmt.Sprintf("TestTypes/TType/%T/Type", typ), func(t *testing.T) {
				require.NotPanics(t, func() { _ = TT.Type() })
				require.NotEmptyf(t, TT.Type(), "type %T returned an empty string", typ)
			})
			_ = t.Run(fmt.Sprintf("TestTypes/TType/%T/ReflectType", typ), func(t *testing.T) {
				require.NotPanics(t, func() { _ = TT.ReflectType() })
				require.Equal(t, fmt.Sprintf("%T", typ), TT.ReflectType())
			})
			_ = t.Run(fmt.Sprintf("TestTypes/TType/%T/Verify", typ), func(t *testing.T) {
				require.NotPanics(t, func() { _ = TT.Verify() })
			})
			_ = t.Run(fmt.Sprintf("TestTypes/TType/%T/String", typ), func(t *testing.T) {
				require.NotPanics(t, func() { _ = TT.String() })
			})
			_ = t.Run(fmt.Sprintf("TestTypes/TType/%T/MarshalJSON", typ), func(t *testing.T) {
				require.NotPanics(t, func() { _, _ = TT.MarshalJSON() })
			})
			_ = t.Run(fmt.Sprintf("TestTypes/TType/%T/UnmarshalJSON", typ), func(t *testing.T) {
				require.NotPanics(t, func() { _ = TT.UnmarshalJSON([]byte("{}")) })
			})
		})
	}
}
