package service

import (
	"testing"

	"github.com/proclaim/mock-slack-api/server"
	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
)

const color = "white"
const fallback = "fallback message"
const text = "secret love letter"
const channel = "core-platform-team"

func TestPostMessageHandler(t *testing.T) {
	mockServer := server.New()

	client := slack.New("SCRET_TOKEN", slack.OptionAPIURL(mockServer.Server.URL+"/"))
	s := NewSlackService(client)

	attachment := slack.Attachment{
		Color:    color,
		Fallback: fallback,
		Text:     text,
	}

	chnl, tstamp, err := s.PostMessage(channel, attachment)

	assert.NoError(t, err, "should not error out")
	assert.Equal(t, chnl, channel, "channel should be correct")
	assert.NotEmpty(t, tstamp, "timestamp should not be empty")

	assert.Equal(t, len(mockServer.Received.Attachment), 1)
	assert.Equal(t, mockServer.Received.Attachment[0].Color, color)
	assert.Equal(t, mockServer.Received.Attachment[0].Fallback, fallback)
	assert.Equal(t, mockServer.Received.Attachment[0].Text, text)
}
