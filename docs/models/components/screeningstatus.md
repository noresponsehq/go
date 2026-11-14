# ScreeningStatus

partial finished with some people or platforms not covered; people_failed and people_partial count them.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v11/models/components"
)

value := components.ScreeningStatusQueued

// Open enum: custom values can be created with a direct type cast
custom := components.ScreeningStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ScreeningStatusQueued`    | queued                     |
| `ScreeningStatusRunning`   | running                    |
| `ScreeningStatusCompleted` | completed                  |
| `ScreeningStatusPartial`   | partial                    |
| `ScreeningStatusFailed`    | failed                     |
| `ScreeningStatusCancelled` | cancelled                  |