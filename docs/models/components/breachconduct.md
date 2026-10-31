# BreachConduct

What a breach table's rows record against the identifier: chargeback, one or more chargebacks; declined_payment, one or more declined payments; forum_sanction, a forum ban, warning or infraction. A category and not a value: the schema has no member for one.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.BreachConductChargeback

// Open enum: custom values can be created with a direct type cast
custom := components.BreachConduct("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `BreachConductChargeback`      | chargeback                     |
| `BreachConductDeclinedPayment` | declined_payment               |
| `BreachConductForumSanction`   | forum_sanction                 |