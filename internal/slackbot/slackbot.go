package slackbot

import (
	"errors"
	"fmt"
	"github.com/nlopes/slack"
	"github.com/riksa/gonkka/internal/slackbot/trigger"
	log "github.com/sirupsen/logrus"
)

type Bot struct {
	api      *slack.Client
	rtm      *slack.RTM
	triggers []*trigger.Trigger
}

func New(token string) *Bot {
	api := slack.New(token)
	return &Bot{
		api:      api,
		rtm:      api.NewRTM(),
		triggers: make([]*trigger.Trigger, 0),
	}
}

func (bot *Bot) Run() error {
	go bot.rtm.ManageConnection()

	for {
		select {
		case msg := <-bot.rtm.IncomingEvents:

			log.Debug("event received")
			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				info := bot.rtm.GetInfo()

				if ev.User != info.User.ID {
					text := ev.Text
					for _, t := range bot.triggers {
						if t.Matches(text) {
							err := t.Handle(text, bot.rtm, ev)
							if err != nil {
								log.WithError(err).Error("error handling event")
							}
						}
					}

				}
			case *slack.RTMError:
				fmt.Printf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				fmt.Printf("Invalid credentials")
				return errors.New("invalid credentials")
			default:
				// Take no action
			}
		}
	}
}

func (bot *Bot) AddTrigger(s string, f trigger.MessageHander) {
	bot.triggers = append(bot.triggers, trigger.NewTrigger(`ios (?P<Id>\d+)`, f))
}
