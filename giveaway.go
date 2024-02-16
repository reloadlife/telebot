package telebot

// GiveawayCreated represents a service message about the creation of a scheduled giveaway.
// Currently holds no information.
type GiveawayCreated struct{}

// Giveaway represents a message about a scheduled giveaway.
type Giveaway struct {
	// Chats is the list of chats, which the user must join to participate in the giveaway.
	Chats []Chat `json:"chats"`

	// WinnersSelectionDate is the point in time (Unix timestamp) when winners of the giveaway will be selected.
	WinnersSelectionDate int `json:"winners_selection_date"`

	// WinnerCount is the number of users which are supposed to be selected as winners of the giveaway.
	WinnerCount int `json:"winner_count"`

	// OnlyNewMembers is true if only users who join the chats after the giveaway started should be eligible to win.
	OnlyNewMembers bool `json:"only_new_members,omitempty"`

	// HasPublicWinners is true if the list of giveaway winners will be visible to everyone.
	HasPublicWinners bool `json:"has_public_winners,omitempty"`

	// PrizeDescription is the description of additional giveaway prize.
	PrizeDescription string `json:"prize_description,omitempty"`

	// CountryCodes is a list of two-letter ISO 3166-1 alpha-2 country codes indicating the countries from which eligible users for the giveaway must come.
	// If empty, then all users can participate in the giveaway.
	// Users with a phone number that was bought on Fragment can always participate in giveaways.
	CountryCodes []string `json:"country_codes,omitempty"`

	// PremiumSubscriptionMonthCount is the number of months the Telegram Premium subscription won from the giveaway will be active for.
	PremiumSubscriptionMonthCount int `json:"premium_subscription_month_count,omitempty"`
}

// GiveawayWinners represents a message about the completion of a giveaway with public winners.
type GiveawayWinners struct {
	// Chat is the chat that created the giveaway.
	Chat Chat `json:"chat"`

	// GiveawayMessageID is the identifier of the message with the giveaway in the chat.
	GiveawayMessageID int `json:"giveaway_message_id"`

	// WinnersSelectionDate is the point in time (Unix timestamp) when winners of the giveaway were selected.
	WinnersSelectionDate int `json:"winners_selection_date"`

	// WinnerCount is the total number of winners in the giveaway.
	WinnerCount int `json:"winner_count"`

	// Winners is the list of up to 100 winners of the giveaway.
	Winners []User `json:"winners"`

	// AdditionalChatCount is the number of other chats the user had to join in order to be eligible for the giveaway.
	AdditionalChatCount int `json:"additional_chat_count,omitempty"`

	// PremiumSubscriptionMonthCount is the number of months the Telegram Premium subscription won from the giveaway will be active for.
	PremiumSubscriptionMonthCount int `json:"premium_subscription_month_count,omitempty"`

	// UnclaimedPrizeCount is the number of undistributed prizes.
	UnclaimedPrizeCount int `json:"unclaimed_prize_count,omitempty"`

	// OnlyNewMembers is true if only users who had joined the chats after the giveaway started were eligible to win.
	OnlyNewMembers bool `json:"only_new_members,omitempty"`

	// WasRefunded is true if the giveaway was canceled because the payment for it was refunded.
	WasRefunded bool `json:"was_refunded,omitempty"`

	// PrizeDescription is the description of additional giveaway prize.
	PrizeDescription string `json:"prize_description,omitempty"`
}

// GiveawayCompleted represents a service message about the completion of a giveaway without public winners.
type GiveawayCompleted struct {
	// WinnerCount is the number of winners in the giveaway.
	WinnerCount int `json:"winner_count"`

	// UnclaimedPrizeCount is the number of undistributed prizes.
	UnclaimedPrizeCount int `json:"unclaimed_prize_count,omitempty"`

	// GiveawayMessage is the message with the giveaway that was completed if it wasn't deleted.
	GiveawayMessage AccessibleMessage `json:"giveaway_message,omitempty"`
}
