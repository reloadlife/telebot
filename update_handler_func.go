package telebot

import (
	"reflect"
	"regexp"
)

func (b *bot) Group() *Group {
	return &Group{b: b}
}

// Use adds middleware to the global bot chain.
func (b *bot) Use(middleware ...MiddlewareFunc) {
	b.group.Use(middleware...)
}

func (b *bot) Handle(endpoint any, h HandlerFunc, m ...MiddlewareFunc) {
	if reflect.TypeOf(endpoint).Kind() == reflect.Slice {
		for _, e := range endpoint.([]any) {
			b.Handle(e, h, m...)
		}
		return
	}
	if len(b.group.middleware) > 0 {
		m = append(b.group.middleware, m...)
	}

	handle := func(c Context) error {
		return applyMiddleware(h, m...)(c)
	}

	switch end := endpoint.(type) {
	case string:
		b.handlers = append(b.handlers, handler{
			e: onMatch,
			f: handle,
			t: end,
		})
	case CallbackEndpoint:
		b.handlers = append(b.handlers, handler{
			t: end.CallbackUnique(),
			e: OnCallback,
			f: handle,
		})
	case *regexp.Regexp:
		b.handlers = append(b.handlers, handler{
			e: OnCommand,
			f: handle,
			t: end,
		})
	case UpdateHandlerOn:
		b.handlers = append(b.handlers, handler{
			e: end,
			f: handle,
			t: nil,
		})

	default:
		panic("telebot: unsupported endpoint")
	}
}

func (b *bot) handle(event UpdateHandlerOn, c Context, i ...any) bool {
	for _, h := range b.handlers {
		if h.e != event {
			continue
		}
		if h.f == nil {
			continue
		}

		if len(i) > 0 {
			switch i[0].(type) {
			case string:
				if h.t == i[0] {
					b.runHandler(h.f, c)
					return true
				}
			}
			return false
		}

		b.runHandler(h.f, c)
	}

	return false
}

func (b *bot) runHandler(h HandlerFunc, c Context) {
	f := func() {
		if err := h(c); err != nil {
			b.onError(err, c)
		}
	}
	if b.synchronous {
		f()
	} else {
		go f()
	}
}
