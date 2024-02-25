package telebot

import (
	"encoding/json"
	"errors"
)

type JSONer interface {
	MarshalJSON() ([]byte, error)
	UnmarshalJSON([]byte) error
}

func (u *Update) MarshalJSON() ([]byte, error) {
	return json.Marshal(*u)
}
func (u *AccessibleMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(*u)
}
func (u *InaccessibleMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(*u)
}
func (u *User) MarshalJSON() ([]byte, error) {
	return json.Marshal(*u)
}
func (c *Chat) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)
}
func (r *MessageReaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
func (r *MessageReactionCountUpdated) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
func (c *InlineQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)
}
func (c *ChosenInlineResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)
}
func (c *Callback) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)
}
func (sh *ShippingQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(*sh)
}
func (c *PreCheckoutQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)
}
func (p *Poll) MarshalJSON() ([]byte, error) {
	return json.Marshal(*p)
}
func (c *PollAnswer) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)
}
func (c *ChatMemberUpdated) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)
}
func (c *ChatJoinRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)
}
func (c *ChatBoostUpdated) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)
}
func (c *ChatBoostRemoved) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)
}
func (i *InlineKeyboardButton) MarshalJSON() ([]byte, error) { return json.Marshal(*i) }
func (i *KeyboardButton) MarshalJSON() ([]byte, error)       { return json.Marshal(*i) }
func (c *ChatPhoto) MarshalJSON() ([]byte, error)            { return json.Marshal(*c) }
func (c *ChatLocation) MarshalJSON() ([]byte, error)         { return json.Marshal(*c) }
func (c *ChatPermissions) MarshalJSON() ([]byte, error)      { return json.Marshal(*c) }
func (c *CallbackGame) MarshalJSON() ([]byte, error)         { return json.Marshal(*c) }
func (c *LoginURL) MarshalJSON() ([]byte, error)             { return json.Marshal(*c) }
func (c *ChatMember) MarshalJSON() ([]byte, error)           { return json.Marshal(*c) }
func (c *ReactionType) MarshalJSON() ([]byte, error)         { return json.Marshal(*c) }
func (c *ReactionCount) MarshalJSON() ([]byte, error)        { return json.Marshal(*c) }
func (c *ChatInviteLink) MarshalJSON() ([]byte, error)       { return json.Marshal(*c) }
func (r *Rights) MarshalJSON() ([]byte, error)               { return json.Marshal(*r) }
func (m *MessageOrigin) MarshalJSON() ([]byte, error)        { return json.Marshal(*m) }
func (u *ExternalReplyInfo) MarshalJSON() ([]byte, error)    { return json.Marshal(*u) }
func (m *ForceReplyMarkup) MarshalJSON() ([]byte, error)     { return json.Marshal(*m) }
func (m *ReplyKeyboardRemove) MarshalJSON() ([]byte, error)  { return json.Marshal(*m) }
func (m *InlineKeyboardMarkup) MarshalJSON() ([]byte, error) { return json.Marshal(*m) }
func (m *ReplyKeyboardMarkup) MarshalJSON() ([]byte, error)  { return json.Marshal(*m) }

func (i *SwitchInlineQueryChosenChat) MarshalJSON() ([]byte, error) { return json.Marshal(*i) }
func (i *KeyboardButtonRequestChat) MarshalJSON() ([]byte, error)   { return json.Marshal(*i) }
func (i *KeyboardButtonRequestUsers) MarshalJSON() ([]byte, error)  { return json.Marshal(*i) }
func (i *KeyboardButtonPollType) MarshalJSON() ([]byte, error)      { return json.Marshal(*i) }

func (c *ChatBoostAdded) MarshalJSON() ([]byte, error)               { return json.Marshal(*c) }
func (c *ShippingAddress) MarshalJSON() ([]byte, error)              { return json.Marshal(*c) }
func (c *OrderInfo) MarshalJSON() ([]byte, error)                    { return json.Marshal(*c) }
func (c *LabeledPrice) MarshalJSON() ([]byte, error)                 { return json.Marshal(*c) }
func (i *Invoice) MarshalJSON() ([]byte, error)                      { return json.Marshal(*i) }
func (c *SuccessfulPayment) MarshalJSON() ([]byte, error)            { return json.Marshal(*c) }
func (c *PassportData) MarshalJSON() ([]byte, error)                 { return json.Marshal(*c) }
func (c *EncryptedPassportElement) MarshalJSON() ([]byte, error)     { return json.Marshal(*c) }
func (c *EncryptedCredentials) MarshalJSON() ([]byte, error)         { return json.Marshal(*c) }
func (c *PassportFile) MarshalJSON() ([]byte, error)                 { return json.Marshal(*c) }
func (c *PassportElementError) MarshalJSON() ([]byte, error)         { return json.Marshal(*c) }
func (s *InputSticker) MarshalJSON() ([]byte, error)                 { return json.Marshal(*s) }
func (c *InputInvoiceMessageContent) MarshalJSON() ([]byte, error)   { return json.Marshal(*c) }
func (c *InputLocationMessageContent) MarshalJSON() ([]byte, error)  { return json.Marshal(*c) }
func (c *InputVenueMessageContent) MarshalJSON() ([]byte, error)     { return json.Marshal(*c) }
func (c *InputContactMessageContent) MarshalJSON() ([]byte, error)   { return json.Marshal(*c) }
func (c *InputTextMessageContent) MarshalJSON() ([]byte, error)      { return json.Marshal(*c) }
func (c *ChatBoostSource) MarshalJSON() ([]byte, error)              { return json.Marshal(*c) }
func (c *ChatBoost) MarshalJSON() ([]byte, error)                    { return json.Marshal(*c) }
func (f *File) MarshalJSON() ([]byte, error)                         { return json.Marshal(*f) }
func (p *PhotoSize) MarshalJSON() ([]byte, error)                    { return json.Marshal(*p) }
func (a *Audio) MarshalJSON() ([]byte, error)                        { return json.Marshal(*a) }
func (d *Document) MarshalJSON() ([]byte, error)                     { return json.Marshal(*d) }
func (v *Video) MarshalJSON() ([]byte, error)                        { return json.Marshal(*v) }
func (a *Animation) MarshalJSON() ([]byte, error)                    { return json.Marshal(*a) }
func (v *Voice) MarshalJSON() ([]byte, error)                        { return json.Marshal(*v) }
func (v *VideoNote) MarshalJSON() ([]byte, error)                    { return json.Marshal(*v) }
func (c *Contact) MarshalJSON() ([]byte, error)                      { return json.Marshal(*c) }
func (l *Location) MarshalJSON() ([]byte, error)                     { return json.Marshal(*l) }
func (v *Venue) MarshalJSON() ([]byte, error)                        { return json.Marshal(*v) }
func (p *PollOption) MarshalJSON() ([]byte, error)                   { return json.Marshal(*p) }
func (d *Dice) MarshalJSON() ([]byte, error)                         { return json.Marshal(*d) }
func (f *Game) MarshalJSON() ([]byte, error)                         { return json.Marshal(*f) }
func (s *Sticker) MarshalJSON() ([]byte, error)                      { return json.Marshal(*s) }
func (f *StickerSet) MarshalJSON() ([]byte, error)                   { return json.Marshal(*f) }
func (f *MaskPosition) MarshalJSON() ([]byte, error)                 { return json.Marshal(*f) }
func (f *ForumTopic) MarshalJSON() ([]byte, error)                   { return json.Marshal(*f) }
func (f *ForumTopicClosed) MarshalJSON() ([]byte, error)             { return json.Marshal(*f) }
func (f *ForumTopicReopened) MarshalJSON() ([]byte, error)           { return json.Marshal(*f) }
func (f *ForumTopicEdited) MarshalJSON() ([]byte, error)             { return json.Marshal(*f) }
func (f *ForumTopicCreated) MarshalJSON() ([]byte, error)            { return json.Marshal(*f) }
func (f *GameHighScore) MarshalJSON() ([]byte, error)                { return json.Marshal(*f) }
func (f *MenuButton) MarshalJSON() ([]byte, error)                   { return json.Marshal(*f) }
func (f *TextQuote) MarshalJSON() ([]byte, error)                    { return json.Marshal(*f) }
func (f *Entity) MarshalJSON() ([]byte, error)                       { return json.Marshal(*f) }
func (f *LinkPreviewOptions) MarshalJSON() ([]byte, error)           { return json.Marshal(*f) }
func (f *Story) MarshalJSON() ([]byte, error)                        { return json.Marshal(*f) }
func (f *AutoDeleteTimerChanged) MarshalJSON() ([]byte, error)       { return json.Marshal(*f) }
func (f *ProximityAlertTriggered) MarshalJSON() ([]byte, error)      { return json.Marshal(*f) }
func (f *UsersShared) MarshalJSON() ([]byte, error)                  { return json.Marshal(*f) }
func (f *ChatShared) MarshalJSON() ([]byte, error)                   { return json.Marshal(*f) }
func (f *WriteAccessAllowed) MarshalJSON() ([]byte, error)           { return json.Marshal(*f) }
func (f *GeneralForumTopicHidden) MarshalJSON() ([]byte, error)      { return json.Marshal(*f) }
func (f *GeneralForumTopicUnhidden) MarshalJSON() ([]byte, error)    { return json.Marshal(*f) }
func (f *GiveawayCreated) MarshalJSON() ([]byte, error)              { return json.Marshal(*f) }
func (f *GiveawayCompleted) MarshalJSON() ([]byte, error)            { return json.Marshal(*f) }
func (f *GiveawayWinners) MarshalJSON() ([]byte, error)              { return json.Marshal(*f) }
func (f *Giveaway) MarshalJSON() ([]byte, error)                     { return json.Marshal(*f) }
func (f *VideoChatScheduled) MarshalJSON() ([]byte, error)           { return json.Marshal(*f) }
func (f *VideoChatEnded) MarshalJSON() ([]byte, error)               { return json.Marshal(*f) }
func (f *VideoChatParticipantsInvited) MarshalJSON() ([]byte, error) { return json.Marshal(*f) }
func (f *VideoChatStarted) MarshalJSON() ([]byte, error)             { return json.Marshal(*f) }
func (w *WebAppData) MarshalJSON() ([]byte, error)                   { return json.Marshal(*w) }
func (w *WebAppInfo) MarshalJSON() ([]byte, error)                   { return json.Marshal(*w) }
func (f *BotCommand) MarshalJSON() ([]byte, error)                   { return json.Marshal(*f) }
func (f *BotCommandScope) MarshalJSON() ([]byte, error)              { return json.Marshal(*f) }
func (u *UserProfilePhotos) MarshalJSON() ([]byte, error)            { return json.Marshal(*u) }
func (f *UserChatBoosts) MarshalJSON() ([]byte, error)               { return json.Marshal(*f) }
func (s *SentWebAppMessage) MarshalJSON() ([]byte, error)            { return json.Marshal(*s) }
func (r *ReplyParameters) MarshalJSON() ([]byte, error)              { return json.Marshal(*r) }
func (q *QueryResult) MarshalJSON() ([]byte, error)                  { return json.Marshal(*q) }
func (m *InputMedia) MarshalJSON() ([]byte, error)                   { return json.Marshal(*m) }

func (u *MaybeInaccessibleMessage) MarshalJSON() ([]byte, error) {
	if u.IsAccessible() {
		return u.AccessibleMessage.MarshalJSON()
	}
	if u.InaccessibleMessage != nil {
		return u.InaccessibleMessage.MarshalJSON()
	}
	return nil, ErrNoMessageToMarshal
}

// /// //// ////// ////// ////// ////// ////// ////// ////// ////// ////// //////

// UnmarshalJSON in MaybeInaccessibleMessage is a custom unmarshaler that
// checks if the message is accessible or not and unmarshals it accordingly.
// It first checks if the message is accessible by checking the date field.
// note from the author: I'm not sure if this is the best way to do it, but it works.
// if you find a better way of doing it that passes the UnitTests, i'd appreciate the PR <3.
func (u *MaybeInaccessibleMessage) UnmarshalJSON(b []byte) error {
	var where struct {
		Date int64 `json:"date"`
	}
	err := json.Unmarshal(b, &where)
	if err != nil {
		return err
	}
	if where.Date == 0 {
		InM := &InaccessibleMessage{}
		err := InM.UnmarshalJSON(b)
		if err == nil {
			u.AccessibleMessage = nil
			u.InaccessibleMessage = InM
		}
		return err
	}
	AcM := &AccessibleMessage{}
	err = AcM.UnmarshalJSON(b)
	if err == nil {
		u.InaccessibleMessage = nil
		u.AccessibleMessage = AcM
	}
	return err
}

func (u *Update) UnmarshalJSON(b []byte) error {
	type Alias Update
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (u *AccessibleMessage) UnmarshalJSON(b []byte) error {
	type Alias AccessibleMessage
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (u *InaccessibleMessage) UnmarshalJSON(b []byte) error {
	type Alias InaccessibleMessage
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (u *User) UnmarshalJSON(b []byte) error {
	type Alias User
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (c *Chat) UnmarshalJSON(b []byte) error {
	type Alias Chat
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (r *MessageReaction) UnmarshalJSON(b []byte) error {
	type Alias MessageReaction
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (r *MessageReactionCountUpdated) UnmarshalJSON(b []byte) error {
	type Alias MessageReactionCountUpdated
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (c *InlineQuery) UnmarshalJSON(b []byte) error {
	type Alias InlineQuery
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (c *ChosenInlineResult) UnmarshalJSON(b []byte) error {
	type Alias ChosenInlineResult
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (c *Callback) UnmarshalJSON(b []byte) error {
	type Alias Callback
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (sh *ShippingQuery) UnmarshalJSON(b []byte) error {
	type Alias ShippingQuery
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(sh),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (c *PreCheckoutQuery) UnmarshalJSON(b []byte) error {
	type Alias PreCheckoutQuery
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (p *Poll) UnmarshalJSON(b []byte) error {
	type Alias Poll
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (c *PollAnswer) UnmarshalJSON(b []byte) error {
	type Alias PollAnswer
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (c *ChatMemberUpdated) UnmarshalJSON(b []byte) error {
	type Alias ChatMemberUpdated
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (c *ChatJoinRequest) UnmarshalJSON(b []byte) error {
	type Alias ChatJoinRequest
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (c *ChatBoostUpdated) UnmarshalJSON(b []byte) error {
	type Alias ChatBoostUpdated
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (c *ChatBoostRemoved) UnmarshalJSON(b []byte) error {
	type Alias ChatBoostRemoved
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(b, &aux); err != nil {
		return err
	}
	return nil
}
func (c *ChatPhoto) UnmarshalJSON(data []byte) error {
	type Alias ChatPhoto
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *ChatLocation) UnmarshalJSON(data []byte) error {
	type Alias ChatLocation
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *ChatPermissions) UnmarshalJSON(data []byte) error {
	type Alias ChatPermissions
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *CallbackGame) UnmarshalJSON(data []byte) error {
	type Alias CallbackGame
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *LoginURL) UnmarshalJSON(data []byte) error {
	type Alias LoginURL
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *ChatMember) UnmarshalJSON(data []byte) error {
	type Alias ChatMember
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *ReactionType) UnmarshalJSON(data []byte) error {
	type Alias ReactionType
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *ReactionCount) UnmarshalJSON(data []byte) error {
	type Alias ReactionCount
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *ChatInviteLink) UnmarshalJSON(data []byte) error {
	type Alias ChatInviteLink
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (r *Rights) UnmarshalJSON(data []byte) error {
	type Alias Rights
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (m *MessageOrigin) UnmarshalJSON(data []byte) error {
	type Alias MessageOrigin
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (u *ExternalReplyInfo) UnmarshalJSON(data []byte) error {
	type Alias ExternalReplyInfo
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (m *ForceReplyMarkup) UnmarshalJSON(data []byte) error {
	type Alias ForceReplyMarkup
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (m *ReplyKeyboardRemove) UnmarshalJSON(data []byte) error {
	type Alias ReplyKeyboardRemove
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (i *SwitchInlineQueryChosenChat) UnmarshalJSON(data []byte) error {
	type Alias SwitchInlineQueryChosenChat
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(i),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (i *KeyboardButtonRequestChat) UnmarshalJSON(data []byte) error {
	type Alias KeyboardButtonRequestChat
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(i),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (i *KeyboardButtonRequestUsers) UnmarshalJSON(data []byte) error {
	type Alias KeyboardButtonRequestUsers
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(i),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (i *KeyboardButtonPollType) UnmarshalJSON(data []byte) error {
	type Alias KeyboardButtonPollType
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(i),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *ChatBoostAdded) UnmarshalJSON(data []byte) error {
	type Alias ChatBoostAdded
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *ShippingAddress) UnmarshalJSON(data []byte) error {
	type Alias ShippingAddress
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *OrderInfo) UnmarshalJSON(data []byte) error {
	type Alias OrderInfo
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *LabeledPrice) UnmarshalJSON(data []byte) error {
	type Alias LabeledPrice
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (i *Invoice) UnmarshalJSON(data []byte) error {
	type Alias Invoice
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(i),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *SuccessfulPayment) UnmarshalJSON(data []byte) error {
	type Alias SuccessfulPayment
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *PassportData) UnmarshalJSON(data []byte) error {
	type Alias PassportData
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *EncryptedPassportElement) UnmarshalJSON(data []byte) error {
	type Alias EncryptedPassportElement
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *EncryptedCredentials) UnmarshalJSON(data []byte) error {
	type Alias EncryptedCredentials
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *PassportFile) UnmarshalJSON(data []byte) error {
	type Alias PassportFile
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *PassportElementError) UnmarshalJSON(data []byte) error {
	type Alias PassportElementError
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (s *InputSticker) UnmarshalJSON(data []byte) error {
	type Alias InputSticker
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *InputInvoiceMessageContent) UnmarshalJSON(data []byte) error {
	type Alias InputInvoiceMessageContent
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *InputLocationMessageContent) UnmarshalJSON(data []byte) error {
	type Alias InputLocationMessageContent
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *InputVenueMessageContent) UnmarshalJSON(data []byte) error {
	type Alias InputVenueMessageContent
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *InputContactMessageContent) UnmarshalJSON(data []byte) error {
	type Alias InputContactMessageContent
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *InputTextMessageContent) UnmarshalJSON(data []byte) error {
	type Alias InputTextMessageContent
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *ChatBoostSource) UnmarshalJSON(data []byte) error {
	type Alias ChatBoostSource
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *ChatBoost) UnmarshalJSON(data []byte) error {
	type Alias ChatBoost
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *File) UnmarshalJSON(data []byte) error {
	type Alias File
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (p *PhotoSize) UnmarshalJSON(data []byte) error {
	type Alias PhotoSize
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (a *Audio) UnmarshalJSON(data []byte) error {
	type Alias Audio
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (d *Document) UnmarshalJSON(data []byte) error {
	type Alias Document
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(d),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (v *Video) UnmarshalJSON(data []byte) error {
	type Alias Video
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(v),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (a *Animation) UnmarshalJSON(data []byte) error {
	type Alias Animation
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(a),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (v *Voice) UnmarshalJSON(data []byte) error {
	type Alias Voice
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(v),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (v *VideoNote) UnmarshalJSON(data []byte) error {
	type Alias VideoNote
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(v),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (c *Contact) UnmarshalJSON(data []byte) error {
	type Alias Contact
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (l *Location) UnmarshalJSON(data []byte) error {
	type Alias Location
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(l),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (v *Venue) UnmarshalJSON(data []byte) error {
	type Alias Venue
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(v),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (p *PollOption) UnmarshalJSON(data []byte) error {
	type Alias PollOption
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(p),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (d *Dice) UnmarshalJSON(data []byte) error {
	type Alias Dice
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(d),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *Game) UnmarshalJSON(data []byte) error {
	type Alias Game
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (s *Sticker) UnmarshalJSON(data []byte) error {
	type Alias Sticker
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *StickerSet) UnmarshalJSON(data []byte) error {
	type Alias StickerSet
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *MaskPosition) UnmarshalJSON(data []byte) error {
	type Alias MaskPosition
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *ForumTopic) UnmarshalJSON(data []byte) error {
	type Alias ForumTopic
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *ForumTopicClosed) UnmarshalJSON(data []byte) error {
	type Alias ForumTopicClosed
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *ForumTopicReopened) UnmarshalJSON(data []byte) error {
	type Alias ForumTopicReopened
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *ForumTopicEdited) UnmarshalJSON(data []byte) error {
	type Alias ForumTopicEdited
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *ForumTopicCreated) UnmarshalJSON(data []byte) error {
	type Alias ForumTopicCreated
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *GameHighScore) UnmarshalJSON(data []byte) error {
	type Alias GameHighScore
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *MenuButton) UnmarshalJSON(data []byte) error {
	type Alias MenuButton
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *TextQuote) UnmarshalJSON(data []byte) error {
	type Alias TextQuote
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *Entity) UnmarshalJSON(data []byte) error {
	type Alias Entity
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *LinkPreviewOptions) UnmarshalJSON(data []byte) error {
	type Alias LinkPreviewOptions
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *Story) UnmarshalJSON(data []byte) error {
	type Alias Story
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *AutoDeleteTimerChanged) UnmarshalJSON(data []byte) error {
	type Alias AutoDeleteTimerChanged
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *ProximityAlertTriggered) UnmarshalJSON(data []byte) error {
	type Alias ProximityAlertTriggered
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *UsersShared) UnmarshalJSON(data []byte) error {
	type Alias UsersShared
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *ChatShared) UnmarshalJSON(data []byte) error {
	type Alias ChatShared
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *WriteAccessAllowed) UnmarshalJSON(data []byte) error {
	type Alias WriteAccessAllowed
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *GeneralForumTopicHidden) UnmarshalJSON(data []byte) error {
	type Alias GeneralForumTopicHidden
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *GeneralForumTopicUnhidden) UnmarshalJSON(data []byte) error {
	type Alias GeneralForumTopicUnhidden
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *GiveawayCreated) UnmarshalJSON(data []byte) error {
	type Alias GiveawayCreated
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *GiveawayCompleted) UnmarshalJSON(data []byte) error {
	type Alias GiveawayCompleted
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *GiveawayWinners) UnmarshalJSON(data []byte) error {
	type Alias GiveawayWinners
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *Giveaway) UnmarshalJSON(data []byte) error {
	type Alias Giveaway
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *VideoChatScheduled) UnmarshalJSON(data []byte) error {
	type Alias VideoChatScheduled
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *VideoChatEnded) UnmarshalJSON(data []byte) error {
	type Alias VideoChatEnded
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *VideoChatParticipantsInvited) UnmarshalJSON(data []byte) error {
	type Alias VideoChatParticipantsInvited
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *VideoChatStarted) UnmarshalJSON(data []byte) error {
	type Alias VideoChatStarted
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (w *WebAppData) UnmarshalJSON(data []byte) error {
	type Alias WebAppData
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(w),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (w *WebAppInfo) UnmarshalJSON(data []byte) error {
	type Alias WebAppInfo
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(w),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *BotCommand) UnmarshalJSON(data []byte) error {
	type Alias BotCommand
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *BotCommandScope) UnmarshalJSON(data []byte) error {
	type Alias BotCommandScope
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (u *UserProfilePhotos) UnmarshalJSON(data []byte) error {
	type Alias UserProfilePhotos
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (f *UserChatBoosts) UnmarshalJSON(data []byte) error {
	type Alias UserChatBoosts
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(f),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (s *SentWebAppMessage) UnmarshalJSON(data []byte) error {
	type Alias SentWebAppMessage
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(s),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (r *ReplyParameters) UnmarshalJSON(data []byte) error {
	type Alias ReplyParameters
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(r),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (q *QueryResult) UnmarshalJSON(data []byte) error {
	type Alias QueryResult
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(q),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}
func (m *InputMedia) UnmarshalJSON(data []byte) error {
	type Alias InputMedia
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(m),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

///// ReplyMarkup Custom UnmarshalJSON / MarshalJSON.

// UnmarshalJSON implements json.Unmarshaler.
// It supports both InlineKeyboardButton and KeyboardButton.
// It is a bit complicated to unmarshal into interface types, believe me.
// this thing works perfectly for our goal, if you have a better idea, a PR is appreciated.
func (r *Row) UnmarshalJSON(data []byte) error {
	var buttons []json.RawMessage
	if err := json.Unmarshal(data, &buttons); err != nil {
		return err
	}

	butts := make([]Button, 0, len(buttons))

	for _, buttonData := range buttons {
		var rawButton struct {
			URL                          *string                      `json:"url,omitempty"`
			CallbackData                 *string                      `json:"callback_data,omitempty"`
			LoginURL                     *LoginURL                    `json:"login_url,omitempty"`
			SwitchInlineQuery            *string                      `json:"switch_inline_query,omitempty"`
			SwitchInlineQueryCurrentChat *string                      `json:"switch_inline_query_current_chat,omitempty"`
			SwitchInlineQueryChosenChat  *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`
			CallbackGame                 *CallbackGame                `json:"callback_game,omitempty"`
			Pay                          *bool                        `json:"pay,omitempty"`

			RequestUsers    *KeyboardButtonRequestUsers `json:"request_users,omitempty"`
			RequestChat     *KeyboardButtonRequestChat  `json:"request_chat,omitempty"`
			RequestContact  *bool                       `json:"request_contact,omitempty"`
			RequestLocation *bool                       `json:"request_location,omitempty"`
			RequestPoll     *KeyboardButtonPollType     `json:"request_poll,omitempty"`

			Text   string      `json:"text"`
			WebApp *WebAppInfo `json:"web_app,omitempty"`
		}

		if err := json.Unmarshal(buttonData, &rawButton); err != nil {
			return err
		}

		switch {
		case rawButton.URL != nil || rawButton.LoginURL != nil || rawButton.CallbackData != nil ||
			rawButton.SwitchInlineQuery != nil || rawButton.SwitchInlineQueryCurrentChat != nil ||
			rawButton.SwitchInlineQueryChosenChat != nil || rawButton.CallbackGame != nil || rawButton.Pay != nil:
			butts = append(butts, &InlineKeyboardButton{
				Text:                         rawButton.Text,
				URL:                          rawButton.URL,
				CallbackData:                 rawButton.CallbackData,
				WebApp:                       rawButton.WebApp,
				LoginURL:                     rawButton.LoginURL,
				SwitchInlineQuery:            rawButton.SwitchInlineQuery,
				SwitchInlineQueryCurrentChat: rawButton.SwitchInlineQueryCurrentChat,
				SwitchInlineQueryChosenChat:  rawButton.SwitchInlineQueryChosenChat,
				CallbackGame:                 rawButton.CallbackGame,
				Pay:                          rawButton.Pay,
			})

		case rawButton.RequestUsers != nil || rawButton.RequestChat != nil ||
			rawButton.RequestContact != nil || rawButton.RequestLocation != nil || rawButton.RequestPoll != nil:
			butts = append(butts, &KeyboardButton{
				Text:            rawButton.Text,
				RequestUsers:    rawButton.RequestUsers,
				RequestChat:     rawButton.RequestChat,
				RequestContact:  rawButton.RequestContact,
				RequestLocation: rawButton.RequestLocation,
				RequestPoll:     rawButton.RequestPoll,
			})

		default:
			return ErrUnsupportedButton
		}
	}

	*r = butts
	return nil
}

func (m *InlineKeyboardMarkup) UnmarshalJSON(data []byte) error {
	var raw struct {
		InlineKeyboard []json.RawMessage `json:"inline_keyboard"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	m.InlineKeyboard = make([]Row, len(raw.InlineKeyboard))
	for i, row := range raw.InlineKeyboard {
		r := Row{}
		err := r.UnmarshalJSON(row)
		if err != nil {
			return err
		}
		m.InlineKeyboard[i] = r
	}
	return nil
}

func (m *ReplyKeyboardMarkup) UnmarshalJSON(data []byte) error {
	var raw struct {
		Keyboard              []json.RawMessage `json:"inline_keyboard"`
		IsPersistent          *bool             `json:"is_persistent,omitempty"`
		ResizeKeyboard        *bool             `json:"resize_keyboard,omitempty"`
		OneTimeKeyboard       *bool             `json:"one_time_keyboard,omitempty"`
		InputFieldPlaceholder *string           `json:"input_field_placeholder,omitempty"`
		Selective             *bool             `json:"selective,omitempty"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	m.Keyboard = make([]Row, len(raw.Keyboard))
	m.IsPersistent = raw.IsPersistent
	m.ResizeKeyboard = raw.ResizeKeyboard
	m.OneTimeKeyboard = raw.OneTimeKeyboard
	m.InputFieldPlaceholder = raw.InputFieldPlaceholder
	m.Selective = raw.Selective

	m.Keyboard = make([]Row, len(raw.Keyboard))
	for i, row := range raw.Keyboard {
		r := Row{}
		err := r.UnmarshalJSON(row)
		if err != nil {
			return err
		}
		m.Keyboard[i] = r
	}
	return nil
}

func (i *InlineKeyboardButton) UnmarshalJSON(data []byte) error {
	type Alias InlineKeyboardButton
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(i),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

func (i *KeyboardButton) UnmarshalJSON(data []byte) error {
	type Alias KeyboardButton
	aux := &struct {
		*Alias
	}{
		Alias: (*Alias)(i),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	return nil
}

func (b *baseMarkUp) UnmarshalJSON(data []byte) error {
	var temp struct {
		InlineKeyboard any `json:"inline_keyboard"`
		Keyboard       any `json:"keyboard"`
		RemoveKeyboard any `json:"remove_keyboard"`
		ForceReply     any `json:"force_reply"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	switch {
	case temp.InlineKeyboard != nil:
		b.markupType = markupTypeInline
		b.inline = &InlineKeyboardMarkup{}
		return b.inline.UnmarshalJSON(data)

	case temp.Keyboard != nil:
		b.markupType = markupTypeKeyboard
		b.keyboard = &ReplyKeyboardMarkup{}
		return b.keyboard.UnmarshalJSON(data)

	case temp.RemoveKeyboard != nil:
		b.markupType = markupTypeRemoveKeyboard
		b.remove = &ReplyKeyboardRemove{}
		return json.Unmarshal(data, b.remove)

	case temp.ForceReply != nil:
		b.markupType = markupTypeForceReply
		b.forceReply = &ForceReplyMarkup{}
		return json.Unmarshal(data, b.forceReply)
	}

	return errors.New("telebot: unknown markup type")
}

// MarshalJSON implements json.Marshaler.
func (b *baseMarkUp) MarshalJSON() ([]byte, error) {
	switch b.markupType {
	case markupTypeInline:
		return json.Marshal(b.inline)
	case markupTypeKeyboard:
		return json.Marshal(b.keyboard)
	case markupTypeRemoveKeyboard:
		return json.Marshal(b.remove)
	case markupTypeForceReply:
		return json.Marshal(b.forceReply)
	}
	return nil, nil
}
