package slackhandler

import (
	"fmt"
	"github.com/nlopes/slack"
	"github.com/riksa/gonkka/pkg/gocd"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

type SlackHandler struct {
	rtm  *slack.RTM
	rest *gocd.Rest
}

func New(rtm *slack.RTM, rest *gocd.Rest) *SlackHandler {
	return &SlackHandler{
		rtm:  rtm,
		rest: rest,
	}
}

func (sh *SlackHandler) HandleIosInstance(m map[string]string, event *slack.MessageEvent) error {
	return sh.handleInstance("ios-build-app", m, event)
}

func (sh *SlackHandler) HandleAndroidInstance(m map[string]string, event *slack.MessageEvent) error {
	return sh.handleInstance("android-build-app", m, event)
}

func (sh *SlackHandler) handleInstance(app string, m map[string]string, event *slack.MessageEvent) error {
	id, err := strconv.ParseInt(m["Id"], 10, 32)
	if err != nil {
		return err
	}
	log.WithField("Id", id).WithField("app", app).Debug("handle instance")

	response, err := sh.rest.Instance(app, id)
	if err != nil {
		sh.rtm.SendMessage(sh.rtm.NewOutgoingMessage(fmt.Sprintf("cannot find %s version %d", app, id), event.Channel))
		return err
	}

	bytes, err := response.MarshalBinary()
	log.WithField("json", string(bytes)).Debug("received instance response")

	sb := strings.Builder{}
	comments := make([]string, 0)
	sb.WriteString(fmt.Sprintf("Build name: %s, number: %d\n", response.Name, id))
	for _, rev := range response.BuildCause.MaterialRevisions {
		for _, mod := range rev.Modifications {
			comments = append(comments, mod.Comment)
		}
	}
	sb.WriteString("Commits:\n")
	sb.WriteString(strings.Join(comments, "\n"))

	msg := sb.String()
	sh.rtm.SendMessage(sh.rtm.NewOutgoingMessage(msg, event.Channel))

	return nil
}
