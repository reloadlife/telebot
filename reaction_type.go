package telebot

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

type ReactionType struct {
	Type          ReactionTypeType `json:"type"`
	Emoji         Emoji            `json:"emoji,omitempty"`
	CustomEmojiID string           `json:"custom_emoji_id,omitempty"`
}

type setMessageReactionRequest struct {
	ChatID    any            `json:"chat_id"`
	MessageID int            `json:"message_id"`
	Reaction  []ReactionType `json:"reaction,omitempty"`
	IsBig     bool           `json:"is_big,omitempty"`
}
