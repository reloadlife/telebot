package telebot

import "encoding/json"

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

// /// //// ////// ////// ////// ////// ////// ////// ////// ////// ////// //////

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

func (u *MaybeInaccessibleMessage) UnmarshalJSON(b []byte) error {
	if u.IsAccessible() {
		return u.AccessibleMessage.UnmarshalJSON(b)
	}
	return u.InaccessibleMessage.UnmarshalJSON(b)
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
