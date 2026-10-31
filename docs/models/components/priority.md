# Priority

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.PriorityLow

// Open enum: custom values can be created with a direct type cast
custom := components.Priority("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `PriorityLow`    | low              |
| `PriorityNormal` | normal           |
| `PriorityHigh`   | high             |
| `PriorityUrgent` | urgent           |