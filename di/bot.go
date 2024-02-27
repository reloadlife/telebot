package di

import (
	"github.com/sarulabs/di"
	tele "go.mamad.dev/telebot"
	"go.mamad.dev/telebot/config"
	"go.mamad.dev/telebot/log"
)

type RouteRegisterFunc func(bot tele.Bot) error

func BotService(configFilePath string, registerRouteFunction RouteRegisterFunc) *di.Def {
	return &di.Def{
		Name: "botService",
		Build: func(c di.Container) (interface{}, error) {
			conf := config.NewConfigFromFile(configFilePath)
			settings := conf.GetSettings()
			bot := tele.New(tele.BotSettings{
				Token:          settings.GetToken(),
				URL:            settings.GetURL(),
				AllowedUpdates: settings.GetAllowedUpdates(),
			})

			botConf := conf.GetBot()

			for _, locale := range conf.GetLocales() {
				localedName := botConf.GetName(locale)
				if localedName != "" {
					localedNameFetched, err := bot.GetName(locale)
					if err != nil {
						return nil, err
					}
					if *localedNameFetched != localedName {
						err = bot.SetName(localedName, locale)
						if err != nil {
							return nil, err
						}
						log.Infof("Bot name set to %s in %s locale", localedName, locale)
					}
				}

				localedDescription := botConf.GetLongDescription(locale)
				if localedDescription != "" {
					localedDescriptionFetched, err := bot.GetDescription(locale)
					if err != nil {
						return nil, err
					}
					if *localedDescriptionFetched != localedDescription {
						err = bot.SetDescription(localedDescription, locale)
						if err != nil {
							return nil, err
						}
						log.Infof("Bot description set to %s in %s locale", localedDescription, locale)
					}
				}

				localedShortDescription := botConf.GetShortDescription(locale)
				if localedShortDescription != "" {
					localedShortDescriptionFetched, err := bot.GetShortDescription(locale)
					if err != nil {
						return nil, err
					}
					if *localedShortDescriptionFetched != localedShortDescription {
						err = bot.SetShortDescription(localedShortDescription, locale)
						if err != nil {
							return nil, err
						}
						log.Infof("Bot short-description set to %s in %s locale", localedShortDescription, locale)
					}
				}

				localedCommands := botConf.GetCommands(locale)
				if len(localedCommands) > 0 {
					commands := make([]tele.BotCommand, 0)
					for command, description := range localedCommands {
						commands = append(commands, tele.BotCommand{
							Command:     command,
							Description: description,
						})
						log.Infof("Command %s -> %s in %s locale", command, description, locale)
					}

					err := bot.SetCommands(commands, locale)
					if err != nil {
						return nil, err
					}
					log.Infof("Commands set in %s locale", locale)
				}
			}

			channelRights := botConf.GetDefaultAdminRights(true)
			if channelRights != nil {
				err := bot.SetDefaultAdministratorRights(channelRights, true)
				if err != nil {
					return nil, err
				}
				log.Infof("Admin rights set (channel)")
			}
			rights := botConf.GetDefaultAdminRights(false)
			if rights != nil {
				err := bot.SetDefaultAdministratorRights(rights, false)
				if err != nil {
					return nil, err
				}
				log.Infof("Admin rights set (group)")
			}

			handles := conf.GetHandles()

			bot.Use(func(next tele.HandlerFunc) tele.HandlerFunc {
				return func(c tele.Context) error {
					// sets default locale (can be changed through whatever, even a middleware)
					c.Set("locale", conf.GetDefaultLocale())
					return next(c)
				}
			})

			for _, handle := range handles {
				bot.Handle(handle.GetCommand(), func(c tele.Context) error {
					h := handle.GetHandler()
					locale := c.Get("locale").(string)
					_, err := c.Reply(h.GetText(locale), h.GetKeyboard(locale))
					return err
				})
			}

			err := registerRouteFunction(bot)
			return bot, err
		},
	}
}

func GetBot() tele.Bot {
	bot, err := getServiceSafe[tele.Bot]("botService")
	if err != nil {
		panic(err)
	}
	return bot
}

func Start() {
	bot := GetBot()
	bot.Start()
}
