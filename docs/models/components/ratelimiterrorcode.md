# RateLimitErrorCode

Machine codes an HTTP 429 answer can carry.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v3/models/components"
)

value := components.RateLimitErrorCodeRateLimited

// Open enum: custom values can be created with a direct type cast
custom := components.RateLimitErrorCode("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `RateLimitErrorCodeRateLimited` | rate_limited                    |