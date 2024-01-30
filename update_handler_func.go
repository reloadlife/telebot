package telebot

func (b *bot) handle(event UpdateHandlerOn, c Context) bool {
	if handler, ok := b.handlers[event]; ok {
		b.runHandler(handler, c)
		return true
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
