# AccountDecision

block when the score reached your block threshold (risk.block of the assessment), otherwise the assessment's recommendation. held only for a flow whose admission is held (settings flows.<flow>.admission), when the account's current assessment was made for that flow and its investigation has not completed (status partial): the final decision arrives as account.risk_changed when the investigation completes, decided by its score, or when the hold's deadline (hold_seconds after the assessment was created) passes first, as the flow's on_expiry action or the decision the evidence so far gives, whichever is stricter. Your application applies it.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v7/models/components"
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
| `AccountDecisionHeld`   | held                    |