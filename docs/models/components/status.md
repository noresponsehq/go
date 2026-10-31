# Status

Processing state, separate from evidence coverage. complete: final; nothing is still running, but unsupported, failed or unanswered lanes remain in incomplete_sources. partial: a recommendation exists and evidence is incomplete; assessment_url identifies continuing work when present. pending: no recommendation yet; poll or wait for the webhook. Finished processing does not mean all sources answered.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
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