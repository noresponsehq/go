# ProfileMatchState

How sure the screening is that a profile is the person's, from the likelihood its record scores. Confirmed reaches the join bar of 0.9. Probable reaches 0.85, which takes a shared value beyond the one that found the account. Possible is the rest.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v20/models/components"
)

value := components.ProfileMatchStateConfirmed

// Open enum: custom values can be created with a direct type cast
custom := components.ProfileMatchState("custom_value")
```


## Values

| Name                         | Value                        |
| ---------------------------- | ---------------------------- |
| `ProfileMatchStateConfirmed` | confirmed                    |
| `ProfileMatchStateProbable`  | probable                     |
| `ProfileMatchStatePossible`  | possible                     |