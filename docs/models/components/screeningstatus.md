# ScreeningStatus

partial finished with some attendees not covered; attendees_failed counts them.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
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