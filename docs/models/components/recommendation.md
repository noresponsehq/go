# Recommendation

The recommended action: allow, step_up or review. Your application applies it.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v5/models/components"
)

value := components.RecommendationAllow

// Open enum: custom values can be created with a direct type cast
custom := components.Recommendation("custom_value")
```


## Values

| Name                   | Value                  |
| ---------------------- | ---------------------- |
| `RecommendationAllow`  | allow                  |
| `RecommendationStepUp` | step_up                |
| `RecommendationReview` | review                 |