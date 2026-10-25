# Admission

immediate (the default): an assessment of this flow is decided from the evidence its answer has, and later evidence may change the decision. held: while the account's current assessment of this flow has not completed, the account's decision for the flow is held, and the final decision arrives as account.risk_changed when the investigation completes or hold_seconds pass, whichever is first. Each flow has its own admission.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v7/models/components"
)

value := components.AdmissionImmediate

// Open enum: custom values can be created with a direct type cast
custom := components.Admission("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `AdmissionImmediate` | immediate            |
| `AdmissionHeld`      | held                 |