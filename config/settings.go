package config

import (
	tele "go.mamad.dev/telebot"
	"strings"
)

type Settings interface {
	GetToken() string
	GetURL() string
	GetAllowedUpdates() []tele.UpdateType
	GetDefaultParseMode() tele.ParseMode
}

func (c *config) GetSettings() Settings {
	return &c.conf.Settings
}

type settings struct {
	RootDir          string            `yaml:"rootDir" json:"root_dir,omitempty"`
	LocalesDir       string            `yaml:"localesDir" json:"locales_dir,omitempty"`
	Token            string            `yaml:"TOKEN" json:"token,omitempty"`
	URL              string            `yaml:"URL" json:"url,omitempty"`
	AllowedUpdates   []tele.UpdateType `yaml:"allowed_updates" json:"allowed_updates,omitempty"`
	DefaultParseMode string            `yaml:"defaultParseMode" json:"default_parse_mode,omitempty"`
}

func (s *settings) GetToken() string {
	return s.Token
}

func (s *settings) GetURL() string {
	return s.URL
}

func (s *settings) GetAllowedUpdates() []tele.UpdateType {
	return s.AllowedUpdates
}

func (s *settings) GetDefaultParseMode() tele.ParseMode {
	switch strings.ToLower(s.DefaultParseMode) {
	case "markdown":
		return tele.ParseModeMarkdown
	case "markdownv2":
		return tele.ParseModeMarkdownV2
	case "html":
		return tele.ParseModeHTML
	}
	return tele.ParseModeDefault
}
