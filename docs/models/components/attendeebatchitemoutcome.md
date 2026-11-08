# AttendeeBatchItemOutcome

unchanged: the row matched an attendee with the same details.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.AttendeeBatchItemOutcomeCreated

// Open enum: custom values can be created with a direct type cast
custom := components.AttendeeBatchItemOutcome("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `AttendeeBatchItemOutcomeCreated`   | created                             |
| `AttendeeBatchItemOutcomeUpdated`   | updated                             |
| `AttendeeBatchItemOutcomeUnchanged` | unchanged                           |