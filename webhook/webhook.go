package webhook

import (
	"github.com/slack-go/slack"
)

func Post(url, text string) error {
	if err := postSlack(url, text); err != nil {
		return err
	}
	return nil
}

var postWebhookFunc = slack.PostWebhook

// postSlack function
func postSlack(url, text string) error {
	payload := &slack.WebhookMessage{
		Blocks: &slack.Blocks{
			BlockSet: []slack.Block{
				getResultSectionBlock(text),
			},
		},
	}

	err := postWebhookFunc(url, payload)
	if err != nil {
		return err
	}

	return nil
}

// getResultSectionBlock function
func getResultSectionBlock(text string) slack.Block {
	textBlockObject := slack.NewTextBlockObject(
		"mrkdwn",
		text,
		false,
		false,
	)
	section := slack.NewSectionBlock(
		textBlockObject,
		nil,
		nil,
	)
	return section
}
