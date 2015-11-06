Trigger a PagerDuty incident if the number of `enqueued` messages in a Sidekiq system is higher than a threshold.

## Usage
```bash
$ HTTP_USERNAME="<basic-auth-username>" \
  HTTP_PASSWORD="<basic-auth-password> \
  THRESHOLD="<number of messages in queue for alert>" \
  PAGERDUTY_KEY="<pagerduty api key>" \
  ./sidekiq-monitor
```
