# BadRequestErrorCode

Machine codes an HTTP 400 answer can carry.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v2/models/components"
)

value := components.BadRequestErrorCodeInvalidRequest

// Open enum: custom values can be created with a direct type cast
custom := components.BadRequestErrorCode("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `BadRequestErrorCodeInvalidRequest`        | invalid_request                            |
| `BadRequestErrorCodeInvalidIdempotencyKey` | invalid_idempotency_key                    |
| `BadRequestErrorCodeInvalidCursor`         | invalid_cursor                             |