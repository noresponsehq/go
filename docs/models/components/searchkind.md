# SearchKind

The kind of the value searched.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v12/models/components"
)

value := components.SearchKindEmail

// Open enum: custom values can be created with a direct type cast
custom := components.SearchKind("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `SearchKindEmail`      | email                  |
| `SearchKindPhone`      | phone                  |
| `SearchKindUsername`   | username               |
| `SearchKindPersonName` | person_name            |
| `SearchKindPassword`   | password               |
| `SearchKindHash`       | hash                   |
| `SearchKindIP`         | ip                     |