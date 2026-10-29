# Population

first_answers uses anonymous counters that can include erased detail. retained_first_answers computes exact intersections over retained first answers within 31 days. Requests exceeding the 100000-record work budget are refused.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
)

value := components.PopulationFirstAnswers

// Open enum: custom values can be created with a direct type cast
custom := components.Population("custom_value")
```


## Values

| Name                             | Value                            |
| -------------------------------- | -------------------------------- |
| `PopulationFirstAnswers`         | first_answers                    |
| `PopulationRetainedFirstAnswers` | retained_first_answers           |