# SearchStatus

A password or hash search is counting, then counted: rows says how many breach rows hold the value, and nothing runs until a person runs it (POST .../searches/{search_id}/run). Any other search is queued, then running. completed: every source it asked answered. partial: some did not. failed: none answered or none could be asked, which says nothing about whether the value is held anywhere.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v14/models/components"
)

value := components.SearchStatusCounting

// Open enum: custom values can be created with a direct type cast
custom := components.SearchStatus("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `SearchStatusCounting`  | counting                |
| `SearchStatusCounted`   | counted                 |
| `SearchStatusQueued`    | queued                  |
| `SearchStatusRunning`   | running                 |
| `SearchStatusCompleted` | completed               |
| `SearchStatusPartial`   | partial                 |
| `SearchStatusFailed`    | failed                  |