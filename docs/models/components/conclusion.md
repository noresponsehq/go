# Conclusion

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
)

value := components.ConclusionBenign

// Open enum: custom values can be created with a direct type cast
custom := components.Conclusion("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `ConclusionBenign`       | benign                   |
| `ConclusionAbuse`        | abuse                    |
| `ConclusionInconclusive` | inconclusive             |