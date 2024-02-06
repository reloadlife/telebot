package telebot

type getUserChatBoosts struct {
	ChatId any `json:"chat_id"`
	UserId any `json:"user_id"`
}

type UserChatBoosts struct {
	Boosts []ChatBoost `json:"boosts"`
}

type ChatBoost struct {
	BoostId   string          `json:"boost_id"`
	Date      int             `json:"date"`
	ExpiresAt int             `json:"expires_at"`
	Source    ChatBoostSource `json:"source"`
}

type BoostSource string

const (
	ChatBoostSourcePremium  BoostSource = "premium"
	ChatBoostSourceGiftCode BoostSource = "gift_code"
	ChatBoostSourceGiveaway BoostSource = "giveaway"
)

type ChatBoostSource struct {
	Source            BoostSource `json:"source"`
	User              *User       `json:"user,omitempty"`
	GiveawayMessageId int         `json:"giveaway_message_id,omitempty"`
	IsUnclaimed       bool        `json:"is_unclaimed,omitempty"`
}
