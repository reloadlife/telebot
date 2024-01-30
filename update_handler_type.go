package telebot

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
