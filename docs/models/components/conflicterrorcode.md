# ConflictErrorCode

Machine codes an HTTP 409 answer can carry.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v2/models/components"
)

value := components.ConflictErrorCodeIdempotencyKeyReused

// Open enum: custom values can be created with a direct type cast
custom := components.ConflictErrorCode("custom_value")
```


## Values

| Name                                     | Value                                    |
| ---------------------------------------- | ---------------------------------------- |
| `ConflictErrorCodeIdempotencyKeyReused`  | idempotency_key_reused                   |
| `ConflictErrorCodeKeyRotationInProgress` | key_rotation_in_progress                 |
| `ConflictErrorCodeTooManyIntegrations`   | too_many_integrations                    |