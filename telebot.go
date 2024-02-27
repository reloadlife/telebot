package telebot

import (
	"errors"
	httpc "go.mamad.dev/telebot/http"
	"go.mamad.dev/telebot/log"
	"time"
)

var telegramSecretToken string

type handler struct {
	f HandlerFunc
	e UpdateHandlerOn
	t any
}

type bot struct {
	self    *User
	token   string
	updates chan Update
	poller  Poller
	onError func(error, Context)

	group    *Group
	handlers []handler

	synchronous bool
	stop        chan chan struct{}
	stopClient  chan struct{}

	httpClient httpc.Client

	offlineMode bool
}

type BotSettings struct {
	// OfflineMode runs the Client in Offline Mode for test-purposes.
	OfflineMode bool

	Token string
	URL   string

	Synchronous  bool
	UpdatesCount int

	Poller Poller

	AllowedUpdates []UpdateType
}

// New creates a new bot instance.
func New(s BotSettings) Bot {
	if s.Token == "" {
		panic("telebot: token is required")
	}

	if s.UpdatesCount == 0 {
		s.UpdatesCount = 100
	}

	if s.Poller == nil {
		s.Poller = &LongPoller{
			AllowedUpdates: s.AllowedUpdates,
		}
	}

	b := &bot{
		self:  nil,
		token: s.Token,
		onError: func(err error, _ Context) {
			if s.OfflineMode {
				panic(err)
			}
			log.Errorf("%v", err)
		},
		poller: s.Poller,

		offlineMode: s.OfflineMode,

		updates:  make(chan Update, s.UpdatesCount),
		handlers: []handler{},
		stop:     make(chan chan struct{}),

		synchronous: s.Synchronous,
		httpClient:  httpc.NewHTTPClient(s.URL, time.Minute),

		group: new(Group),
	}

	telegramSecretToken = s.Token
	if !s.OfflineMode {
		self, err := b.GetMe()
		if err != nil {
			panic(errors.Join(err, errors.New("telebot: can't get bot info")))
		}
		b.self = self
	}
	if s.OfflineMode {
		b.self = &User{
			ID:        123456789,
			IsBot:     true,
			FirstName: "TeleBot",
		}
	}

	return b
}

// Start brings bot into motion by consuming incoming updates (see bot.updates channel).
func (b *bot) Start() {
	if b.poller == nil {
		panic("telebot: can't start without a poller, either set one or use NewBot instead")
	}

	if b.stopClient != nil {
		return
	}

	log.Infof("telebot: bot started and logged in As %s", b.self)

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

// OnError Debug sends a debug message to the bot.
func (b *bot) OnError(e error, c Context) {
	b.onError(e, c)
}
