# Flow

The event an assessment is made for: account_creation (a user creates an account), login (a user returns to an existing account) or purchase (a payment or a credit grant). The account's risk is reported with the decision for this flow.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.FlowAccountCreation

// Open enum: custom values can be created with a direct type cast
custom := components.Flow("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `FlowAccountCreation` | account_creation      |
| `FlowLogin`           | login                 |
| `FlowPurchase`        | purchase              |