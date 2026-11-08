# MatchState

How sure the latest screening is that it found the attendee's accounts. unscreened until a screening reaches the attendee.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.MatchStateUnscreened

// Open enum: custom values can be created with a direct type cast
custom := components.MatchState("custom_value")
```


## Values

| Name                     | Value                    |
| ------------------------ | ------------------------ |
| `MatchStateUnscreened`   | unscreened               |
| `MatchStateConfirmed`    | confirmed                |
| `MatchStateProbable`     | probable                 |
| `MatchStatePossible`     | possible                 |
| `MatchStateNoneFound`    | none_found               |
| `MatchStateInaccessible` | inaccessible             |