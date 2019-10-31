package main

import (
	"context"
	"fmt"
	"github.com/nlopes/slack"
	"github.com/riksa/gonkka/internal/slackbot"
	"github.com/riksa/gonkka/pkg/gocd"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"strings"
)

var(
	rest *gocd.Rest
)


func main() {
	log.SetLevel(log.DebugLevel)
	token := os.Getenv("SLACKTOKEN")
	if token == "" {
		log.Fatal("missing SLACKTOKEN environment variable")
	}
	rest = gocd.NewRest("jenkins2.local:8153", context.TODO())

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

	app := "ios-build-app"
	response, err := rest.Instance(app, id)
	if err != nil {
		rtm.SendMessage(rtm.NewOutgoingMessage(fmt.Sprintf("cannot find %s version %d", app, id), event.Channel))
		return err
	}

	bytes, err := response.MarshalBinary()
	log.WithField("json", string(bytes)).Debug("received instance response")

	sb := strings.Builder{}
	comments := make([]string, 0)
	sb.WriteString( fmt.Sprintf("Build name: %s, number: %d\n", response.Name, id))
	for _, rev := range response.BuildCause.MaterialRevisions {
		for _, mod := range rev.Modifications {
			comments = append(comments, mod.Comment)
		}
	}
	sb.WriteString("Commits:\n")
	sb.WriteString( strings.Join(comments, "\n"))

	msg := sb.String()
	rtm.SendMessage(rtm.NewOutgoingMessage(msg, event.Channel))

	return nil
}
