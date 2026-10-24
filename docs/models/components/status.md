# Status

complete: final; nothing is still running. A lane that failed or missed its 10 s deadline stays listed in incomplete_sources, and confidence is lower. partial: a recommendation exists and at least one lane is still running or did not settle (see incomplete_sources). pending: nothing usable before the deadline; no recommendation yet, poll or wait for the webhook.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v7/models/components"
)

value := components.StatusComplete

// Open enum: custom values can be created with a direct type cast
custom := components.Status("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `StatusComplete` | complete         |
| `StatusPartial`  | partial          |
| `StatusPending`  | pending          |