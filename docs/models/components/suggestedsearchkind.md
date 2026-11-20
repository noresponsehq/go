# SuggestedSearchKind

The kind of the value: one of the kinds the environment's enrichment setting names.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v20/models/components"
)

value := components.SuggestedSearchKindUsername

// Open enum: custom values can be created with a direct type cast
custom := components.SuggestedSearchKind("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `SuggestedSearchKindUsername` | username                      |
| `SuggestedSearchKindEmail`    | email                         |
| `SuggestedSearchKindPhone`    | phone                         |