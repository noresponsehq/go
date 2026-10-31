# RiskConfidence

How much of the planned evidence answered; it says nothing about risk.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.RiskConfidenceLow

// Open enum: custom values can be created with a direct type cast
custom := components.RiskConfidence("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `RiskConfidenceLow`    | low                    |
| `RiskConfidenceMedium` | medium                 |
| `RiskConfidenceHigh`   | high                   |