# Outcome

What the provider later learned about the assessed account. confirmed_abuse, account_takeover and chargeback report abuse; false_positive and legitimate withdraw an earlier abuse report of this workspace for the assessment; unknown is recorded only. Outcomes are append-only and the latest one recorded for an assessment wins.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v8/models/components"
)

value := components.OutcomeConfirmedAbuse

// Open enum: custom values can be created with a direct type cast
custom := components.Outcome("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `OutcomeConfirmedAbuse`  | confirmed_abuse          |
| `OutcomeAccountTakeover` | account_takeover         |
| `OutcomeChargeback`      | chargeback               |
| `OutcomeFalsePositive`   | false_positive           |
| `OutcomeLegitimate`      | legitimate               |
| `OutcomeUnknown`         | unknown                  |