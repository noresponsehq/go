# AccountDecision

block when the score reached your block threshold (risk.block of the assessment), otherwise the assessment's recommendation. Your application applies it.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v6/models/components"
)

value := components.AccountDecisionAllow

// Open enum: custom values can be created with a direct type cast
custom := components.AccountDecision("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `AccountDecisionAllow`  | allow                   |
| `AccountDecisionStepUp` | step_up                 |
| `AccountDecisionReview` | review                  |
| `AccountDecisionBlock`  | block                   |