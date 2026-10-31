# AccountType

personal, or business: an account a company or school manages.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.AccountTypePersonal

// Open enum: custom values can be created with a direct type cast
custom := components.AccountType("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `AccountTypePersonal` | personal              |
| `AccountTypeBusiness` | business              |