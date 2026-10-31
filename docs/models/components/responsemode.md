# ResponseMode

standard returns the consolidated risk response; detailed adds merged provider-neutral facts when the key has detailed scope. There is no raw mode.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.ResponseModeStandard

// Open enum: custom values can be created with a direct type cast
custom := components.ResponseMode("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `ResponseModeStandard` | standard               |
| `ResponseModeDetailed` | detailed               |