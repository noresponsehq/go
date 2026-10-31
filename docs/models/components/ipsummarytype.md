# IPSummaryType

The address's network type.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.IPSummaryTypeResidential

// Open enum: custom values can be created with a direct type cast
custom := components.IPSummaryType("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `IPSummaryTypeResidential` | residential                |
| `IPSummaryTypeHosting`     | hosting                    |
| `IPSummaryTypeVpn`         | vpn                        |
| `IPSummaryTypeProxy`       | proxy                      |
| `IPSummaryTypeTor`         | tor                        |
| `IPSummaryTypeMobile`      | mobile                     |
| `IPSummaryTypeUnknown`     | unknown                    |