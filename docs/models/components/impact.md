# Impact

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v7/models/components"
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