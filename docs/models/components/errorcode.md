# ErrorCode

Why a failed search found nothing: its sources did not answer, none could be asked, or it could not run before its attempts were spent.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v12/models/components"
)

value := components.ErrorCodeNotAnswered

// Open enum: custom values can be created with a direct type cast
custom := components.ErrorCode("custom_value")
```


## Values

| Name                               | Value                              |
| ---------------------------------- | ---------------------------------- |
| `ErrorCodeNotAnswered`             | not_answered                       |
| `ErrorCodeNoSource`                | no_source                          |
| `ErrorCodeWorkerAttemptsExhausted` | worker_attempts_exhausted          |