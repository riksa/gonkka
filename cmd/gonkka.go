package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"github.com/riksa/gonkka/internal/slackbot"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
)

func main() {
	log.SetLevel(log.DebugLevel)
	token := os.Getenv("SLACKTOKEN")
	if token == "" {
		log.Fatal("missing SLACKTOKEN environment variable")
	}

	bot := slackbot.New(token)
	bot.AddTrigger(`ios (?P<Id>\d+)`, handleIosInstance)
	log.WithError(bot.Run()).Fatal("bot is died")
}

func handleIosInstance(m map[string]string, rtm *slack.RTM, event *slack.MessageEvent) error {
	id, err := strconv.ParseInt(m["Id"], 10, 32)
	if err != nil {
		return err
	}
	log.WithField("Id", id).Debug("handle ios instance")
	msg := fmt.Sprintf("TODO: something about ios instance %d", id)
	rtm.SendMessage(rtm.NewOutgoingMessage(msg, event.Channel))
	return nil
}
