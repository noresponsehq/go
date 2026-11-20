# LastError

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v20/models/components"
)

value := components.LastErrorScreeningCapExceeded

// Open enum: custom values can be created with a direct type cast
custom := components.LastError("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `LastErrorScreeningCapExceeded` | screening_cap_exceeded          |
| `LastErrorScreeningUnavailable` | screening_unavailable           |
| `LastErrorPersonNotFound`       | person_not_found                |
| `LastErrorIdempotencyKeyReused` | idempotency_key_reused          |