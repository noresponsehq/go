# Decision

An assessment's own decision: block when its stored risk.block is true, otherwise its recommendation. It is not the account's decision or held-admission state.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.DecisionAllow

// Open enum: custom values can be created with a direct type cast
custom := components.Decision("custom_value")
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `DecisionAllow`  | allow            |
| `DecisionStepUp` | step_up          |
| `DecisionReview` | review           |
| `DecisionBlock`  | block            |