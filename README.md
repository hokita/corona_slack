# corona_slack
This command sends the number of corona-infected people in Japan and the countries with the highest number of corona-infected people to Slack.

Source: https://disease.sh/

About slack webhook: https://api.slack.com/messaging/webhooks

## usage
```shell
$ export SLACK_WEBHOOK_URL=<slack webhoook url>

$ docker build --build-arg SLACK_WEBHOOK_URL=$SLACK_WEBHOOK_URL -t corona_slack .

$ docker run corona_slack
```
