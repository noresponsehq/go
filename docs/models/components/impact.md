# Impact

The selected reason's served impact band. Requires reason.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
)

value := components.ImpactNone

// Open enum: custom values can be created with a direct type cast
custom := components.Impact("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `ImpactNone`   | none           |
| `ImpactLow`    | low            |
| `ImpactMedium` | medium         |
| `ImpactHigh`   | high           |