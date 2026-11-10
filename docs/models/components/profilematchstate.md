# ProfileMatchState

How sure the screening is that a profile is the attendee's, from its evidence: confirmed is a handle on the attendee row, or two or more hard identifiers (email, phone, a supplied handle, or a fact an existing source returned); probable is one identifier beyond the name; possible is the name alone.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
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