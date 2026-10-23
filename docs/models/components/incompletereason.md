# IncompleteReason

Why a lane did not settle. deferred means the lane is asynchronous and still running; a continuation settles it. timed_out means the lane missed its deadline. insufficient_evidence is a settled inconclusive signal rather than an incomplete lane.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v3/models/components"
)

value := components.IncompleteReasonTimedOut

// Open enum: custom values can be created with a direct type cast
custom := components.IncompleteReason("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `IncompleteReasonTimedOut`    | timed_out                     |
| `IncompleteReasonRateLimited` | rate_limited                  |
| `IncompleteReasonBlocked`     | blocked                       |
| `IncompleteReasonSourceError` | source_error                  |
| `IncompleteReasonUnavailable` | unavailable                   |
| `IncompleteReasonDeferred`    | deferred                      |