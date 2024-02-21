package telebot

// a place for all Custom Types, that are either String or Int or ENUM or what ever (except Struct and Interface)

type CustomEmoji string

type IconColor int

const (
	IconColorBlue      IconColor = 7322096
	IconColorOrange    IconColor = 16766590
	IconColorPurple    IconColor = 13338331
	IconColorGreen     IconColor = 9367192
	IconColorRed       IconColor = 16749490
	IconColorOrangeRed IconColor = 16478047
)

type UpdateHandlerOn int

const (
	OnMessage UpdateHandlerOn = iota

	OnCommand
	onMatch

	OnText

	OnPhoto
	OnVideo
	OnVoice
	OnAudio
	OnDocument
	OnSticker
	OnAnimation
	OnVideoNote

	OnLocation
	OnVenue
	OnContact
	OnGame
	OnDice
	OnInvoice
	OnPayment

	OnEditedMessage
	OnChannelPost
	OnEditedChannelPost
	OnCallback

	OnInlineQuery
	OnChosenInlineResult

	OnShippingQuery
	OnPreCheckoutQuery
	OnPoll
	OnPollAnswer
	OnMyChatMember
	OnChatMember
	OnChatJoinRequest

	OnMedia

	OnPinnedMessage

	OnChatBoost
	OnChatBoostRemoved

	OnMessageReaction
	OnMessageReactionCount

	OnAny
)

// UpdateType represents the type of update.
type UpdateType int

// Update types.
const (
	UpdateTypeMessage UpdateType = iota
	UpdateTypeEditedMessage
	UpdateTypeChannelPost
	UpdateTypeEditedChannelPost
	UpdateTypeMessageReaction
	UpdateTypeMessageReactionCount
	UpdateTypeInlineQuery
	UpdateTypeChosenInlineResult
	UpdateTypeCallbackQuery
	UpdateTypeShippingQuery
	UpdateTypePreCheckoutQuery
	UpdateTypePoll
	UpdateTypePollAnswer
	UpdateTypeMyChatMember
	UpdateTypeChatMember
	UpdateTypeChatJoinRequest
	UpdateTypeChatBoost
	UpdateTypeChatBoostRemoved
	UpdateTypeUnknown
)

type MenuButtonType string

const (
	MenuButtonTypeCommands MenuButtonType = "commands"
	MenuButtonTypeWebApp   MenuButtonType = "web_app"
	MenuButtonTypeDefault  MenuButtonType = "default"
)

type PollType string

const (
	PollTypeRegular PollType = "regular"
	PollTypeQuiz    PollType = "quiz"
)

type SwitchInlineQueryStringType string
type SwitchInlineQueryChosenChatStringType string
type RequestContact bool
type RequestLocation bool

type Option int

const (
	Silent Option = iota
	Protected
	HasSpoiler

	RemoveCaption

	DisableContentTypeDetection
	Streamable

	ResizeKeyboard
	OneTimeKeyboard
	PersistentKeyboard
	Selective
)

// markupType is the type of the markup.
// telegram as if v7.1 has 4 types of markup:
type markupType int

const (
	markupTypeInline markupType = iota
	markupTypeKeyboard
	markupTypeRemoveKeyboard
	markupTypeForceReply
)
