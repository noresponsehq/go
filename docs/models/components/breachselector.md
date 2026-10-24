# BreachSelector

The submitted identifier a breach table matched on: email or username. Phone and name are not breach selectors.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v6/models/components"
)

value := components.BreachSelectorEmail

// Open enum: custom values can be created with a direct type cast
custom := components.BreachSelector("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `BreachSelectorEmail`    | email                    |
| `BreachSelectorUsername` | username                 |
| `BreachSelectorIP`       | ip                       |