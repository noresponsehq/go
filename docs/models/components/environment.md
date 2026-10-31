# Environment

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.EnvironmentLive

// Open enum: custom values can be created with a direct type cast
custom := components.Environment("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `EnvironmentLive`      | live                   |
| `EnvironmentTest`      | test                   |
| `EnvironmentWorkspace` | workspace              |