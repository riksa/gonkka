package main

import (
	"context"
	"github.com/riksa/gonkka/internal/slackbot"
	"github.com/riksa/gonkka/internal/slackhandler"
	"github.com/riksa/gonkka/pkg/gocd"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	log.SetLevel(log.DebugLevel)
	token := os.Getenv("SLACKTOKEN")
	if token == "" {
		log.Fatal("missing SLACKTOKEN=xoxb-... environment variable")
	}

	host := os.Getenv("GOCD_HOST")
	if host == "" {
		log.Fatal("missing GOCD_HOST=jenkins2.local:8153 environment variable")
	}

	rest := gocd.NewRest(host, context.TODO())

	bot := slackbot.New(token)
	sh := slackhandler.New(bot.Rtm, rest)
	bot.AddTrigger(`ios (?P<Id>\d+)`, sh.HandleIosInstance)
	bot.AddTrigger(`android (?P<Id>\d+)`, sh.HandleAndroidInstance)
	log.WithError(bot.Run()).Fatal("bot is dead")
}
