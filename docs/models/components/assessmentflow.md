# AssessmentFlow

The customer event an assessment is made for: signup (a new account), login (a returning sign-in) or purchase (a payment or credit grant). The account's risk is reported with the decision for this flow.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v2/models/components"
)

value := components.AssessmentFlowSignup
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `AssessmentFlowSignup`   | signup                   |
| `AssessmentFlowLogin`    | login                    |
| `AssessmentFlowPurchase` | purchase                 |