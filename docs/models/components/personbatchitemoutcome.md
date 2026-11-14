# PersonBatchItemOutcome

unchanged: the row matched a person with the same details.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v11/models/components"
)

value := components.PersonBatchItemOutcomeCreated

// Open enum: custom values can be created with a direct type cast
custom := components.PersonBatchItemOutcome("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `PersonBatchItemOutcomeCreated`   | created                           |
| `PersonBatchItemOutcomeUpdated`   | updated                           |
| `PersonBatchItemOutcomeUnchanged` | unchanged                         |