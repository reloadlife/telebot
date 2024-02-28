package telebot

import (
	"reflect"
	"regexp"
)

type handlers struct {
	handlers map[any]HandlerFunc
}

func (h *handlers) add(endpoint any, handler HandlerFunc) {
	h.handlers[endpoint] = handler
}

func newHandlers() *handlers {
	return &handlers{
		handlers: make(map[any]HandlerFunc),
	}
}

// runHandler runs the handler function with Context.
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

// Handle registers a handler (HandlerFunc) for the given endpoint, with optional middlewares (MiddlewareFunc).
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

	// handle is a middleware wrapped handler.
	handle := func(c Context) error {
		return applyMiddleware(h, m...)(c)
	}

	// Register the handler.
	switch end := endpoint.(type) {
	case string:
		b.handlers.add(end, handle)
	case CallbackEndpoint:
		b.handlers.add(end.CallbackUnique(), handle)
	case *regexp.Regexp:
		b.handlers.add(end, handle)
	case UpdateHandlerOn:
		b.handlers.add(end, handle)
	default:
		panic("telebot: unsupported endpoint type, only string, *regexp.Regexp, callback, and UpdateHandlerOn are supported.")
	}

}

func (b *bot) handle(event any, c Context) bool {
	if h, ok := b.handlers.handlers[event]; ok {
		b.runHandler(h, c)
		return true
	}
	return false
}
