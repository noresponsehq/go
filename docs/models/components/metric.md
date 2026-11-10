# Metric

assessments: assessments admitted. webhook_deliveries: events sent to endpoints, one per event per endpoint. key_usage_events: key-usage events accepted. screenings: attendees a screening recorded a final result for, one per attendee per screening; a failed attendee counts only when its lookups were kept.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.MetricAssessments

// Open enum: custom values can be created with a direct type cast
custom := components.Metric("custom_value")
```


## Values

| Name                      | Value                     |
| ------------------------- | ------------------------- |
| `MetricAssessments`       | assessments               |
| `MetricWebhookDeliveries` | webhook_deliveries        |
| `MetricKeyUsageEvents`    | key_usage_events          |
| `MetricScreenings`        | screenings                |