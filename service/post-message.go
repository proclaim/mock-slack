package service

import (
	"github.com/slack-go/slack"
)

func (s *SlackService) PostMessage(channel string, attachment slack.Attachment) (string, string, error) {
	// ...
	return s.api.PostMessage(
		channel,
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionAsUser(true),
	)
}
