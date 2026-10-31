# EnforcementStatus

applied when the identity provider accepted the action, failed when it refused it or could not be reached. Only applied counts as enforced.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.EnforcementStatusApplied

// Open enum: custom values can be created with a direct type cast
custom := components.EnforcementStatus("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `EnforcementStatusApplied` | applied                    |
| `EnforcementStatusFailed`  | failed                     |