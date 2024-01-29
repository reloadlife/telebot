package telebot

type Giveaway struct {
	Chats                         []Chat   `json:"chats"`
	WinnersSelectionDate          int      `json:"winners_selection_date"`
	WinnerCount                   int      `json:"winner_count"`
	OnlyNewMembers                bool     `json:"only_new_members,omitempty"`
	HasPublicWinners              bool     `json:"has_public_winners,omitempty"`
	PrizeDescription              string   `json:"prize_description,omitempty"`
	CountryCodes                  []string `json:"country_codes,omitempty"`
	PremiumSubscriptionMonthCount int      `json:"premium_subscription_month_count,omitempty"`
}

type GiveawayWinners struct {
	Chat                          Chat   `json:"chat"`
	GiveawayMessageID             int    `json:"giveaway_message_id"`
	WinnersSelectionDate          int    `json:"winners_selection_date"`
	WinnerCount                   int    `json:"winner_count"`
	Winners                       []User `json:"winners"`
	AdditionalChatCount           int    `json:"additional_chat_count,omitempty"`
	PremiumSubscriptionMonthCount int    `json:"premium_subscription_month_count,omitempty"`
	UnclaimedPrizeCount           int    `json:"unclaimed_prize_count,omitempty"`
	OnlyNewMembers                bool   `json:"only_new_members,omitempty"`
	WasRefunded                   bool   `json:"was_refunded,omitempty"`
	PrizeDescription              string `json:"prize_description,omitempty"`
}

type GiveawayCompleted struct {
	WinnerCount         int     `json:"winner_count"`
	UnclaimedPrizeCount int     `json:"unclaimed_prize_count,omitempty"`
	GiveawayMessage     Message `json:"giveaway_message,omitempty"`
}

type ChatBoostUpdated struct {
	Chat  Chat      `json:"chat"`
	Boost ChatBoost `json:"boost"`
}

type ChatBoostRemoved struct {
	Chat       Chat            `json:"chat"`
	BoostID    string          `json:"boost_id"`
	RemoveDate int             `json:"remove_date"`
	Source     ChatBoostSource `json:"source"`
}

type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"`
}

type ChatBoostSourcePremium struct {
	Source string `json:"source"`
	User   User   `json:"user"`
}

func (c ChatBoostSourcePremium) GetSource() string {
	return c.Source
}

type ChatBoostSourceGiftCode struct {
	Source string `json:"source"`
	User   User   `json:"user"`
}

func (c ChatBoostSourceGiftCode) GetSource() string {
	return c.Source
}

type ChatBoostSourceGiveaway struct {
	Source            string `json:"source"`
	GiveawayMessageID int    `json:"giveaway_message_id"`
	User              User   `json:"user,omitempty"`
	IsUnclaimed       bool   `json:"is_unclaimed,omitempty"`
}

func (c ChatBoostSourceGiveaway) GetSource() string {
	return c.Source
}

type ChatBoostSource struct {
	*ChatBoostSourcePremium
	*ChatBoostSourceGiftCode
	*ChatBoostSourceGiveaway
}

type ChatBoost struct {
	BoostID        string          `json:"boost_id"`
	AddDate        int             `json:"add_date"`
	ExpirationDate int             `json:"expiration_date"`
	Source         ChatBoostSource `json:"source"`
}
