# PersonFlaggedMatchState

How sure the screening is that the accounts it read are the person's. A finding on a possible match is at most moderate, on a probable match at most high.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v11/models/components"
)

value := components.PersonFlaggedMatchStateConfirmed
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `PersonFlaggedMatchStateConfirmed`    | confirmed                             |
| `PersonFlaggedMatchStateProbable`     | probable                              |
| `PersonFlaggedMatchStatePossible`     | possible                              |
| `PersonFlaggedMatchStateNoneFound`    | none_found                            |
| `PersonFlaggedMatchStateInaccessible` | inaccessible                          |