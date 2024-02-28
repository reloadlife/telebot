package telebot

import (
	"regexp"
	"strings"
)

func (b *bot) NewContext(u Update) Context {
	return &nativeContext{b: b, u: u}
}

func (b *bot) ProcessUpdate(u Update) {
	c := b.NewContext(u)

	if u.Message != nil {
		m := u.Message

		if m.ForwardOrigin != nil {
			if b.handle(OnForwardedMessage, c) {
				return
			}
		}

		if m.PinnedMessage != nil {
			if b.handle(OnPinnedMessage, c) {
				return
			}
		}

		if m.Text != nil {
			text := *m.Text
			match := commandRegex.FindAllStringSubmatch(text, -1)
			if match != nil {
				command, botName := match[0][1], match[0][3]
				if botName != "" && !strings.EqualFold(*b.self.Username, botName) {
					return
				}
				m.Command = command
				m.Payload = match[0][5]
				if b.handle(m.Command, c) {
					return
				}
			}
			if b.handle(text, c) {
				return
			}
			if b.handle(OnText, c) {
				return
			}
		}

		// well if we're here, the message clearly was not a Text.
		// so now, we can handle other stuff, first and foremost, we need to handle ServiceMessage (OnServiceMessage)

		if m.IsService() {
			if b.handle(OnServiceMessage, c) {
				// and if they're being handled by a handler function, well
				// we need to break the cycle.
				return
			}
		}

		// now we can check for Medias (OnMedia) and other media types.

		mediaEvents := map[UpdateHandlerOn]interface{}{
			OnAnimation: m.Animation,
			OnAudio:     m.Audio,
			OnDocument:  m.Document,
			OnStory:     m.Story,
			OnVideo:     m.Video,
			OnVideoNote: m.VideoNote,
			OnVoice:     m.Voice,
			OnPhoto:     m.Photo,
		}

		for event, media := range mediaEvents {
			if media != nil {
				if b.handle(event, c) {
					return
				}

				// if it is a media and not being handled by it's own handler,
				// we might try feeding it to OnMedia handler.

				if b.handle(OnMedia, c) {
					return
				}
			}
		}

		// now we're sure that it's not a media.

		if m.Sticker != nil {
			if b.handle(OnSticker, c) {
				return
			}
		}

		// not even a sticker?

		if m.Location != nil {
			if b.handle(OnLocation, c) {
				return
			}
		}

		if m.Venue != nil {
			if b.handle(OnVenue, c) {
				return
			}
		}

		// well, it was not a location message either

		if m.Contact != nil {
			if b.handle(OnContact, c) {
				return
			}
		}

		if m.Game != nil {
			if b.handle(OnGame, c) {
				return
			}
		}

		if m.Poll != nil {
			if b.handle(OnPoll, c) {
				return
			}
		}

		if m.Dice != nil {
			if b.handle(OnDice, c) {
				return
			}
		}

		if m.Invoice != nil {
			if b.handle(OnInvoice, c) {
				return
			}
		}

		if m.SuccessfulPayment != nil {
			if b.handle(OnPayment, c) {
				return
			}
		}

		// well it was not a message or it was not handled even by OnMessage handler.

		if b.handle(OnMessage, c) {
			return
		}
		// well it was not a message or it was not handled even by OnMessage handler.
	}

	if u.EditedMessage != nil {
		if b.handle(OnEditedMessage, c) {
			return
		}
	}

	// is it a channel post?
	if u.ChannelPost != nil {
		if b.handle(OnChannelPost, c) {
			return
		}
	}

	// is it an edited channel post?
	if u.EditedChannelPost != nil {
		if b.handle(OnEditedChannelPost, c) {
			return
		}
	}

	if u.CallbackQuery != nil {
		if b.handleCallbackEvent(u.CallbackQuery, c) {
			return
		}
	}

	if u.Query != nil {
		if b.handle(OnInlineQuery, c) {
			return
		}
	}

	if u.InlineResult != nil {
		if b.handle(OnChosenInlineResult, c) {
			return
		}
	}

	if u.ShippingQuery != nil {
		if b.handle(OnShippingQuery, c) {
			return
		}
	}

	if u.PreCheckoutQuery != nil {
		if b.handle(OnPreCheckoutQuery, c) {
			return
		}
	}

	if u.Poll != nil {
		if b.handle(OnPoll, c) {
			return
		}
	}

	if u.PollAnswer != nil {
		if b.handle(OnPollAnswer, c) {
			return
		}
	}

	if u.MyChatMember != nil {
		if b.handle(OnMyChatMember, c) {
			return
		}
	}

	if u.ChatMember != nil {
		if b.handle(OnChatMember, c) {
			return
		}
	}

	if u.ChatJoinRequest != nil {
		if b.handle(OnChatJoinRequest, c) {
			return
		}
	}

	if u.ChatBoost != nil {
		if b.handle(OnChatBoost, c) {
			return
		}
	}

	if u.RemovedChatBoost != nil {
		if b.handle(OnChatBoostRemoved, c) {
			return
		}
	}

	if u.Reaction != nil {
		if b.handle(OnMessageReaction, c) {
			return
		}
	}

	if u.ReactionCount != nil {
		if b.handle(OnMessageReactionCount, c) {
			return
		}
	}

	b.handle(OnAny, c)
}

func (b *bot) handleCallbackEvent(callback *Callback, c Context) bool {
	data := callback.Data
	if data != "" && data[0] == '\f' {
		match := callbackRegex.FindAllStringSubmatch(data, -1)
		if match != nil {
			unique, payload := match[0][1], match[0][3]
			callback.Unique = unique
			callback.Data = payload

			if b.handle(unique, c) {
				return true
			}
		}
	}

	return b.handle(OnCallback, c)
}

var (
	commandRegex  = regexp.MustCompile(`^(/\w+)(@(\w+))?(\s|$)(.+)?`)
	callbackRegex = regexp.MustCompile(`^\f([-\w]+)(\|(.+))?$`)
)
