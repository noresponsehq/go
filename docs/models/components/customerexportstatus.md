# CustomerExportStatus

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.CustomerExportStatusQueued

// Open enum: custom values can be created with a direct type cast
custom := components.CustomerExportStatus("custom_value")
```


## Values

| Name                            | Value                           |
| ------------------------------- | ------------------------------- |
| `CustomerExportStatusQueued`    | queued                          |
| `CustomerExportStatusRunning`   | running                         |
| `CustomerExportStatusCompleted` | completed                       |
| `CustomerExportStatusFailed`    | failed                          |
| `CustomerExportStatusCancelled` | cancelled                       |
| `CustomerExportStatusExpired`   | expired                         |