package trigger

import (
	"errors"
	"github.com/nlopes/slack"
	"regexp"
)

type MessageHander func(m map[string]string, rtm *slack.RTM, event *slack.MessageEvent) error

type Trigger struct {
	regexp  *regexp.Regexp
	handler MessageHander
}

func NewTrigger(s string, f MessageHander) *Trigger {
	return &Trigger{
		regexp:  regexp.MustCompile(s),
		handler: f,
	}
}

func (t *Trigger) Matches(s string) bool {
	return t.regexp.MatchString(s)
}

func (t *Trigger) Params(s string) (m map[string]string) {

	match := t.regexp.FindStringSubmatch(s)

	m = make(map[string]string)
	for i, name := range t.regexp.SubexpNames() {
		if i > 0 && i <= len(match) {
			m[name] = match[i]
		}
	}
	return
}

func (t *Trigger) Handle(s string, rtm *slack.RTM, event *slack.MessageEvent) error {
	if !t.Matches(s) {
		return errors.New("trigger not matched")
	}

	params := t.Params(s)
	return t.handler(params, rtm, event)
}
