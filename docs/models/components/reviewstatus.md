# ReviewStatus

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.ReviewStatusOpen

// Open enum: custom values can be created with a direct type cast
custom := components.ReviewStatus("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ReviewStatusOpen`       | open                     |
| `ReviewStatusInProgress` | in_progress              |
| `ReviewStatusResolved`   | resolved                 |