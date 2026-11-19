# InvestigationStatus

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v16/models/components"
)

value := components.InvestigationStatusQueued

// Open enum: custom values can be created with a direct type cast
custom := components.InvestigationStatus("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `InvestigationStatusQueued`    | queued                         |
| `InvestigationStatusRunning`   | running                        |
| `InvestigationStatusCompleted` | completed                      |
| `InvestigationStatusPartial`   | partial                        |
| `InvestigationStatusFailed`    | failed                         |
| `InvestigationStatusCancelled` | cancelled                      |
| `InvestigationStatusUnknown`   | unknown                        |