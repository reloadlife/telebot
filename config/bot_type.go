package config

import (
	tele "go.mamad.dev/telebot"
	"strings"
)

type BotConfig interface {
	GetName(l ...string) string
	GetShortDescription(l ...string) string
	GetLongDescription(l ...string) string
	GetCommands(l ...string) map[string]string

	GetDefaultAdminRights(forChannels bool) *tele.Rights
}

type botConfig struct {
	conf *config

	Name               string            `yaml:"name" json:"name,omitempty"`
	ShortDescription   string            `yaml:"ShortDescription" json:"short_description,omitempty"`
	LongDescription    string            `yaml:"LongDescription" json:"long_description,omitempty"`
	Commands           map[string]string `yaml:"commands" json:"commands,omitempty"`
	DefaultAdminRights struct {
		Channels []string `yaml:"channels" json:"channels,omitempty"`
		Groups   []string `yaml:"groups" json:"groups,omitempty"`
	} `yaml:"defaultAdminRights" json:"default_admin_rights"`
}

func (b *botConfig) GetName(l ...string) string {
	if b.Name == "" {
		return ""
	}
	locale := b.conf.GetDefaultLocale()
	if len(l) > 0 {
		locale = l[0]
	}

	if locale != "" {
		lKey := strings.TrimPrefix(b.Name, "locale:")
		return b.conf.l(locale, lKey)
	}

	return b.Name
}

func (b *botConfig) GetShortDescription(l ...string) string {
	if b.ShortDescription == "" {
		return ""
	}
	locale := b.conf.GetDefaultLocale()
	if len(l) > 0 {
		locale = l[0]
	}

	if locale != "" {
		lKey := strings.TrimPrefix(b.ShortDescription, "locale:")
		return b.conf.l(locale, lKey)
	}

	return b.ShortDescription
}

func (b *botConfig) GetLongDescription(l ...string) string {
	if b.LongDescription == "" {
		return ""
	}
	locale := b.conf.GetDefaultLocale()
	if len(l) > 0 {
		locale = l[0]
	}

	if locale != "" {
		lKey := strings.TrimPrefix(b.LongDescription, "locale:")
		return b.conf.l(locale, lKey)
	}

	return b.LongDescription
}

func (b *botConfig) GetCommands(l ...string) map[string]string {
	if len(b.Commands) == 0 {
		return nil
	}
	commands := make(map[string]string)

	locale := b.conf.GetDefaultLocale()
	if len(l) > 0 {
		locale = l[0]
	}

	for k, v := range b.Commands {
		if locale != "" {
			lKey := strings.TrimPrefix(v, "locale:")
			commands[k] = b.conf.l(locale, lKey)
		} else {
			commands[k] = v
		}
	}

	return commands
}

func (b *botConfig) GetDefaultAdminRights(forChannels bool) *tele.Rights {
	if forChannels {
		if len(b.DefaultAdminRights.Channels) == 0 {
			return nil
		}
	}
	if len(b.DefaultAdminRights.Groups) == 0 {
		return nil
	}

	t := tele.Rights{}
	rights := b.DefaultAdminRights.Groups
	if forChannels {
		rights = b.DefaultAdminRights.Channels
	}

	ok := true

	for _, r := range rights {
		switch r {
		case "can_manage_chat":
			t.CanManageChat = ok

		case "can_delete_messages":
			t.CanDeleteMessages = ok

		case "can_manage_video_chats":
			t.CanManageVideoChats = ok

		case "can_restrict_members":
			t.CanRestrictMembers = ok

		case "can_promote_members":
			t.CanPromoteMembers = ok

		case "can_change_info":
			t.CanChangeInfo = ok

		case "can_invite_users":
			t.CanInviteUsers = ok

		case "can_post_messages":
			t.CanPostMessages = &ok

		case "can_edit_messages":
			t.CanEditMessages = &ok

		case "can_pin_messages":
			t.CanPinMessages = &ok

		case "can_post_stories":
			t.CanPostStories = &ok

		case "can_edit_stories":
			t.CanEditStories = &ok

		case "can_delete_stories":
			t.CanDeleteStories = &ok

		case "can_manage_topics":
			t.CanManageTopics = &ok

		default:
			panic("telebot-config: unknown default admin right, " + r)
		}
	}

	return &t
}

func (c *config) GetBot() BotConfig {
	c.conf.Bot.conf = c
	return &c.conf.Bot
}
