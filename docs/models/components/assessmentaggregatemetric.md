# AssessmentAggregateMetric

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.AssessmentAggregateMetricCount

// Open enum: custom values can be created with a direct type cast
custom := components.AssessmentAggregateMetric("custom_value")
```


## Values

| Name                                      | Value                                     |
| ----------------------------------------- | ----------------------------------------- |
| `AssessmentAggregateMetricCount`          | count                                     |
| `AssessmentAggregateMetricScoreHistogram` | score_histogram                           |
| `AssessmentAggregateMetricScoreOutcomes`  | score_outcomes                            |
| `AssessmentAggregateMetricLateFlips`      | late_flips                                |