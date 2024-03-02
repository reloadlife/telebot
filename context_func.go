package telebot

import "strings"

func (c *nativeContext) Text() string {
	if c.Message() != nil {
		text := c.Message().Text
		if text != nil {
			return *text
		}
		return ""
	}
	return ""
}

func (c *nativeContext) ReplyTo(msg Message, text string, options ...any) (Message, error) {
	chat, id := msg.MessageSig()
	for _, o := range options {
		switch o.(type) {
		case *ReplyParameters, ReplyParameters:
			panic("telebot: cannot use ReplyTo with ReplyParameters")
		}
	}
	options = append(options, &ReplyParameters{
		MessageID: id,
		ChatID:    chat,
	})
	return c.b.SendMessage(chat, text, options...)
}

func (c *nativeContext) Reply(text string, options ...any) (Message, error) {
	return c.ReplyTo(c.Message(), text, options...)
}

func (c *nativeContext) Send(s any, options ...any) (Message, error) {
	chat, _ := c.Message().MessageSig()
	switch x := s.(type) {
	case string:
		return c.b.SendMessage(chat, x, options...)
	case Sendable:
		return x.Send(c.b, chat, options...)
	}

	panic("telebot: invalid Sendable type")
}

func (c *nativeContext) Edit(text string, options ...any) (Message, error) {
	return c.b.EditMessageText(c.Message(), text, options...)
}

func (c *nativeContext) EditOrReply(text string, options ...any) (Message, error) {
	if c.Query() == nil {
		return c.Reply(text, options...)
	}
	return c.Edit(text, options...)
}

func (c *nativeContext) Args() []string {
	var args []string
	if c.Message() != nil {
		args = append(args, c.Message().Payload)
		args = append(args, strings.ReplaceAll(c.Text(), c.Message().Command+" "+c.Message().Payload, " "))
	}

	if c.Callback() != nil {
		args = append(args, strings.Split(c.Callback().Data, "|")...)
	}

	return args
}
