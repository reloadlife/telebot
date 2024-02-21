// Func: String() for all types
// the reason for this, is that so if anything goes wrong with the String, outputs,
//we can just come here and know where to look.

package telebot

import (
	"encoding/json"
	"fmt"
)

func (u *Update) String() string {
	indented, _ := json.MarshalIndent(u, "", "  ")
	return fmt.Sprintf("%s{%d}\n%s\n", u.Type(), u.ID, indented)
}

func (u *User) String() string {
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

func (c *Poll) String() string {
	indented, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("%s{ID: %s}\n%s\n", c.ReflectType(), c.ID, indented)
}

func (c *PollAnswer) String() string {
	indented, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("%s{ID: %s}\n%s\n", c.ReflectType(), c.ID, indented)
}

func (c *ChatMemberUpdated) String() string {
	indented, _ := json.MarshalIndent(c, "", "  ")
	return fmt.Sprintf("%s{}\n%s\n", c.ReflectType(), indented)
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
