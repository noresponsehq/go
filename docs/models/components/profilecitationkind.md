# ProfileCitationKind

record: a record in lookup.records or profiles[].record. post: an item in posts. check: an item in lookup.checks.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v16/models/components"
)

value := components.ProfileCitationKindRecord

// Open enum: custom values can be created with a direct type cast
custom := components.ProfileCitationKind("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `ProfileCitationKindRecord` | record                      |
| `ProfileCitationKindPost`   | post                        |
| `ProfileCitationKindCheck`  | check                       |