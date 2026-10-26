# ConsistencyValue

How a submitted identifier agrees with what the evidence holds for the other identifiers. unknown is a definite 'cannot tell' (for example a truncated or too-common breach result); absent means the comparison was not made.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v8/models/components"
)

value := components.ConsistencyValueMatch

// Open enum: custom values can be created with a direct type cast
custom := components.ConsistencyValue("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ConsistencyValueMatch`    | match                      |
| `ConsistencyValuePartial`  | partial                    |
| `ConsistencyValueMismatch` | mismatch                   |
| `ConsistencyValueUnknown`  | unknown                    |