package main

import (
	"github.com/proclaim/mock-slack-api/service"
	"github.com/slack-go/slack"
)

func main() {
	slackAPI := slack.New("SECRET_SLACK_TOKEN")
	s := service.NewSlackService(slackAPI)

	attachment := slack.Attachment{
		Fallback: "pasta",
		Text:     "Linguine ai Frutti di Mare",
	}

	s.PostMessage("FOOD_CHANNEL", attachment)
}
