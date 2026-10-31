# From

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.FromAllow

// Open enum: custom values can be created with a direct type cast
custom := components.From("custom_value")
```


## Values

| Name         | Value        |
| ------------ | ------------ |
| `FromAllow`  | allow        |
| `FromStepUp` | step_up      |
| `FromReview` | review       |
| `FromBlock`  | block        |
| `FromHeld`   | held         |