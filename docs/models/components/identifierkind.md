# IdentifierKind

Which submitted identifier a signal concerns. ip is a selector of breach_exposure only.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v7/models/components"
)

value := components.IdentifierKindEmail

// Open enum: custom values can be created with a direct type cast
custom := components.IdentifierKind("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `IdentifierKindEmail`    | email                    |
| `IdentifierKindUsername` | username                 |
| `IdentifierKindPhone`    | phone                    |
| `IdentifierKindIP`       | ip                       |