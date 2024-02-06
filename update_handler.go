package telebot

import (
	"regexp"
	"strings"
)

func (b *bot) NewContext(u Update) Context {
	return &nativeContext{b: b, u: u}
}

func (b *bot) ProcessUpdate(u Update) bool {
	c := b.NewContext(u)

	handled := false

	switch {
	case u.Message != nil:
		handled = b.handleMessageEvents(u.Message, c)
	case u.EditedMessage != nil:
		handled = b.handle(OnEditedMessage, c)
	case u.ChannelPost != nil:
		handled = b.handleMessageEvents(u.ChannelPost, c)
		handled = handled || b.handle(OnChannelPost, c)
	case u.EditedChannelPost != nil:
		handled = b.handle(OnEditedChannelPost, c)
	case u.MessageReaction != nil:
		handled = b.handle(OnMessageReaction, c)
	case u.MessageReactionCount != nil:
		handled = b.handle(OnMessageReactionCount, c)
	case u.Query != nil:
		handled = b.handle(OnInlineQuery, c)
	case u.InlineResult != nil:
		handled = b.handle(OnChosenInlineResult, c)
	case u.Callback != nil:
		handled = b.handleCallbackEvent(u.Callback, c)
	case u.ShippingQuery != nil:
		handled = b.handle(OnShippingQuery, c)
	case u.PreCheckoutQuery != nil:
		handled = b.handle(OnPreCheckoutQuery, c)
	case u.Poll != nil:
		handled = b.handle(OnPoll, c)
	case u.PollAnswer != nil:
		handled = b.handle(OnPollAnswer, c)
	case u.MyChatMember != nil:
		handled = b.handle(OnMyChatMember, c)
	case u.ChatMember != nil:
		handled = b.handle(OnChatMember, c)
	case u.ChatJoinRequest != nil:
		handled = b.handle(OnChatJoinRequest, c)
	case u.ChatBoost != nil:
		handled = b.handle(OnChatBoost, c)
	case u.ChatBoostRemoved != nil:
		handled = b.handle(OnChatBoostRemoved, c)

	default:
		handled = b.handle(OnAny, c)
	}

	if !handled {
		handled = b.handle(OnAny, c)
	}

	return handled
}

func (b *bot) handleMessageEvents(m *Message, c Context) bool {
	if m.PinnedMessage != nil {
		return b.handle(OnPinnedMessage, c)
	}

	if m.Text != "" {
		match := commandRegex.FindAllStringSubmatch(m.Text, -1)
		if match != nil {
			command, botName := match[0][1], match[0][3]
			if botName != "" && !strings.EqualFold(b.self.Username, botName) {
				return false
			}

			m.Command = command
			m.Payload = match[0][5]

			if b.handle(OnCommand, c, m.Command) {
				return true
			}
		}

		if b.handle(onMatch, c, m.Text) || b.handle(OnText, c) {
			return true
		}
	}

	return b.handleMediaAndOtherEvents(m, c)
}

func (b *bot) handleMediaAndOtherEvents(m *Message, c Context) bool {
	mediaHandlers := map[interface{}]UpdateHandlerOn{
		m.Photo:     OnPhoto,
		m.Voice:     OnVoice,
		m.Audio:     OnAudio,
		m.Animation: OnAnimation,
		m.Document:  OnDocument,
		m.Sticker:   OnSticker,
		m.Video:     OnVideo,
		m.VideoNote: OnVideoNote,
	}

	// Check for media events
	for media, event := range mediaHandlers {
		if media != nil {
			return b.handle(event, c)
		}
	}

	// Check for non-media events
	switch {
	case m.Contact != nil:
		return b.handle(OnContact, c)
	case m.Location != nil:
		return b.handle(OnLocation, c)
	case m.Venue != nil:
		return b.handle(OnVenue, c)
	case m.Game != nil:
		return b.handle(OnGame, c)
	case m.Dice != nil:
		return b.handle(OnDice, c)
	case m.Invoice != nil:
		return b.handle(OnInvoice, c)
	case m.Payment != nil:
		return b.handle(OnPayment, c)
	default:
		return b.handle(OnAny, c)
	}
}

func (b *bot) handleCallbackEvent(callback *Callback, c Context) bool {
	data := callback.Data
	if data != "" && data[0] == '\f' {
		match := callbackRegex.FindAllStringSubmatch(data, -1)
		if match != nil {
			unique, payload := match[0][1], match[0][3]
			if unique == "" {
				return b.handle(OnCallback, c)
			}

			for _, h := range b.handlers {
				if h.e != OnCallback {
					continue
				}
				if h.f == nil {
					continue
				}

				if h.t == unique {
					callback.Unique = unique
					callback.Data = payload
					b.runHandler(h.f, c)
					return true
				}
			}
		}
	}

	return b.handle(OnCallback, c)
}

var (
	commandRegex  = regexp.MustCompile(`^(/\w+)(@(\w+))?(\s|$)(.+)?`)
	callbackRegex = regexp.MustCompile(`^\f([-\w]+)(\|(.+))?$`)
)
