# WebhookEndpointState

unverified until the endpoint echoes an endpoint.verification challenge (send one with POST /v1/webhooks/{endpoint_id}/test); verified receives assessment.updated events; disabled after 72 hours of continuous delivery failure, and verified again by the same challenge.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v2/models/components"
)

value := components.WebhookEndpointStateUnverified

// Open enum: custom values can be created with a direct type cast
custom := components.WebhookEndpointState("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `WebhookEndpointStateUnverified` | unverified                       |
| `WebhookEndpointStateVerified`   | verified                         |
| `WebhookEndpointStateDisabled`   | disabled                         |