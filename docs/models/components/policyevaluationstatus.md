# PolicyEvaluationStatus

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.PolicyEvaluationStatusQueued

// Open enum: custom values can be created with a direct type cast
custom := components.PolicyEvaluationStatus("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PolicyEvaluationStatusQueued`    | queued                            |
| `PolicyEvaluationStatusRunning`   | running                           |
| `PolicyEvaluationStatusCompleted` | completed                         |
| `PolicyEvaluationStatusFailed`    | failed                            |
| `PolicyEvaluationStatusCancelled` | cancelled                         |