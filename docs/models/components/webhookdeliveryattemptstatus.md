# WebhookDeliveryAttemptStatus

delivered when the endpoint answered 2xx; failed otherwise, with failure naming why.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
)

value := components.WebhookDeliveryAttemptStatusDelivered

// Open enum: custom values can be created with a direct type cast
custom := components.WebhookDeliveryAttemptStatus("custom_value")
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `WebhookDeliveryAttemptStatusDelivered` | delivered                               |
| `WebhookDeliveryAttemptStatusFailed`    | failed                                  |