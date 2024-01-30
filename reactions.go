package telebot

import (
	"strconv"
)

type MessageReaction struct {
	Chat        *Chat          `json:"chat"`
	MessageId   int            `json:"message_id"`
	User        *User          `json:"user,omitempty"`
	ActorChat   *Chat          `json:"actor_chat,omitempty"`
	Date        int            `json:"date"`
	OldReaction []ReactionType `json:"old_reaction"`
	NewReaction []ReactionType `json:"new_reaction"`
}

type ReactionType struct {
	// Type one of "emoji" | "custom_emoji".
	Type string `json:"type"`

	Emoji         *Emoji `json:"emoji,omitempty"`
	CustomEmojiId string `json:"custom_emoji_id,omitempty"`
}

// Emoji Reaction emoji
type Emoji string

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

type MessageReactionCountUpdated struct {
	Chat      *Chat          `json:"chat"`
	MessageId int            `json:"message_id"`
	Date      int            `json:"date"`
	Reactions []ReactionType `json:"reactions"`
}

func (b *OldBot) SetMessageReaction(where Recipient, messageId int, isBig bool, reaction ...ReactionType) (bool, error) {
	params := map[string]any{
		"chat_id":    where.Recipient(),
		"message_id": strconv.Itoa(messageId),
		"is_big":     isBig,
		"reactions":  reaction,
	}

	_, err := b.Raw("SetMessageReaction", params)
	if err != nil {
		return false, err
	}

	return true, nil
}
