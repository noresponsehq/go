# Delay

How long after the assessment's creation the change was made.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.DelayUnder10s

// Open enum: custom values can be created with a direct type cast
custom := components.Delay("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `DelayUnder10s` | under_10s       |
| `DelayUnder1m`  | under_1m        |
| `DelayUnder10m` | under_10m       |
| `DelayUnder1h`  | under_1h        |
| `DelayUnder1d`  | under_1d        |
| `DelayOver1d`   | over_1d         |