package _old

import (
	"errors"
	"go.mamad.dev/telebot"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	// required to test send and edit methods
	token     = os.Getenv("TELEBOT_SECRET")
	chatID, _ = strconv.ParseInt(os.Getenv("CHAT_ID"), 10, 64)
	userID, _ = strconv.ParseInt(os.Getenv("USER_ID"), 10, 64)

	b, _ = newTestBot()      // cached bot instance to avoid getMe method flooding
	to   = &Chat{ID: chatID} // to chat recipient for send and edit methods
	user = &User{ID: userID} // to user recipient for some special cases
)

func defaultSettings() Settings {
	return Settings{Token: token}
}

func newTestBot() (*OldBot, error) {
	return NewBot(defaultSettings())
}

func TestNewBot(t *testing.T) {
	var pref Settings
	_, err := NewBot(pref)
	assert.Error(t, err)

	pref.Token = "BAD TOKEN"
	_, err = NewBot(pref)
	assert.Error(t, err)

	pref.URL = "BAD URL"
	_, err = NewBot(pref)
	assert.Error(t, err)

	b, err := NewBot(Settings{Offline: true})
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, b.Me)
	assert.Equal(t, DefaultApiURL, b.URL)
	assert.Equal(t, 100, cap(b.Updates))
	assert.NotZero(t, b.client.Timeout)

	pref = defaultSettings()
	client := &http.Client{Timeout: time.Minute}
	pref.URL = "http://api.telegram.org" // not https
	pref.Client = client
	pref.Poller = &telebot.LongPoller{Timeout: time.Second}
	pref.Updates = 50
	pref.ParseMode = ModeHTML
	pref.Offline = true

	b, err = NewBot(pref)
	require.NoError(t, err)
	assert.Equal(t, client, b.client)
	assert.Equal(t, pref.URL, b.URL)
	assert.Equal(t, pref.Poller, b.Poller)
	assert.Equal(t, 50, cap(b.Updates))
	assert.Equal(t, ModeHTML, b.parseMode)
}

func TestBotHandle(t *testing.T) {
	if b == nil {
		t.Skip("Cached bot instance is bad (probably wrong or empty TELEBOT_SECRET)")
	}

	b.Handle("/start", func(c telebot.Context) error { return nil })
	assert.Contains(t, b.handlers, "/start")

	reply := ReplyButton{Text: "reply"}
	b.Handle(&reply, func(c telebot.Context) error { return nil })

	inline := InlineButton{Unique: "inline"}
	b.Handle(&inline, func(c telebot.Context) error { return nil })

	btnReply := (&ReplyMarkup{}).Text("btnReply")
	b.Handle(&btnReply, func(c telebot.Context) error { return nil })

	btnInline := (&ReplyMarkup{}).Data("", "btnInline")
	b.Handle(&btnInline, func(c telebot.Context) error { return nil })

	assert.Contains(t, b.handlers, btnReply.CallbackUnique())
	assert.Contains(t, b.handlers, btnInline.CallbackUnique())
	assert.Contains(t, b.handlers, reply.CallbackUnique())
	assert.Contains(t, b.handlers, inline.CallbackUnique())
}

func TestBotStart(t *testing.T) {
	if token == "" {
		t.Skip("TELEBOT_SECRET is required")
	}

	pref := defaultSettings()
	pref.Poller = &telebot.LongPoller{}

	b, err := NewBot(pref)
	if err != nil {
		t.Fatal(err)
	}

	// remove webhook to be sure that bot can poll
	require.NoError(t, b.RemoveWebhook())

	go b.Start()
	b.Stop()

	tp := telebot.newTestPoller()
	go func() {
		tp.updates <- telebot.Update{Message: &Message{Text: "/start"}}
	}()

	b, err = NewBot(pref)
	require.NoError(t, err)
	b.Poller = tp

	var ok bool
	b.Handle("/start", func(c telebot.Context) error {
		assert.Equal(t, c.Text(), "/start")
		tp.done <- struct{}{}
		ok = true
		return nil
	})

	go b.Start()
	<-tp.done
	b.Stop()

	assert.True(t, ok)
}

func TestBotProcessUpdate(t *testing.T) {
	b, err := NewBot(Settings{Synchronous: true, Offline: true})
	if err != nil {
		t.Fatal(err)
	}

	b.Handle(telebot.OnMedia, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().Photo)
		return nil
	})
	b.ProcessUpdate(telebot.Update{Message: &Message{Photo: &Photo{}}})

	b.Handle("/start", func(c telebot.Context) error {
		assert.Equal(t, "/start", c.Text())
		return nil
	})
	b.Handle("hello", func(c telebot.Context) error {
		assert.Equal(t, "hello", c.Text())
		return nil
	})
	b.Handle(telebot.OnText, func(c telebot.Context) error {
		assert.Equal(t, "text", c.Text())
		return nil
	})
	b.Handle(OnPinned, func(c telebot.Context) error {
		assert.NotNil(t, c.Message())
		return nil
	})
	b.Handle(telebot.OnPhoto, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().Photo)
		return nil
	})
	b.Handle(telebot.OnVoice, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().Voice)
		return nil
	})
	b.Handle(telebot.OnAudio, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().Audio)
		return nil
	})
	b.Handle(telebot.OnAnimation, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().Animation)
		return nil
	})
	b.Handle(telebot.OnDocument, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().Document)
		return nil
	})
	b.Handle(telebot.OnSticker, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().Sticker)
		return nil
	})
	b.Handle(telebot.OnVideo, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().Video)
		return nil
	})
	b.Handle(telebot.OnVideoNote, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().VideoNote)
		return nil
	})
	b.Handle(telebot.OnContact, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().Contact)
		return nil
	})
	b.Handle(telebot.OnLocation, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().Location)
		return nil
	})
	b.Handle(telebot.OnVenue, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().Venue)
		return nil
	})
	b.Handle(telebot.OnDice, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().Dice)
		return nil
	})
	b.Handle(telebot.OnInvoice, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().Invoice)
		return nil
	})
	b.Handle(telebot.OnPayment, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().Payment)
		return nil
	})
	b.Handle(OnAddedToGroup, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().GroupCreated)
		return nil
	})
	b.Handle(OnUserJoined, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().UserJoined)
		return nil
	})
	b.Handle(OnUserLeft, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().UserLeft)
		return nil
	})
	b.Handle(OnNewGroupTitle, func(c telebot.Context) error {
		assert.Equal(t, "title", c.Message().NewGroupTitle)
		return nil
	})
	b.Handle(OnNewGroupPhoto, func(c telebot.Context) error {
		assert.NotNil(t, c.Message().NewGroupPhoto)
		return nil
	})
	b.Handle(OnGroupPhotoDeleted, func(c telebot.Context) error {
		assert.True(t, c.Message().GroupPhotoDeleted)
		return nil
	})
	b.Handle(OnMigration, func(c telebot.Context) error {
		from, to := c.Migration()
		assert.Equal(t, int64(1), from)
		assert.Equal(t, int64(2), to)
		return nil
	})
	b.Handle(OnEdited, func(c telebot.Context) error {
		assert.Equal(t, "edited", c.Message().Text)
		return nil
	})
	b.Handle(telebot.OnChannelPost, func(c telebot.Context) error {
		assert.Equal(t, "post", c.Message().Text)
		return nil
	})
	b.Handle(telebot.OnEditedChannelPost, func(c telebot.Context) error {
		assert.Equal(t, "edited post", c.Message().Text)
		return nil
	})
	b.Handle(telebot.OnCallback, func(c telebot.Context) error {
		if data := c.Callback().Data; data[0] != '\f' {
			assert.Equal(t, "callback", data)
		}
		return nil
	})
	b.Handle("\funique", func(c telebot.Context) error {
		assert.Equal(t, "callback", c.Callback().Data)
		return nil
	})
	b.Handle(OnQuery, func(c telebot.Context) error {
		assert.Equal(t, "query", c.Data())
		return nil
	})
	b.Handle(OnInlineResult, func(c telebot.Context) error {
		assert.Equal(t, "result", c.InlineResult().ResultID)
		return nil
	})
	b.Handle(OnShipping, func(c telebot.Context) error {
		assert.Equal(t, "shipping", c.ShippingQuery().ID)
		return nil
	})
	b.Handle(OnCheckout, func(c telebot.Context) error {
		assert.Equal(t, "checkout", c.PreCheckoutQuery().ID)
		return nil
	})
	b.Handle(telebot.OnPoll, func(c telebot.Context) error {
		assert.Equal(t, "poll", c.Poll().ID)
		return nil
	})
	b.Handle(telebot.OnPollAnswer, func(c telebot.Context) error {
		assert.Equal(t, "poll", c.PollAnswer().PollID)
		return nil
	})

	b.Handle(OnWebApp, func(c telebot.Context) error {
		assert.Equal(t, "webapp", c.Message().WebAppData.Data)
		return nil
	})

	b.ProcessUpdate(telebot.Update{Message: &Message{Text: "/start"}})
	b.ProcessUpdate(telebot.Update{Message: &Message{Text: "/start@other_bot"}})
	b.ProcessUpdate(telebot.Update{Message: &Message{Text: "hello"}})
	b.ProcessUpdate(telebot.Update{Message: &Message{Text: "text"}})
	b.ProcessUpdate(telebot.Update{Message: &Message{PinnedMessage: &Message{}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{Photo: &Photo{}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{Voice: &Voice{}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{Audio: &Audio{}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{Animation: &Animation{}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{Document: &Document{}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{Sticker: &Sticker{}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{Video: &Video{}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{VideoNote: &VideoNote{}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{Contact: &Contact{}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{Location: &Location{}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{Venue: &Venue{}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{Invoice: &Invoice{}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{Payment: &Payment{}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{Dice: &Dice{}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{GroupCreated: true}})
	b.ProcessUpdate(telebot.Update{Message: &Message{UserJoined: &User{ID: 1}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{UsersJoined: []User{{ID: 1}}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{UserLeft: &User{}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{NewGroupTitle: "title"}})
	b.ProcessUpdate(telebot.Update{Message: &Message{NewGroupPhoto: &Photo{}}})
	b.ProcessUpdate(telebot.Update{Message: &Message{GroupPhotoDeleted: true}})
	b.ProcessUpdate(telebot.Update{Message: &Message{Chat: &Chat{ID: 1}, MigrateTo: 2}})
	b.ProcessUpdate(telebot.Update{EditedMessage: &Message{Text: "edited"}})
	b.ProcessUpdate(telebot.Update{ChannelPost: &Message{Text: "post"}})
	b.ProcessUpdate(telebot.Update{ChannelPost: &Message{PinnedMessage: &Message{}}})
	b.ProcessUpdate(telebot.Update{EditedChannelPost: &Message{Text: "edited post"}})
	b.ProcessUpdate(telebot.Update{Callback: &telebot.Callback{MessageID: "inline", Data: "callback"}})
	b.ProcessUpdate(telebot.Update{Callback: &telebot.Callback{Data: "callback"}})
	b.ProcessUpdate(telebot.Update{Callback: &telebot.Callback{Data: "\funique|callback"}})
	b.ProcessUpdate(telebot.Update{Query: &Query{Text: "query"}})
	b.ProcessUpdate(telebot.Update{InlineResult: &InlineResult{ResultID: "result"}})
	b.ProcessUpdate(telebot.Update{ShippingQuery: &ShippingQuery{ID: "shipping"}})
	b.ProcessUpdate(telebot.Update{PreCheckoutQuery: &PreCheckoutQuery{ID: "checkout"}})
	b.ProcessUpdate(telebot.Update{Poll: &Poll{ID: "poll"}})
	b.ProcessUpdate(telebot.Update{PollAnswer: &PollAnswer{PollID: "poll"}})
	b.ProcessUpdate(telebot.Update{Message: &Message{WebAppData: &WebAppData{Data: "webapp"}}})
}

func TestBotOnError(t *testing.T) {
	b, err := NewBot(Settings{Synchronous: true, Offline: true})
	if err != nil {
		t.Fatal(err)
	}

	var ok bool
	b.onError = func(err error, c telebot.Context) {
		assert.Equal(t, b, c.(*telebot.nativeContext).b)
		assert.NotNil(t, err)
		ok = true
	}

	b.runHandler(func(c telebot.Context) error {
		return errors.New("not nil")
	}, &telebot.nativeContext{b: b})

	assert.True(t, ok)
}

func TestBot(t *testing.T) {
	if b == nil {
		t.Skip("Cached bot instance is bad (probably wrong or empty TELEBOT_SECRET)")
	}
	if chatID == 0 {
		t.Skip("CHAT_ID is required for OldBot methods test")
	}

	_, err := b.Send(to, nil)
	assert.Equal(t, ErrUnsupportedWhat, err)
	_, err = b.Edit(&Message{Chat: &Chat{}}, nil)
	assert.Equal(t, ErrUnsupportedWhat, err)

	_, err = b.Send(nil, "")
	assert.Equal(t, ErrBadRecipient, err)
	_, err = b.Forward(nil, nil)
	assert.Equal(t, ErrBadRecipient, err)

	photo := &Photo{
		File:    telebot.FromURL("https://telegra.ph/file/65c5237b040ebf80ec278.jpg"),
		Caption: t.Name(),
	}
	var msg *Message

	t.Run("Send(what=Sendable)", func(t *testing.T) {
		msg, err = b.Send(to, photo)
		require.NoError(t, err)
		assert.NotNil(t, msg.Photo)
		assert.Equal(t, photo.Caption, msg.Caption)
	})

	t.Run("SendAlbum()", func(t *testing.T) {
		_, err = b.SendAlbum(nil, nil)
		assert.Equal(t, ErrBadRecipient, err)

		_, err = b.SendAlbum(to, nil)
		assert.Error(t, err)

		photo2 := *photo
		photo2.Caption = ""

		msgs, err := b.SendAlbum(to, Album{photo, &photo2}, ModeHTML)
		require.NoError(t, err)
		assert.Len(t, msgs, 2)
		assert.NotEmpty(t, msgs[0].AlbumID)
	})

	t.Run("EditCaption()+ParseMode", func(t *testing.T) {
		b.parseMode = ModeHTML

		edited, err := b.EditCaption(msg, "<b>new caption with html</b>")
		require.NoError(t, err)
		assert.Equal(t, "new caption with html", edited.Caption)
		assert.Equal(t, EntityBold, edited.CaptionEntities[0].Type)

		sleep()

		edited, err = b.EditCaption(msg, "*new caption with markdown*", ModeMarkdown)
		require.NoError(t, err)
		assert.Equal(t, "new caption with markdown", edited.Caption)
		assert.Equal(t, EntityBold, edited.CaptionEntities[0].Type)

		sleep()

		edited, err = b.EditCaption(msg, "_new caption with markdown \\(V2\\)_", ModeMarkdownV2)
		require.NoError(t, err)
		assert.Equal(t, "new caption with markdown (V2)", edited.Caption)
		assert.Equal(t, EntityItalic, edited.CaptionEntities[0].Type)

		b.parseMode = ModeDefault
	})

	t.Run("Edit(what=Media)", func(t *testing.T) {
		edited, err := b.Edit(msg, photo)
		require.NoError(t, err)
		assert.Equal(t, edited.Photo.UniqueID, photo.UniqueID)

		resp, err := http.Get("https://telegra.ph/file/274e5eb26f348b10bd8ee.mp4")
		require.NoError(t, err)
		defer resp.Body.Close()

		file, err := ioutil.TempFile("", "")
		require.NoError(t, err)

		_, err = io.Copy(file, resp.Body)
		require.NoError(t, err)

		animation := &Animation{
			File:     telebot.FromDisk(file.Name()),
			Caption:  t.Name(),
			FileName: "animation.gif",
		}

		msg, err := b.Send(msg.Chat, animation)
		require.NoError(t, err)

		if msg.Animation != nil {
			assert.Equal(t, msg.Animation.FileID, animation.FileID)
		} else {
			assert.Equal(t, msg.Document.FileID, animation.FileID)
		}

		_, err = b.Edit(edited, animation)
		require.NoError(t, err)
	})

	t.Run("Edit(what=Animation)", func(t *testing.T) {})

	t.Run("Send(what=string)", func(t *testing.T) {
		msg, err = b.Send(to, t.Name())
		require.NoError(t, err)
		assert.Equal(t, t.Name(), msg.Text)

		rpl, err := b.Reply(msg, t.Name())
		require.NoError(t, err)
		assert.Equal(t, rpl.Text, msg.Text)
		assert.NotNil(t, rpl.ReplyTo)
		assert.Equal(t, rpl.ReplyTo, msg)
		assert.True(t, rpl.IsReply())

		fwd, err := b.Forward(to, msg)
		require.NoError(t, err)
		assert.NotNil(t, msg, fwd)
		assert.True(t, fwd.IsForwarded())

		fwd.ID += 1 // nonexistent message
		_, err = b.Forward(to, fwd)
		assert.Equal(t, telebot.ErrNotFoundToForward, err)
	})

	t.Run("Edit(what=string)", func(t *testing.T) {
		msg, err = b.Edit(msg, t.Name())
		require.NoError(t, err)
		assert.Equal(t, t.Name(), msg.Text)

		_, err = b.Edit(msg, msg.Text)
		assert.Error(t, err) // message is not modified
	})

	t.Run("Edit(what=ReplyMarkup)", func(t *testing.T) {
		good := &ReplyMarkup{
			InlineKeyboard: [][]InlineButton{
				{{
					Data: "btn",
					Text: "Hi Telebot!",
				}},
			},
		}
		bad := &ReplyMarkup{
			InlineKeyboard: [][]InlineButton{
				{{
					Data: strings.Repeat("*", 65),
					Text: "Bad Button",
				}},
			},
		}

		edited, err := b.Edit(msg, good)
		require.NoError(t, err)
		assert.Equal(t, edited.ReplyMarkup.InlineKeyboard, good.InlineKeyboard)

		edited, err = b.EditReplyMarkup(edited, nil)
		require.NoError(t, err)
		assert.Nil(t, edited.ReplyMarkup)

		_, err = b.Edit(edited, bad)
		assert.Equal(t, telebot.ErrBadButtonData, err)
	})

	t.Run("Edit(what=Location)", func(t *testing.T) {
		loc := &Location{Lat: 42, Lng: 69, LivePeriod: 60}
		edited, err := b.Send(to, loc)
		require.NoError(t, err)
		assert.NotNil(t, edited.Location)

		loc = &Location{Lat: loc.Lng, Lng: loc.Lat}
		edited, err = b.Edit(edited, *loc)
		require.NoError(t, err)
		assert.NotNil(t, edited.Location)
	})

	// Should be after the Edit tests.
	t.Run("Delete()", func(t *testing.T) {
		require.NoError(t, b.Delete(msg))
	})

	t.Run("Notify()", func(t *testing.T) {
		assert.Equal(t, ErrBadRecipient, b.Notify(nil, Typing))
		require.NoError(t, b.Notify(to, Typing))
	})

	t.Run("Answer()", func(t *testing.T) {
		assert.Error(t, b.Answer(&Query{}, &QueryResponse{
			Results: Results{&ArticleResult{}},
		}))
	})

	t.Run("Respond()", func(t *testing.T) {
		assert.Error(t, b.Respond(&telebot.Callback{}, &telebot.CallbackResponse{}))
	})

	t.Run("Payments", func(t *testing.T) {
		assert.NotPanics(t, func() {
			b.Accept(&PreCheckoutQuery{})
			b.Accept(&PreCheckoutQuery{}, "error")
		})
		assert.NotPanics(t, func() {
			b.Ship(&ShippingQuery{})
			b.Ship(&ShippingQuery{}, "error")
			b.Ship(&ShippingQuery{}, ShippingOption{}, ShippingOption{})
			assert.Equal(t, ErrUnsupportedWhat, b.Ship(&ShippingQuery{}, 0))
		})
	})

	t.Run("Commands", func(t *testing.T) {
		var (
			set1 = []Command{{
				Text:        "test1",
				Description: "test command 1",
			}}
			set2 = []Command{{
				Text:        "test2",
				Description: "test command 2",
			}}
			scope = CommandScope{
				Type:   CommandScopeChat,
				ChatID: chatID,
			}
		)

		err := b.SetCommands(set1)
		require.NoError(t, err)

		cmds, err := b.Commands()
		require.NoError(t, err)
		assert.Equal(t, set1, cmds)

		err = b.SetCommands(set2, "en", scope)
		require.NoError(t, err)

		cmds, err = b.Commands()
		require.NoError(t, err)
		assert.Equal(t, set1, cmds)

		cmds, err = b.Commands("en", scope)
		require.NoError(t, err)
		assert.Equal(t, set2, cmds)

		require.NoError(t, b.DeleteCommands("en", scope))
		require.NoError(t, b.DeleteCommands())
	})

	t.Run("InviteLink", func(t *testing.T) {
		inviteLink, err := b.CreateInviteLink(&Chat{ID: chatID}, nil)
		require.NoError(t, err)
		assert.True(t, len(inviteLink.InviteLink) > 0)

		sleep()

		response, err := b.EditInviteLink(&Chat{ID: chatID}, &ChatInviteLink{InviteLink: inviteLink.InviteLink})
		require.NoError(t, err)
		assert.True(t, len(response.InviteLink) > 0)

		sleep()

		response, err = b.RevokeInviteLink(&Chat{ID: chatID}, inviteLink.InviteLink)
		require.Nil(t, err)
		assert.True(t, len(response.InviteLink) > 0)
	})
}

func sleep() {
	time.Sleep(time.Second)
}
