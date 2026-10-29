# Confidence

Evidence coverage: how much of the planned source evidence gave a definite answer. This is not confidence that an identity is genuine, agreement between identifiers, a risk score or a probability. unavailable means no usable evidence answered; a terminal uncovered assessment can have this value and recommend review.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
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