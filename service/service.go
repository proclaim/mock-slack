package service

import "github.com/slack-go/slack"

type SlackAPI interface {
	PostMessage(channelID string, options ...slack.MsgOption) (string, string, error)
}

type SlackService struct {
	api SlackAPI
}

func NewSlackService(api SlackAPI) *SlackService {
	return &SlackService{
		api: api,
	}
}
