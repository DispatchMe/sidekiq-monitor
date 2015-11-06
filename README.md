Trigger a PagerDuty incident if the number of `enqueued` messages in a Sidekiq system is higher than a threshold.

## Usage
```bash
$ SIDEKIQ_URL="<url with basic auth>" \
  THRESHOLD="<number of messages in queue for alert>" \
  PAGERDUTY_KEY="<pagerduty api key>" \
  ./sidekiq-monitor
```
