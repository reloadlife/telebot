package telebot

// a place for all Custom Types, that are either String or Int or ENUM or what ever (except Struct and Interface)

const Unknown = "unknown"

type InputMediaType string

const (
	InputMediaAnimation InputMediaType = "animation"
	InputMediaDocument  InputMediaType = "document"
	InputMediaAudio     InputMediaType = "audio"
	InputMediaPhoto     InputMediaType = "photo"
	InputMediaVideo     InputMediaType = "video"
)

type InlineQueryResultType string

const (
	InlineQueryResultCachedAudio    InlineQueryResultType = "audio"
	InlineQueryResultCachedDocument InlineQueryResultType = "document"
	InlineQueryResultCachedGif      InlineQueryResultType = "gif"
	InlineQueryResultCachedMpeg4Gif InlineQueryResultType = "mpeg4_gif"
	InlineQueryResultCachedPhoto    InlineQueryResultType = "photo"
	InlineQueryResultCachedSticker  InlineQueryResultType = "sticker"
	InlineQueryResultCachedVideo    InlineQueryResultType = "video"
	InlineQueryResultCachedVoice    InlineQueryResultType = "voice"
	InlineQueryResultArticle        InlineQueryResultType = "article"
	InlineQueryResultAudio          InlineQueryResultType = "audio"
	InlineQueryResultContact        InlineQueryResultType = "contact"
	InlineQueryResultGame           InlineQueryResultType = "game"
	InlineQueryResultDocument       InlineQueryResultType = "document"
	InlineQueryResultGif            InlineQueryResultType = "gif"
	InlineQueryResultLocation       InlineQueryResultType = "location"
	InlineQueryResultMpeg4Gif       InlineQueryResultType = "mpeg4_gif"
	InlineQueryResultPhoto          InlineQueryResultType = "photo"
	InlineQueryResultVenue          InlineQueryResultType = "venue"
	InlineQueryResultVideo          InlineQueryResultType = "video"
	InlineQueryResultVoice          InlineQueryResultType = "voice"
)

type BoostSource string

const (
	ChatBoostSourcePremium  BoostSource = "premium"
	ChatBoostSourceGiftCode BoostSource = "gift_code"
	ChatBoostSourceGiveaway BoostSource = "giveaway"
)

type BotCommandScopeType string

const (
	BotCommandScopeDefault               BotCommandScopeType = "default"
	BotCommandScopeAllPrivateChats       BotCommandScopeType = "all_private_chats"
	BotCommandScopeAllGroupChats         BotCommandScopeType = "all_group_chats"
	BotCommandScopeAllChatAdministrators BotCommandScopeType = "all_chat_administrators"
	BotCommandScopeChat                  BotCommandScopeType = "chat"
	BotCommandScopeChatAdministrators    BotCommandScopeType = "chat_administrators"
	BotCommandScopeChatMember            BotCommandScopeType = "chat_member"
)

type MessageOriginType string

const (
	MessageOriginUser       MessageOriginType = "user"
	MessageOriginHiddenUser MessageOriginType = "hidden_user"
	MessageOriginChat       MessageOriginType = "chat"
	MessageOriginChannel    MessageOriginType = "channel"
)

type StickerEmoji string
type Emoji StickerEmoji

const (
	EmojiLike       Emoji = "👍"
	EmojiDislike    Emoji = "👎"
	EmojiHeart      Emoji = "❤"
	EmojiFire       Emoji = "🔥"
	EmojiLove       Emoji = "🥰"
	EmojiClap       Emoji = "👏"
	EmojiSmile      Emoji = "😁"
	EmojiThink      Emoji = "🤔"
	EmojiMindBlown  Emoji = "🤯"
	EmojiScream     Emoji = "😱"
	EmojiAngry      Emoji = "🤬"
	EmojiCry        Emoji = "😢"
	EmojiParty      Emoji = "🎉"
	EmojiExcited    Emoji = "🤩"
	EmojiVomit      Emoji = "🤮"
	EmojiPoop       Emoji = "💩"
	EmojiPray       Emoji = "🙏"
	EmojiOk         Emoji = "👌"
	EmojiDove       Emoji = "🕊"
	EmojiClown      Emoji = "🤡"
	EmojiYawn       Emoji = "🥱"
	EmojiDrunk      Emoji = "🥴"
	EmojiHeartEyes  Emoji = "😍"
	EmojiDolphin    Emoji = "🐳"
	EmojiHeartFire  Emoji = "❤‍🔥"
	EmojiMoon       Emoji = "🌚"
	EmojiHotdog     Emoji = "🌭"
	EmojiHundred    Emoji = "💯"
	EmojiLaugh      Emoji = "🤣"
	EmojiLightning  Emoji = "⚡"
	EmojiBanana     Emoji = "🍌"
	EmojiTrophy     Emoji = "🏆"
	EmojiBroken     Emoji = "💔"
	EmojiSkeptical  Emoji = "🤨"
	EmojiNeutral    Emoji = "😐"
	EmojiStrawberry Emoji = "🍓"
	EmojiChampagne  Emoji = "🍾"
	EmojiKiss       Emoji = "💋"
	EmojiFuckYou    Emoji = "🖕"
	EmojiEvil       Emoji = "😈"
	EmojiSleep      Emoji = "😴"
	EmojiCrying     Emoji = "😭"
	EmojiNerd       Emoji = "🤓"
	EmojiGhost      Emoji = "👻"
	EmojiComputer   Emoji = "👨‍💻"
	EmojiEyes       Emoji = "👀"
	EmojiJackO      Emoji = "🎃"
	EmojiSeeNoEvil  Emoji = "🙈"
	EmojiAngel      Emoji = "😇"
	EmojiFearful    Emoji = "😨"
	EmojiHandshake  Emoji = "🤝"
	EmojiWriting    Emoji = "✍"
	EmojiHugging    Emoji = "🤗"
	EmojiSmiling    Emoji = "🫡"
	EmojiSanta      Emoji = "🎅"
	EmojiChristmas  Emoji = "🎄"
	EmojiSnowman    Emoji = "☃"
	EmojiNailPolish Emoji = "💅"
	EmojiZany       Emoji = "🤪"
	EmojiMoai       Emoji = "🗿"
	EmojiCool       Emoji = "🆒"
	EmojiHearts     Emoji = "💘"
	EmojiHearNoEvil Emoji = "🙉"
	EmojiUnicorn    Emoji = "🦄"
	EmojiKissing    Emoji = "😘"
	EmojiPill       Emoji = "💊"
	EmojiSayNoEvil  Emoji = "🙊"
	EmojiSunglasses Emoji = "😎"
	EmojiAlien      Emoji = "👾"
	EmojiMan        Emoji = "🤷‍♂"
	EmojiWoman      Emoji = "🤷"
	EmojiWoman2     Emoji = "🤷‍♀"
	EmojiAngry2     Emoji = "😡"
)

type ReactionTypeType string

const (
	ReactionTypeTypeEmoji       ReactionTypeType = "emoji"
	ReactionTypeTypeCustomEmoji ReactionTypeType = "custom_emoji"
)

// ChatMemberPermission represents the chat member role permissions.
type ChatMemberPermission int

const (
	// IsAnonymous indicates whether the administrator's presence in the chat is hidden.
	IsAnonymous ChatMemberPermission = iota

	// CanManageChat indicates whether the administrator can access the chat event log,
	// boost list in channels, see channel members, report spam messages,
	// see anonymous administrators in supergroups, and ignore slow mode.
	CanManageChat

	// CanDeleteMessages indicates whether the administrator can delete messages of other users.
	CanDeleteMessages

	// CanManageVideoChats indicates whether the administrator can manage video chats.
	CanManageVideoChats

	// CanRestrictMembers indicates whether the administrator can restrict, ban, or unban chat members,
	// or access supergroup statistics.
	CanRestrictMembers

	// CanPromoteMembers indicates whether the administrator can add new administrators with a subset of their own privileges
	// or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by him).
	CanPromoteMembers

	// CanChangeInfo indicates whether the administrator can change chat title, photo, and other settings.
	CanChangeInfo

	// CanInviteUsers indicates whether the administrator can invite new users to the chat.
	CanInviteUsers

	// CanPostMessages indicates whether the administrator can post messages in the channel
	// or access channel statistics (channels only).
	CanPostMessages

	// CanEditMessages indicates whether the administrator can edit messages of other users
	// and can pin messages (channels only).
	CanEditMessages

	// CanPinMessages indicates whether the administrator can pin messages (supergroups only).
	CanPinMessages

	// CanPostStories indicates whether the administrator can post stories in the channel (channels only).
	CanPostStories

	// CanEditStories indicates whether the administrator can edit stories posted by other users (channels only).
	CanEditStories

	// CanDeleteStories indicates whether the administrator can delete stories posted by other users (channels only).
	CanDeleteStories

	// CanManageTopics indicates whether the user is allowed to create, rename, close, and reopen forum topics (supergroups only).
	CanManageTopics
)

// Status represents the possible statuses of a chat member.
type Status string

// Enum values for the different statuses of a chat member.
const (
	StatusCreator       Status = "creator"
	StatusAdministrator Status = "administrator"
	StatusMember        Status = "member"
	StatusRestricted    Status = "restricted"
	StatusLeft          Status = "left"
	StatusKicked        Status = "kicked"
)

type ChatType string

const (
	ChatTypeSender ChatType = "sender" // Only for InlineQuery.

	ChatPrivate        ChatType = "private"
	ChatGroup          ChatType = "group"
	ChatSuperGroup     ChatType = "supergroup"
	ChatChannel        ChatType = "channel"
	ChatChannelPrivate ChatType = "privatechannel"
)

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
