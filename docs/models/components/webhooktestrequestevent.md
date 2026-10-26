# WebhookTestRequestEvent

endpoint.verification sends a fresh challenge now, in this call, and the endpoint is verified if it echoes it. assessment.updated queues an event whose data is a fixed example assessment, with test true, delivered and retried like any other; the endpoint must be verified.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v8/models/components"
)

value := components.WebhookTestRequestEventEndpointVerification

// Open enum: custom values can be created with a direct type cast
custom := components.WebhookTestRequestEvent("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `WebhookTestRequestEventEndpointVerification` | endpoint.verification                         |
| `WebhookTestRequestEventAssessmentUpdated`    | assessment.updated                            |