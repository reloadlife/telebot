package telebot

import (
	httpc "go.mamad.dev/telebot/http"
	"time"
)

type bot struct {
	self    *User
	token   string
	updates chan Update
	poller  Poller
	onError func(error, Context)

	group    *Group
	handlers map[string]HandlerFunc

	synchronous bool
	stop        chan chan struct{}
	stopClient  chan struct{}

	httpClient httpc.Client
}

type Bot interface {
	Debug(...any)

	GetUpdates(offset, limit int, timeout time.Duration, allowed ...UpdateType) ([]Update, error)

	Start()
	Stop()
}

type BotSettings struct {
	// OfflineMode runs the Client in Offline Mode for test-purposes.
	OfflineMode bool

	Token string
	URL   string

	Synchronous  bool
	UpdatesCount int
}

// New creates a new bot instance.
func New(s BotSettings) Bot {
	if s.Token == "" {
		panic("telebot: token is required")
	}

	if s.UpdatesCount == 0 {
		s.UpdatesCount = 100
	}

	return &bot{
		token: s.Token,
		onError: func(err error, ctx Context) {
			if s.OfflineMode {
				panic(err)
			}
		},
		poller: &LongPoller{},

		updates:  make(chan Update, s.UpdatesCount),
		handlers: make(map[string]HandlerFunc),
		stop:     make(chan chan struct{}),

		synchronous: s.Synchronous,
		httpClient:  httpc.NewHTTPClient(s.URL, time.Minute),
	}
}

// Start brings bot into motion by consuming incoming updates (see bot.updates channel).
func (b *bot) Start() {
	if b.poller == nil {
		panic("telebot: can't start without a poller, either set one or use NewBot instead")
	}

	if b.stopClient != nil {
		return
	}
	b.stopClient = make(chan struct{})

	stop := make(chan struct{})
	stopConfirm := make(chan struct{})

	go func() {
		b.poller.Poll(b, b.updates, stop)
		close(stopConfirm)
	}()

	for {
		select {

		case upd := <-b.updates:
			b.ProcessUpdate(upd)

		case confirm := <-b.stop:
			close(stop)
			<-stopConfirm
			close(confirm)
			b.stopClient = nil
		}
	}
}

// Stop gracefully shuts the poller down.
func (b *bot) Stop() {
	if b.stopClient != nil {
		close(b.stopClient)
	}
	confirm := make(chan struct{})
	b.stop <- confirm
	<-confirm
}

// StartInWebhook starts the bot in webhook mode.
func (b *bot) StartInWebhook() {
	panic("telebot: webhook mode is not implemented yet")
}

// StopInWebhook stops the bot from webhook mode.
func (b *bot) StopInWebhook() {
	panic("telebot: webhook mode is not implemented yet")
}
