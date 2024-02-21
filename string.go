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
