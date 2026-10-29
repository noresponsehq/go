# SignalOutcome

Per-signal outcome. A source that timed out, was blocked or errored reports blocked, unavailable or inconclusive with a reason and is listed in incomplete_sources; not_found is a definite absence.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
)

value := components.SignalOutcomeFound

// Open enum: custom values can be created with a direct type cast
custom := components.SignalOutcome("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `SignalOutcomeFound`        | found                       |
| `SignalOutcomeNotFound`     | not_found                   |
| `SignalOutcomeInconclusive` | inconclusive                |
| `SignalOutcomeBlocked`      | blocked                     |
| `SignalOutcomeUnavailable`  | unavailable                 |