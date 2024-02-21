package telebot

import (
	"encoding/json"
	"errors"
)

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
func (c *ShippingQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)
}
func (c *PreCheckoutQuery) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)
}
func (c *Poll) MarshalJSON() ([]byte, error) {
	return json.Marshal(*c)
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
func (c *ShippingQuery) UnmarshalJSON(b []byte) error {
	type Alias ShippingQuery
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
func (c *Poll) UnmarshalJSON(b []byte) error {
	type Alias Poll
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
