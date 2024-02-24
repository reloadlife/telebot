// Func: String() for all types
// the reason for this, is that so if anything goes wrong with the String, outputs,
//we can just come here and know where to look.

package telebot

import (
	"encoding/json"
	"fmt"
)

type Stringer interface {
	String() string
}

func (d *MessageThreadID) String() string {
	return fmt.Sprintf("%d", d)
}

func (d *DiceEmoji) String() string {
	return string(*d)
}

func (d *StickerEmoji) String() string {
	return string(*d)
}

func (p *PollType) String() string {
	return string(*p)
}

func (p *ParseMode) String() string {
	return string(*p)
}

func (u *Update) String() string {
	indented, _ := json.MarshalIndent(u, "", "  ")
	return fmt.Sprintf("%s{%d}\n%s\n", u.Type(), u.ID, indented)
}

func (u *User) String() string {
	if u == nil {
		return "User{nil}"
	}
	indented, _ := json.MarshalIndent(u, "", "  ")
	isBot := ""
	if u.IsBot {
		isBot = "(Bot)"
	}
	return fmt.Sprintf("%s{%sID: %d, User: @%v}\n%s\n", u.ReflectType(), isBot, u.ID, u.Username, indented)
}

func (u *MaybeInaccessibleMessage) String() string {
	if u.IsAccessible() {
		return u.AccessibleMessage.String()
	}
	if u.InaccessibleMessage != nil {
		return u.InaccessibleMessage.String()
	}
	return "MaybeInaccessibleMessage{nil}"
}

func (u *InaccessibleMessage) String() string {
	return fmt.Sprintf("%s{ID: %d}", u.ReflectType(), u.ID)
}

func (u *AccessibleMessage) String() string {
	indented, _ := json.MarshalIndent(u, "", "  ")
	return fmt.Sprintf("%s{ID: %d, %s %s}\n%s\n", u.ReflectType(), u.ID, u.Chat, u.From, indented)
}

func (c *Chat) String() string {
	indented, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("%s(%s){ID: %d, Title: %v}\n%s\n", c.ReflectType(), c.ChatType, c.ID, c.Title, indented)
}

func (r *MessageReactionCountUpdated) String() string {
	indented, _ := json.MarshalIndent(r, "", "  ")
	return fmt.Sprintf("%s{ID: %d}\n%s\n", r.ReflectType(), r.ID, indented)
}

func (r *MessageReaction) String() string {
	indented, _ := json.MarshalIndent(r, "", "  ")
	return fmt.Sprintf("%s{ID: %d}\n%s\n", r.ReflectType(), r.ID, indented)
}

func (c *InlineQuery) String() string {
	indented, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("%s{ID: %s, From: %s, Query: %s}\n%s\n", c.ReflectType(), c.ID, c.From.String(), c.Query, indented)
}

func (c *ChosenInlineResult) String() string {
	indented, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("%s{ID: %s, From: %s, Query: %s}\n%s\n", c.ReflectType(), c.ID, c.From.String(), c.Query, indented)
}

func (c *Callback) String() string {
	indented, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("%s{ID: %s, From: %s, data: %s (%s)}\n%s\n", c.ReflectType(), c.ID, c.Sender.String(), c.Data, c.Unique, indented)
}

func (c *ShippingQuery) String() string {
	indented, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("%s{ID: %s, From: %s}\n%s\n", c.ReflectType(), c.ID, c.From.String(), indented)
}

func (c *PreCheckoutQuery) String() string {
	indented, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("%s{ID: %s, From: %s}\n%s\n", c.ReflectType(), c.ID, c.From.String(), indented)
}

func (p *Poll) String() string {
	indented, _ := json.MarshalIndent(p, "", "  ")
	return fmt.Sprintf("%s{ID: %s}\n%s\n", p.ReflectType(), p.ID, indented)
}

func (c *PollAnswer) String() string {
	indented, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("%s{ID: %s}\n%s\n", c.ReflectType(), c.ID, indented)
}

func (c *ChatJoinRequest) String() string {
	indented, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("%s{}\n%s\n", c.ReflectType(), indented)
}

func (c *ChatBoostUpdated) String() string {
	indented, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("%s{}\n%s\n", c.ReflectType(), indented)
}

func (c *ChatBoostRemoved) String() string {
	indented, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("%s{}\n%s\n", c.ReflectType(), indented)
}

func (i *InlineKeyboardButton) String() string {
	return fmt.Sprintf("InlineButton{%s %v}", i.Text, i.CallbackData)
}

func (i *KeyboardButton) String() string { return fmt.Sprintf("Keyboard{%si.Text}", i.Text) }

func (c *ChatPhoto) String() string {
	return fmt.Sprintf("%s{\nSmallFileID: %s, \nSmallUniqueID: %s, \nBigFileID: %s, \nBigUniqueID: %s\n}",
		c.ReflectType(),
		c.SmallFileID, c.SmallUniqueID, c.BigFileID, c.BigUniqueID)
}

func (c *ChatLocation) String() string {
	return fmt.Sprintf("%s{Address: %s}",
		c.ReflectType(),
		c.Address,
	)
}

func (c *ChatPermissions) String() string {
	return fmt.Sprintf("%s{}",
		c.ReflectType(),
	)
}

func (p *ChatMemberPermission) String() string {
	switch *p {
	case IsAnonymous:
		return "IsAnonymous"
	case CanManageChat:
		return "CanManageChat"
	case CanDeleteMessages:
		return "CanDeleteMessages"
	case CanManageVideoChats:
		return "CanManageVideoChats"
	case CanRestrictMembers:
		return "CanRestrictMembers"
	case CanPromoteMembers:
		return "CanPromoteMembers"
	case CanChangeInfo:
		return "CanChangeInfo"
	case CanInviteUsers:
		return "CanInviteUsers"
	case CanPostMessages:
		return "CanPostMessages"
	case CanEditMessages:
		return "CanEditMessages"
	case CanPinMessages:
		return "CanPinMessages"
	case CanPostStories:
		return "CanPostStories"
	case CanEditStories:
		return "CanEditStories"
	case CanDeleteStories:
		return "CanDeleteStories"
	case CanManageTopics:
		return "CanManageTopics"
	default:
		return fmt.Sprintf("Unknown ChatMemberPermission: %d", p)
	}
}

func (c *CallbackGame) String() string {
	return fmt.Sprintf("%s{}",
		c.ReflectType(),
	)
}

func (c *LoginURL) String() string {
	return fmt.Sprintf("%s{}",
		c.ReflectType(),
	)
}

func (c *ChatMember) String() string {
	return fmt.Sprintf("%s{%s %s}",
		c.ReflectType(),
		c.User.String(),
		c.Status,
	)
}

func (c *ReactionType) String() string {
	return fmt.Sprintf("%s{%s %s}", c.ReflectType(), c.ReactionType, c.Emoji)
}

func (c *ReactionCount) String() string {
	return fmt.Sprintf("%s{%s %d}", c.ReflectType(), c.ReactionType, c.Count)
}

func (c *ChatMemberUpdated) String() string {
	return fmt.Sprintf("%s{%s %s}",
		c.ReflectType(),
		c.Chat.String(),
		c.From.String(),
	)
}

func (c *ChatInviteLink) String() string {
	return fmt.Sprintf("%s{ID: %s, Creator: %s, IsPrimary: %v}",
		c.ReflectType(),
		c.InviteLink,
		c.Creator,
		c.IsPrimary,
	)
}

func (r *Rights) String() string {
	return fmt.Sprintf("%s{}",
		r.ReflectType(),
	)
}

func (m *MessageOrigin) String() string {
	return fmt.Sprintf("%s{ID: %d, %s}",
		m.ReflectType(),
		m.ID,
		m.Type(),
	)
}

func (u *ExternalReplyInfo) String() string {
	return fmt.Sprintf("%s{ID: %s}",
		u.ReflectType(),
		u.Origin.String(),
	)
}

func (m *ForceReplyMarkup) String() string {
	return fmt.Sprintf("%s{ %p, %v, %v }", m.ReflectType(),
		m.InputFieldPlaceholder,
		m.Selective,
		m.ForceReply)
}
func (m *ReplyKeyboardRemove) String() string {
	return fmt.Sprintf("%s{ %v }", m.ReflectType(), m.Selective)
}
func (m *InlineKeyboardMarkup) String() string {
	return fmt.Sprintf("%s{ rows: %d }", m.ReflectType(), len(m.InlineKeyboard))
}
func (m *ReplyKeyboardMarkup) String() string {
	return fmt.Sprintf("%s{ rows: %d,  %p, %v, %v, %v, %v }", m.ReflectType(),
		len(m.Keyboard),
		m.InputFieldPlaceholder,
		m.Selective,
		m.OneTimeKeyboard,
		m.IsPersistent,
		m.ResizeKeyboard,
	)
}

func (i *SwitchInlineQueryChosenChat) String() string {
	return fmt.Sprintf("%s{ %p }", i.ReflectType(), i.Query)
}
func (i *KeyboardButtonRequestChat) String() string {
	return fmt.Sprintf("%s{ %d }", i.ReflectType(), i.RequestID)
}
func (i *KeyboardButtonRequestUsers) String() string {
	return fmt.Sprintf("%s{ %d }", i.ReflectType(), i.RequestID)
}
func (i *KeyboardButtonPollType) String() string {
	return fmt.Sprintf("%s{ %s }", i.ReflectType(), i.Type())
}
func (f *File) String() string {
	if f == nil {
		return "File{nil}"
	}
	return fmt.Sprintf("%s{ID: %s, Size: %dMB}", f.ReflectType(), f.UniqueID, f.FileSize/1024/1024)
}

func (f *TextQuote) String() string { return fmt.Sprintf("%s{%s}", f.ReflectType(), f.Text) }

func (f *Entity) String() string {
	return fmt.Sprintf("%s{%s %s %s %p}", f.ReflectType(), f.Type(), f.URL, f.User, f.Language)
}

func (f *LinkPreviewOptions) String() string { return fmt.Sprintf("%s{%p}", f.ReflectType(), f.URL) }
func (a *Animation) String() string          { return fmt.Sprintf("%s{%s}", a.ReflectType(), a.File()) }
func (a *Audio) String() string              { return fmt.Sprintf("%s{%s}", a.ReflectType(), a.File()) }
func (d *Document) String() string           { return fmt.Sprintf("%s{%s}", d.ReflectType(), d.File()) }
func (p *PhotoSize) String() string          { return fmt.Sprintf("%s{%s}", p.ReflectType(), p.File()) }
func (f *Sticker) String() string            { return fmt.Sprintf("%s{%s}", f.ReflectType(), f.File()) }
func (v *Video) String() string              { return fmt.Sprintf("%s{%s}", v.ReflectType(), v.File()) }
func (v *Voice) String() string              { return fmt.Sprintf("%s{%s}", v.ReflectType(), v.File()) }
func (v *VideoNote) String() string          { return fmt.Sprintf("%s{%s}", v.ReflectType(), v.File()) }

func (f *Story) String() string {
	return fmt.Sprintf("%s{%d}", f.ReflectType(), f.ID)
}

func (c *Contact) String() string {
	return fmt.Sprintf("%s{%s %d %p %s%p}", c.ReflectType(), c.PhoneNumber, c.UserID, c.VCard, c.FirstName, c.LastName)
}

func (d *Dice) String() string {
	return fmt.Sprintf("%s{%s %d}", d.ReflectType(), d.Emoji, d.Value)
}

func (f *Game) String() string {
	return fmt.Sprintf("%s{%s %s}", f.ReflectType(), f.Title, f.Description)
}

func (f *GameHighScore) String() string {
	return fmt.Sprintf("%s{%s %d}", f.ReflectType(), f.User.String(), f.Score)
}

func (p *PollOption) String() string {
	return fmt.Sprintf("%s{%s %d}", p.ReflectType(), p.Text, p.VoterCount)
}

func (v *Venue) String() string {
	return fmt.Sprintf("%s{%s %s %s}", v.ReflectType(), v.Location.String(), v.Title, v.Address)
}

func (l *Location) String() string {
	return fmt.Sprintf("%s{%f %f}", l.ReflectType(), l.Latitude, l.Longitude)
}

func (f *Invoice) String() string {
	return fmt.Sprintf("%s{%s %s %s%d}", f.ReflectType(), f.Title, f.Description, f.Currency, f.TotalAmount)
}

func (f *SuccessfulPayment) String() string {
	return fmt.Sprintf("%s{%s%d %s %s}", f.ReflectType(), f.Currency, f.TotalAmount, f.InvoicePayload, f.ShippingOptionID)
}

func (f *OrderInfo) String() string {
	return fmt.Sprintf("%s{%s %s %s %s}", f.ReflectType(), f.Name, f.PhoneNumber, f.Email, f.ShippingAddress.String())
}

func (f *ShippingAddress) String() string {
	return fmt.Sprintf("%s{%s %s %s %s %s}", f.ReflectType(), f.CountryCode, f.State, f.City, f.StreetLine1, f.StreetLine2)
}

func (f *LabeledPrice) String() string {
	return fmt.Sprintf("%s{%s %d}", f.ReflectType(), f.Label, f.Amount)
}

func (f *ProximityAlertTriggered) String() string {
	return fmt.Sprintf("%s{%s %s}", f.ReflectType(), f.Traveler.String(), f.Watcher.String())
}

func (f *UsersShared) String() string {
	return fmt.Sprintf("%s{%d %v}", f.ReflectType(), f.RequestID, f.UserIDs)
}

func (f *ChatShared) String() string {
	return fmt.Sprintf("%s{%d %d}", f.ReflectType(), f.ChatID, f.RequestID)
}

func (f *AutoDeleteTimerChanged) String() string {
	return fmt.Sprintf("%s{%d}", f.ReflectType(), f.AutoDeleteTime)
}

func (f *PassportData) String() string {
	return fmt.Sprintf("%s{}", f.ReflectType())
}

func (f *EncryptedPassportElement) String() string {
	return fmt.Sprintf("%s{%s %s}", f.ReflectType(), f.Type(), f.Data)
}

func (f *EncryptedCredentials) String() string {
	return fmt.Sprintf("%s{}", f.ReflectType())
}

func (f *PassportFile) String() string {
	return fmt.Sprintf("%s{%s}", f.ReflectType(), f.File())
}

func (f *WriteAccessAllowed) String() string {
	return fmt.Sprintf("%s{}", f.ReflectType())
}

func (f *ForumTopicCreated) String() string {
	return fmt.Sprintf("%s{%s}", f.ReflectType(), f.Name)
}

func (f *ForumTopicEdited) String() string {
	return fmt.Sprintf("%s{%s}", f.ReflectType(), f.Name)
}

func (f *ForumTopicClosed) String() string {
	return fmt.Sprintf("%s{}", f.ReflectType())
}

func (f *ForumTopicReopened) String() string {
	return fmt.Sprintf("%s{}", f.ReflectType())
}

func (f *GeneralForumTopicHidden) String() string {
	return fmt.Sprintf("%s{}", f.ReflectType())
}

func (f *GeneralForumTopicUnhidden) String() string {
	return fmt.Sprintf("%s{}", f.ReflectType())
}

func (f *GiveawayCreated) String() string {
	return fmt.Sprintf("%s{}", f.ReflectType())
}

func (f *GiveawayCompleted) String() string {
	return fmt.Sprintf("%s{}", f.ReflectType())
}

func (f *GiveawayWinners) String() string {
	return fmt.Sprintf("%s{}", f.ReflectType())
}

func (f *Giveaway) String() string {
	return fmt.Sprintf("%s{}", f.ReflectType())
}

func (f *VideoChatScheduled) String() string {
	return fmt.Sprintf("%s{}", f.ReflectType())
}

func (f *VideoChatEnded) String() string {
	return fmt.Sprintf("%s{}", f.ReflectType())
}

func (f *VideoChatStarted) String() string {
	return fmt.Sprintf("%s{}", f.ReflectType())
}

func (f *VideoChatParticipantsInvited) String() string {
	return fmt.Sprintf("%s{}", f.ReflectType())
}

func (f *WebAppData) String() string {
	return fmt.Sprintf("%s{}", f.ReflectType())
}

func (f *SentWebAppMessage) String() string {
	return fmt.Sprintf("%s{}", f.ReflectType())
}

func (f *WebAppInfo) String() string {
	return fmt.Sprintf("%s{%s}", f.ReflectType(), f.URL)
}

func (f *BotCommand) String() string {
	return fmt.Sprintf("%s{ /%s %s }", f.ReflectType(), f.Command, f.Description)
}

func (f *BotCommandScope) String() string {
	return fmt.Sprintf("%s{ %s %d %d}", f.ReflectType(), f.Type(), f.ChatID, f.UserID)
}

func (u *UserProfilePhotos) String() string {
	return fmt.Sprintf("%s{ %d %d }", u.ReflectType(), u.TotalCount, len(u.Photos))
}

func (f *UserChatBoosts) String() string {
	return fmt.Sprintf("%s{ %d }", f.ReflectType(), len(f.Boosts))
}

func (f *StickerSet) String() string {
	return fmt.Sprintf("%s{ %s %s }", f.ReflectType(), f.Name, f.Title)
}

func (r *ReplyParameters) String() string {
	return fmt.Sprintf("%s{ %d -> %s }", r.ReflectType(), r.MessageID, r.ChatID)
}

func (f *QueryResult) String() string {
	return fmt.Sprintf("%s{ %s %s }", f.ReflectType(), f.Type(), f.ID)
}

func (f *PassportElementError) String() string {
	return fmt.Sprintf("%s{ }", f.ReflectType())
}

func (f *MenuButton) String() string {
	return fmt.Sprintf("%s{ %s %p }", f.ReflectType(), f.Type(), f.Text)
}

func (f *MaskPosition) String() string {
	return fmt.Sprintf("%s{ }", f.ReflectType())
}

func (f *ForumTopic) String() string {
	return fmt.Sprintf("%s{ %s }", f.ReflectType(), f.Name)
}

func (f *ChatBoost) String() string {
	return fmt.Sprintf("%s{ %s }", f.ReflectType(), f.BoostID)
}

func (f *ChatBoostSource) String() string {
	return fmt.Sprintf("%s{ %s }", f.ReflectType(), f.Source)
}

func (f *InputContactMessageContent) String() string {
	return fmt.Sprintf("%s{ %s }", f.ReflectType(), f.PhoneNumber)
}

func (f *InputLocationMessageContent) String() string {
	return fmt.Sprintf("%s{ %v %v }", f.ReflectType(), f.Latitude, f.Longitude)
}

func (f *InputTextMessageContent) String() string {
	return fmt.Sprintf("%s{ %s }", f.ReflectType(), f.MessageText)
}

func (f *InputVenueMessageContent) String() string {
	return fmt.Sprintf("%s{ %s }", f.ReflectType(), f.Title)
}

func (f *InputInvoiceMessageContent) String() string {
	return fmt.Sprintf("%s{ %s }", f.ReflectType(), f.Title)
}

func (f *InputSticker) String() string {
	return fmt.Sprintf("%s{ %s }", f.ReflectType(), f.Sticker.String())
}

func (m *InputMedia) String() string {
	return fmt.Sprintf("%s{ %s %s }", m.ReflectType(), m.Type(), m.Media.String())
}

func (f *ChatBoostAdded) String() string {
	return fmt.Sprintf("%s{ %d }", f.ReflectType(), f.BoostCount)
}
