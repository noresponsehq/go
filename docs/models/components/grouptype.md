# GroupType

What the group's people share: an event, an organisation, a place or another attribute.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v19/models/components"
)

value := components.GroupTypeEvent

// Open enum: custom values can be created with a direct type cast
custom := components.GroupType("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `GroupTypeEvent`        | event                   |
| `GroupTypeOrganisation` | organisation            |
| `GroupTypePlace`        | place                   |
| `GroupTypeCustom`       | custom                  |