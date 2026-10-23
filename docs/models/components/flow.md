# Flow

The customer event an assessment is made for: signup (a new account), login (a returning sign-in) or purchase (a payment or credit grant). The account's risk is reported with the decision for this flow.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v4/models/components"
)

value := components.FlowSignup

// Open enum: custom values can be created with a direct type cast
custom := components.Flow("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `FlowSignup`   | signup         |
| `FlowLogin`    | login          |
| `FlowPurchase` | purchase       |