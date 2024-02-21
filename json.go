package telebot

import (
	"encoding/json"
)

func (u *Update) MarshalJSON() ([]byte, error) {
	return json.Marshal(*u)
}
func (u *MaybeInaccessibleMessage) MarshalJSON() ([]byte, error) {
	if u.IsAccessible() {
		return u.AccessibleMessage.MarshalJSON()
	}
	return u.InaccessibleMessage.MarshalJSON()
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
