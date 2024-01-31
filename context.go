package telebot

import (
	"errors"
	"go.mamad.dev/telebot/.old"
	"strings"
	"sync"
	"time"
)

// HandlerFunc represents a handler function, which is used to handle actual endpoints.
type HandlerFunc func(Context) error

// Context wraps an update and represents the context of the current event.
type Context interface {
	// Bot returns the bot instance.
	Bot() Bot

	// Update returns the original update.
	Update() Update

	// Message returns stored message if such presented.
	Message() *_old.Message

	// Callback returns stored callback if such presented.
	Callback() *Callback

	// Query returns stored query if such presented.
	Query() *_old.Query

	// InlineResult returns stored inline result if such presented.
	InlineResult() *_old.InlineResult

	// ShippingQuery returns stored shipping query if such presented.
	ShippingQuery() *_old.ShippingQuery

	// PreCheckoutQuery returns stored pre checkout query if such presented.
	PreCheckoutQuery() *_old.PreCheckoutQuery

	// Poll returns stored poll if such presented.
	Poll() *_old.Poll

	// PollAnswer returns stored poll answer if such presented.
	PollAnswer() *_old.PollAnswer

	// ChatMember returns chat member changes.
	ChatMember() *_old.ChatMemberUpdate

	// ChatJoinRequest returns cha
	ChatJoinRequest() *_old.ChatJoinRequest

	// Migration returns both migration from and to chat IDs.
	Migration() (int64, int64)

	// Sender returns the current recipient, depending on the context type.
	// Returns nil if user is not presented.
	Sender() *_old.User

	// Chat returns the current chat, depending on the context type.
	// Returns nil if chat is not presented.
	Chat() *_old.Chat

	// Recipient combines both Sender and Chat functions. If there is no user
	// the chat will be returned. The native context cannot be without sender,
	// but it is useful in the case when the context created intentionally
	// by the NewContext constructor and have only Chat field inside.
	Recipient() _old.Recipient

	// Text returns the message text, depending on the context type.
	// In the case when no related data presented, returns an empty string.
	Text() string

	// Entities returns the message entities, whether it's media caption's or the text's.
	// In the case when no entities presented, returns a nil.
	Entities() _old.Entities

	// Data returns the current data, depending on the context type.
	// If the context contains command, returns its arguments string.
	// If the context contains payment, returns its payload.
	// In the case when no related data presented, returns an empty string.
	Data() string

	// Args returns a raw slice of command or callback arguments as strings.
	// The message arguments split by space, while the callback's ones by a "|" symbol.
	Args() []string

	// Send sends a message to the current recipient.
	// See Send from bot.go.
	Send(what interface{}, opts ...interface{}) error

	// SendAlbum sends an album to the current recipient.
	// See SendAlbum from bot.go.
	SendAlbum(a _old.Album, opts ...interface{}) error

	// Reply replies to the current message.
	// See Reply from bot.go.
	Reply(what interface{}, opts ...interface{}) error

	// Forward forwards the given message to the current recipient.
	// See Forward from bot.go.
	Forward(msg _old.Editable, opts ...interface{}) error

	// ForwardTo forwards the current message to the given recipient.
	// See Forward from bot.go
	ForwardTo(to _old.Recipient, opts ...interface{}) error

	// Edit edits the current message.
	// See Edit from bot.go.
	Edit(what interface{}, opts ...interface{}) error

	// EditCaption edits the caption of the current message.
	// See EditCaption from bot.go.
	EditCaption(caption string, opts ...interface{}) error

	// EditOrSend edits the current message if the update is callback,
	// otherwise the content is sent to the chat as a separate message.
	EditOrSend(what interface{}, opts ...interface{}) error

	// EditOrReply edits the current message if the update is callback,
	// otherwise the content is replied as a separate message.
	EditOrReply(what interface{}, opts ...interface{}) error

	// Delete removes the current message.
	// See Delete from bot.go.
	Delete() error

	// DeleteAfter waits for the duration to elapse and then removes the
	// message. It handles an error automatically using b.OnError callback.
	// It returns a Timer that can be used to cancel the call using its Stop method.
	DeleteAfter(d time.Duration) *time.Timer

	// Notify updates the chat action for the current recipient.
	// See Notify from bot.go.
	Notify(action ChatAction) error

	// Ship replies to the current shipping query.
	// See Ship from bot.go.
	Ship(what ...interface{}) error

	// Accept finalizes the current deal.
	// See Accept from bot.go.
	Accept(errorMessage ...string) error

	// Answer sends a response to the current inline query.
	// See Answer from bot.go.
	Answer(resp *_old.QueryResponse) error

	// Respond sends a response for the current callback query.
	// See Respond from bot.go.
	Respond(resp ...*CallbackResponse) error

	// Get retrieves data from the context.
	Get(key string) interface{}

	// Set saves data in the context.
	Set(key string, val interface{})
}

// nativeContext is a native implementation of the Context interface; "context" is taken by context package, maybe there is a better name.
type nativeContext struct {
	b     *bot
	u     Update
	lock  sync.RWMutex
	store map[string]interface{}
}

func (c *nativeContext) Bot() Bot {
	return c.b
}

func (c *nativeContext) Update() Update {
	return c.u
}

func (c *nativeContext) Message() *_old.Message {
	switch {
	case c.u.Message != nil:
		return c.u.Message
	case c.u.Callback != nil:
		return c.u.Callback.Message
	case c.u.EditedMessage != nil:
		return c.u.EditedMessage
	case c.u.ChannelPost != nil:
		if c.u.ChannelPost.PinnedMessage != nil {
			return c.u.ChannelPost.PinnedMessage
		}
		return c.u.ChannelPost
	case c.u.EditedChannelPost != nil:
		return c.u.EditedChannelPost
	default:
		return nil
	}
}

func (c *nativeContext) Callback() *Callback {
	return c.u.Callback
}

func (c *nativeContext) Query() *_old.Query {
	return c.u.Query
}

func (c *nativeContext) InlineResult() *_old.InlineResult {
	return c.u.InlineResult
}

func (c *nativeContext) ShippingQuery() *_old.ShippingQuery {
	return c.u.ShippingQuery
}

func (c *nativeContext) PreCheckoutQuery() *_old.PreCheckoutQuery {
	return c.u.PreCheckoutQuery
}

func (c *nativeContext) ChatMember() *_old.ChatMemberUpdate {
	switch {
	case c.u.ChatMember != nil:
		return c.u.ChatMember
	case c.u.MyChatMember != nil:
		return c.u.MyChatMember
	default:
		return nil
	}
}

func (c *nativeContext) ChatJoinRequest() *_old.ChatJoinRequest {
	return c.u.ChatJoinRequest
}

func (c *nativeContext) Poll() *_old.Poll {
	return c.u.Poll
}

func (c *nativeContext) PollAnswer() *_old.PollAnswer {
	return c.u.PollAnswer
}

func (c *nativeContext) Migration() (int64, int64) {
	return c.u.Message.MigrateFrom, c.u.Message.MigrateTo
}

func (c *nativeContext) Sender() *_old.User {
	switch {
	case c.u.Callback != nil:
		return c.u.Callback.Sender
	case c.Message() != nil:
		return c.Message().Sender
	case c.u.Query != nil:
		return c.u.Query.Sender
	case c.u.InlineResult != nil:
		return c.u.InlineResult.Sender
	case c.u.ShippingQuery != nil:
		return c.u.ShippingQuery.Sender
	case c.u.PreCheckoutQuery != nil:
		return c.u.PreCheckoutQuery.Sender
	case c.u.PollAnswer != nil:
		return c.u.PollAnswer.Sender
	case c.u.MyChatMember != nil:
		return c.u.MyChatMember.Sender
	case c.u.ChatMember != nil:
		return c.u.ChatMember.Sender
	case c.u.ChatJoinRequest != nil:
		return c.u.ChatJoinRequest.Sender
	default:
		return nil
	}
}

func (c *nativeContext) Chat() *_old.Chat {
	switch {
	case c.Message() != nil:
		return c.Message().Chat
	case c.u.MyChatMember != nil:
		return c.u.MyChatMember.Chat
	case c.u.ChatMember != nil:
		return c.u.ChatMember.Chat
	case c.u.ChatJoinRequest != nil:
		return c.u.ChatJoinRequest.Chat
	default:
		return nil
	}
}

func (c *nativeContext) Recipient() _old.Recipient {
	chat := c.Chat()
	if chat != nil {
		return chat
	}
	return c.Sender()
}

func (c *nativeContext) Text() string {
	m := c.Message()
	if m == nil {
		return ""
	}
	if m.Caption != "" {
		return m.Caption
	}
	return m.Text
}

func (c *nativeContext) Entities() _old.Entities {
	m := c.Message()
	if m == nil {
		return nil
	}
	if len(m.CaptionEntities) > 0 {
		return m.CaptionEntities
	}
	return m.Entities
}

func (c *nativeContext) Data() string {
	switch {
	case c.u.Message != nil:
		return c.u.Message.Payload
	case c.u.Callback != nil:
		return c.u.Callback.Data
	case c.u.Query != nil:
		return c.u.Query.Text
	case c.u.InlineResult != nil:
		return c.u.InlineResult.Query
	case c.u.ShippingQuery != nil:
		return c.u.ShippingQuery.Payload
	case c.u.PreCheckoutQuery != nil:
		return c.u.PreCheckoutQuery.Payload
	default:
		return ""
	}
}

func (c *nativeContext) Args() []string {
	switch {
	case c.u.Message != nil:
		payload := strings.Trim(c.u.Message.Payload, " ")
		if payload != "" {
			return strings.Split(payload, " ")
		}
	case c.u.Callback != nil:
		return strings.Split(c.u.Callback.Data, "|")
	case c.u.Query != nil:
		return strings.Split(c.u.Query.Text, " ")
	case c.u.InlineResult != nil:
		return strings.Split(c.u.InlineResult.Query, " ")
	}
	return nil
}

func (c *nativeContext) Send(what interface{}, opts ...interface{}) error {
	_, err := c.b.Send(c.Recipient(), what, opts...)
	return err
}

func (c *nativeContext) SendAlbum(a _old.Album, opts ...interface{}) error {
	_, err := c.b.SendAlbum(c.Recipient(), a, opts...)
	return err
}

func (c *nativeContext) Reply(what interface{}, opts ...interface{}) error {
	msg := c.Message()
	if msg == nil {
		return ErrBadContext
	}
	_, err := c.b.Reply(msg, what, opts...)
	return err
}

func (c *nativeContext) Forward(msg _old.Editable, opts ...interface{}) error {
	_, err := c.b.Forward(c.Recipient(), msg, opts...)
	return err
}

func (c *nativeContext) ForwardTo(to _old.Recipient, opts ...interface{}) error {
	msg := c.Message()
	if msg == nil {
		return ErrBadContext
	}
	_, err := c.b.Forward(to, msg, opts...)
	return err
}

func (c *nativeContext) Edit(what interface{}, opts ...interface{}) error {
	if c.u.InlineResult != nil {
		_, err := c.b.Edit(c.u.InlineResult, what, opts...)
		return err
	}
	if c.u.Callback != nil {
		_, err := c.b.Edit(c.u.Callback, what, opts...)
		return err
	}
	return ErrBadContext
}

func (c *nativeContext) EditCaption(caption string, opts ...interface{}) error {
	if c.u.InlineResult != nil {
		_, err := c.b.EditCaption(c.u.InlineResult, caption, opts...)
		return err
	}
	if c.u.Callback != nil {
		_, err := c.b.EditCaption(c.u.Callback, caption, opts...)
		return err
	}
	return ErrBadContext
}

func (c *nativeContext) EditOrSend(what interface{}, opts ...interface{}) error {
	err := c.Edit(what, opts...)
	if err == ErrBadContext {
		return c.Send(what, opts...)
	}
	return err
}

func (c *nativeContext) EditOrReply(what interface{}, opts ...interface{}) error {
	err := c.Edit(what, opts...)
	if err == ErrBadContext {
		return c.Reply(what, opts...)
	}
	return err
}

func (c *nativeContext) Delete() error {
	msg := c.Message()
	if msg == nil {
		return ErrBadContext
	}
	return c.b.Delete(msg)
}

func (c *nativeContext) DeleteAfter(d time.Duration) *time.Timer {
	return time.AfterFunc(d, func() {
		if err := c.Delete(); err != nil {
			c.b.OnError(err, c)
		}
	})
}

func (c *nativeContext) Notify(action ChatAction) error {
	return c.b.Notify(c.Recipient(), action)
}

func (c *nativeContext) Ship(what ...interface{}) error {
	if c.u.ShippingQuery == nil {
		return errors.New("telebot: context shipping query is nil")
	}
	return c.b.Ship(c.u.ShippingQuery, what...)
}

func (c *nativeContext) Accept(errorMessage ...string) error {
	if c.u.PreCheckoutQuery == nil {
		return errors.New("telebot: context pre checkout query is nil")
	}
	return c.b.Accept(c.u.PreCheckoutQuery, errorMessage...)
}

func (c *nativeContext) Respond(resp ...*CallbackResponse) error {
	if c.u.Callback == nil {
		return errors.New("telebot: context callback is nil")
	}
	return c.b.Respond(c.u.Callback, resp...)
}

func (c *nativeContext) Answer(resp *_old.QueryResponse) error {
	if c.u.Query == nil {
		return errors.New("telebot: context inline query is nil")
	}
	return c.b.Answer(c.u.Query, resp)
}

func (c *nativeContext) Set(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.store == nil {
		c.store = make(map[string]interface{})
	}
	c.store[key] = value
}

func (c *nativeContext) Get(key string) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.store[key]
}
