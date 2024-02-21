package telebot

import "strconv"

type Recipient interface {
	Recipient() string
}

func (c *Chat) Recipient() string {
	return strconv.FormatInt(c.ID, 10)
}

func (u *User) Recipient() string {
	return strconv.FormatInt(u.ID, 10)
}
