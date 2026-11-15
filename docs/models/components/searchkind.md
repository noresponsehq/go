# SearchKind

The kind of the value searched. A domain search finds the leaked rows of every email address at the domain.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v14/models/components"
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
| `SearchKindDomain`     | domain                 |