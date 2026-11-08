# AttendeeFlaggedMatchState

How sure the screening is that the accounts it read are the attendee's. A finding on a possible match is at most moderate, on a probable match at most high.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.AttendeeFlaggedMatchStateConfirmed
```


## Values

| Name                                    | Value                                   |
| --------------------------------------- | --------------------------------------- |
| `AttendeeFlaggedMatchStateConfirmed`    | confirmed                               |
| `AttendeeFlaggedMatchStateProbable`     | probable                                |
| `AttendeeFlaggedMatchStatePossible`     | possible                                |
| `AttendeeFlaggedMatchStateNoneFound`    | none_found                              |
| `AttendeeFlaggedMatchStateInaccessible` | inaccessible                            |