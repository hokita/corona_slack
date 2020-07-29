package webhook

import (
	"errors"
	"testing"

	"github.com/slack-go/slack"
)

type mockSlack struct {
	err error
}

func (c mockSlack) PostWebhook(url string, msg *slack.WebhookMessage) error {
	return c.err
}

func TestPost(t *testing.T) {
	postErr := errors.New("post error")

	tests := map[string]struct {
		err  error
		want error
	}{
		"success": {
			err:  nil,
			want: nil,
		},
		"error": {
			err:  postErr,
			want: postErr,
		},
	}

	for name, test := range tests {
		postWebhookFunc = (mockSlack{test.err}).PostWebhook

		t.Run(name, func(t *testing.T) {
			if err := Post("test.com", "hogehoge"); err != test.want {
				t.Errorf(`want="%v" err="%v"`, test.want, err)
			}
		})
	}
}
