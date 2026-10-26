# BreachSelector

The submitted email or username whose own search matched the table. A submitted full name is searched as well, and what it finds feeds the identity's reasons rather than this list.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v8/models/components"
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