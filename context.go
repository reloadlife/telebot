package telebot

import (
	"sync"
)

// HandlerFunc represents a handler function, which is used to handle actual endpoints.
type HandlerFunc func(Context) error

// nativeContext is a native implementation of the Context interface; "context" is taken by context package, maybe there is a better name.
type nativeContext struct {
	b     *bot
	u     Update
	lock  sync.RWMutex
	store map[string]any
}

func (c *nativeContext) Sender() *User {
	switch {
	case c.u.CallbackQuery != nil:
		return c.u.CallbackQuery.Sender
	case c.Message() != nil:
		return c.Message().From
	case c.u.Query != nil:
		return &c.u.Query.From
	case c.u.InlineResult != nil:
		return &c.u.InlineResult.From
	case c.u.ShippingQuery != nil:
		return &c.u.ShippingQuery.From
	case c.u.PreCheckoutQuery != nil:
		return &c.u.PreCheckoutQuery.From
	case c.u.PollAnswer != nil:
		return c.u.PollAnswer.User
	case c.u.MyChatMember != nil:
		return &c.u.MyChatMember.From
	case c.u.ChatMember != nil:
		return &c.u.ChatMember.From
	case c.u.ChatJoinRequest != nil:
		return &c.u.ChatJoinRequest.From
	default:
		return nil
	}
}

func (c *nativeContext) Bot() Bot {
	return c.b
}

func (c *nativeContext) Update() Update {
	return c.u
}

func (c *nativeContext) Message() *AccessibleMessage {
	switch {
	case c.u.Message != nil:
		return c.u.Message
	case c.u.CallbackQuery != nil:
		return c.u.CallbackQuery.Message.AccessibleMessage
	case c.u.EditedMessage != nil:
		return c.u.EditedMessage
	case c.u.ChannelPost != nil:
		if c.u.ChannelPost.PinnedMessage != nil {
			return c.u.ChannelPost.PinnedMessage.AccessibleMessage
		}
		return c.u.ChannelPost
	case c.u.EditedChannelPost != nil:
		return c.u.EditedChannelPost
	default:
		return nil
	}
}

func (c *nativeContext) Callback() *Callback {
	return c.u.CallbackQuery
}

func (c *nativeContext) Query() *InlineQuery {
	return c.u.Query
}

func (c *nativeContext) Set(key string, value any) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.store == nil {
		c.store = make(map[string]any)
	}
	c.store[key] = value
}

func (c *nativeContext) Get(key string) any {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.store[key]
}
