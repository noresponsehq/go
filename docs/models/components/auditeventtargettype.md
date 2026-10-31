# AuditEventTargetType

The kind of resource the change acted on.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.AuditEventTargetTypeWorkspace

// Open enum: custom values can be created with a direct type cast
custom := components.AuditEventTargetType("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `AuditEventTargetTypeWorkspace`       | workspace                             |
| `AuditEventTargetTypeMember`          | member                                |
| `AuditEventTargetTypeInvite`          | invite                                |
| `AuditEventTargetTypeKey`             | key                                   |
| `AuditEventTargetTypeSetting`         | setting                               |
| `AuditEventTargetTypeWebhookEndpoint` | webhook_endpoint                      |
| `AuditEventTargetTypeIntegration`     | integration                           |
| `AuditEventTargetTypeBilling`         | billing                               |
| `AuditEventTargetTypeReview`          | review                                |