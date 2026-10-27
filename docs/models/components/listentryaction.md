# ListEntryAction

block decides block for the accounts the entry matches; allow decides allow. Block wins when both match.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v8/models/components"
)

value := components.ListEntryActionBlock

// Open enum: custom values can be created with a direct type cast
custom := components.ListEntryAction("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `ListEntryActionBlock` | block                  |
| `ListEntryActionAllow` | allow                  |