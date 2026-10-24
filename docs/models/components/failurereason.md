# FailureReason

Why a signal or source did not produce a definite found/not_found. timed_out means the source missed the request deadline; a continuation may still settle it. deferred means the (source, identifier) pair is async-only and was not attempted on the synchronous path; a continuation settles it.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v7/models/components"
)

value := components.FailureReasonTimedOut

// Open enum: custom values can be created with a direct type cast
custom := components.FailureReason("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `FailureReasonTimedOut`             | timed_out                           |
| `FailureReasonRateLimited`          | rate_limited                        |
| `FailureReasonBlocked`              | blocked                             |
| `FailureReasonSourceError`          | source_error                        |
| `FailureReasonUnavailable`          | unavailable                         |
| `FailureReasonDeferred`             | deferred                            |
| `FailureReasonInsufficientEvidence` | insufficient_evidence               |