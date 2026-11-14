# FindingBand

How severe a finding is, most severe first.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v11/models/components"
)

value := components.FindingBandCritical

// Open enum: custom values can be created with a direct type cast
custom := components.FindingBand("custom_value")
```


## Values

| Name                  | Value                 |
| --------------------- | --------------------- |
| `FindingBandCritical` | critical              |
| `FindingBandHigh`     | high                  |
| `FindingBandModerate` | moderate              |
| `FindingBandLow`      | low                   |
| `FindingBandInfo`     | info                  |