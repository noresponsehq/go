# Confidence

Ordinal confidence in the recommendation under the current policy: low, medium, high or unavailable.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v2/models/components"
)

value := components.ConfidenceLow

// Open enum: custom values can be created with a direct type cast
custom := components.Confidence("custom_value")
```


## Values

| Name                    | Value                   |
| ----------------------- | ----------------------- |
| `ConfidenceLow`         | low                     |
| `ConfidenceMedium`      | medium                  |
| `ConfidenceHigh`        | high                    |
| `ConfidenceUnavailable` | unavailable             |