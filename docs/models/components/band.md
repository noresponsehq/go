# Band

low 0-29, medium 30-59, high 60-84, critical 85-100.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v2/models/components"
)

value := components.BandLow

// Open enum: custom values can be created with a direct type cast
custom := components.Band("custom_value")
```


## Values

| Name           | Value          |
| -------------- | -------------- |
| `BandLow`      | low            |
| `BandMedium`   | medium         |
| `BandHigh`     | high           |
| `BandCritical` | critical       |