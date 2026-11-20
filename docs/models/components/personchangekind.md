# PersonChangeKind

account_added and account_removed: an account or profile appeared or disappeared. breach_added and breach_removed: a breach record appeared or disappeared. identifier_added and identifier_removed: an email, phone, username, name or account id on an account the person already held. field_changed: another field on such an account. posts_added: new posts, counted in count.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v20/models/components"
)

value := components.PersonChangeKindAccountAdded

// Open enum: custom values can be created with a direct type cast
custom := components.PersonChangeKind("custom_value")
```


## Values

| Name                                | Value                               |
| ----------------------------------- | ----------------------------------- |
| `PersonChangeKindAccountAdded`      | account_added                       |
| `PersonChangeKindAccountRemoved`    | account_removed                     |
| `PersonChangeKindBreachAdded`       | breach_added                        |
| `PersonChangeKindBreachRemoved`     | breach_removed                      |
| `PersonChangeKindIdentifierAdded`   | identifier_added                    |
| `PersonChangeKindIdentifierRemoved` | identifier_removed                  |
| `PersonChangeKindFieldChanged`      | field_changed                       |
| `PersonChangeKindPostsAdded`        | posts_added                         |