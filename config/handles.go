package config

import (
	"fmt"
	tele "go.mamad.dev/telebot"
	"go.mamad.dev/telebot/log"
	"strings"
)

type Handler interface {
	GetKeyboard(l ...string) tele.ReplyMarkup
	GetText(l ...string) string

	GetParseMode() string
	GetLinkPreview() *tele.LinkPreviewOptions
}

type Handle interface {
	GetCommand(l ...string) []any
	GetHandler() Handler
}

type handle struct {
	conf    *config
	Command any     `yaml:"cmd" json:"cmd"`
	Handler handler `yaml:"handlerResponse" json:"handler"`
}

type handler struct {
	conf *config

	Keyboard string `yaml:"keyboard" json:"keyboard"`
	Text     string `yaml:"text" json:"text"`

	ParseMode   string `yaml:"parseMode" json:"parse_mode"`
	LinkPreview any    `yaml:"linkPreview" json:"link_preview"`
}

func (h *handle) GetCommand(l ...string) []any {
	commands := make([]any, 0)
	switch hc := h.Command.(type) {
	case string:
		if strings.HasPrefix(hc, "locale:") {
			lKey := strings.TrimPrefix(hc, "locale:")
			for _, locale := range l {
				commands = append(commands, h.conf.L(locale, lKey))
			}
		} else {
			commands = append(commands, hc)
		}
	case []any:
		for _, c := range hc {
			switch cc := c.(type) {
			case string:
				if strings.HasPrefix(cc, "locale:") {
					lKey := strings.TrimPrefix(cc, "locale:")
					for _, locale := range l {
						commands = append(commands, h.conf.L(locale, lKey))
					}
				} else {
					commands = append(commands, c)
				}
			default:
				panic(fmt.Errorf("`cmd` should be either string or array of strings, but got %T", c))
			}
		}
	default:
		panic(fmt.Errorf("`cmd` should be either string or array of strings, but got %T", hc))
	}
	return commands
}

func (h *handle) GetHandler() Handler {
	g := h.Handler
	g.conf = h.conf
	return &g
}

func (h *handler) GetParseMode() string {
	if h.ParseMode == "" {
		return "-"
	}
	return h.ParseMode
}

func (h *handler) GetLinkPreview() *tele.LinkPreviewOptions {
	t := true
	switch l := h.LinkPreview.(type) {
	case *tele.LinkPreviewOptions:
		return l
	case bool:
		t = h.LinkPreview.(bool)
		return &tele.LinkPreviewOptions{
			IsDisabled: &t,
		}
	case string:
		switch h.LinkPreview.(string) {
		case "true":
			t = false
			return &tele.LinkPreviewOptions{
				IsDisabled: &t,
			}
		case "false":
			t = true
			return &tele.LinkPreviewOptions{
				IsDisabled: &t,
			}
		default:
			return nil
		}
	default:
		return &tele.LinkPreviewOptions{
			IsDisabled: &t,
		}
	}
}

func (h *handler) GetKeyboard(l ...string) tele.ReplyMarkup {
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Error getting keyboard: %v", r)
		}
	}()
	locale := h.conf.GetDefaultLocale()
	if len(l) > 0 {
		locale = l[0]
	}

	return h.conf.Keyboard(h.Keyboard, locale)
}

func (h *handler) GetText(l ...string) string {
	locale := h.conf.GetDefaultLocale()
	if len(l) > 0 {
		locale = l[0]
	}

	if locale != "" {
		lKey := strings.TrimPrefix(h.Text, "locale:")
		return h.conf.L(locale, lKey)
	}
	return h.Text
}

func (c *config) GetHandles() []Handle {
	handles := make([]Handle, 0)

	for _, h := range c.conf.Handles {
		ha := h
		ha.conf = c
		handles = append(handles, &ha)
	}

	return handles
}
