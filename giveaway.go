package telebot

type Giveaway struct {
	Chats                         []Chat   `json:"chats"`                                      // The list of chats which the user must join to participate in the giveaway
	WinnersSelectionDate          int      `json:"winners_selection_date"`                     // Point in time (Unix timestamp) when winners of the giveaway will be selected
	WinnerCount                   int      `json:"winner_count"`                               // The number of users which are supposed to be selected as winners of the giveaway
	OnlyNewMembers                bool     `json:"only_new_members,omitempty"`                 // Optional. True, if only users who join the chats after the giveaway started should be eligible to win
	HasPublicWinners              bool     `json:"has_public_winners,omitempty"`               // Optional. True, if the list of giveaway winners will be visible to everyone
	PrizeDescription              string   `json:"prize_description,omitempty"`                // Optional. Description of additional giveaway prize
	CountryCodes                  []string `json:"country_codes,omitempty"`                    // Optional. A list of two-letter ISO 3166-1 alpha-2 country codes indicating the countries from which eligible users for the giveaway must come. If empty, then all users can participate in the giveaway. Users with a phone number that was bought on Fragment can always participate in giveaways.
	PremiumSubscriptionMonthCount int      `json:"premium_subscription_month_count,omitempty"` // Optional. The number of months the Telegram Premium subscription won from the giveaway will be active for
}

type GiveawayWinners struct {
	Chat                          Chat   `json:"chat"`                                       // The chat that created the giveaway
	GiveawayMessageID             int    `json:"giveaway_message_id"`                        // Identifier of the message with the giveaway in the chat
	WinnersSelectionDate          int    `json:"winners_selection_date"`                     // Point in time (Unix timestamp) when winners of the giveaway were selected
	WinnerCount                   int    `json:"winner_count"`                               // Total number of winners in the giveaway
	Winners                       []User `json:"winners"`                                    // List of up to 100 winners of the giveaway
	AdditionalChatCount           int    `json:"additional_chat_count,omitempty"`            // Optional. The number of other chats the user had to join in order to be eligible for the giveaway
	PremiumSubscriptionMonthCount int    `json:"premium_subscription_month_count,omitempty"` // Optional. The number of months the Telegram Premium subscription won from the giveaway will be active for
	UnclaimedPrizeCount           int    `json:"unclaimed_prize_count,omitempty"`            // Optional. Number of undistributed prizes
	OnlyNewMembers                bool   `json:"only_new_members,omitempty"`                 // Optional. True, if only users who had joined the chats after the giveaway started were eligible to win
	WasRefunded                   bool   `json:"was_refunded,omitempty"`                     // Optional. True, if the giveaway was canceled because the payment for it was refunded
	PrizeDescription              string `json:"prize_description,omitempty"`                // Optional. Description of additional giveaway prize
}

type GiveawayCompleted struct {
	WinnerCount         int     `json:"winner_count"`                    // Number of winners in the giveaway
	UnclaimedPrizeCount int     `json:"unclaimed_prize_count,omitempty"` // Optional. Number of undistributed prizes
	GiveawayMessage     Message `json:"giveaway_message,omitempty"`      // Optional. Message with the giveaway that was completed, if it wasn't deleted
}
