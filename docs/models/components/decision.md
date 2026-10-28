# Decision

The decision of the assessment's latest stored revision: block when its stored risk.block is true, otherwise its recommendation. This is not the account's current decision or held-admission state. Combined with recommendation using AND.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v8/models/components"
)

value := components.DecisionAllow
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `DecisionAllow`  | allow            |
| `DecisionStepUp` | step_up          |
| `DecisionReview` | review           |
| `DecisionBlock`  | block            |