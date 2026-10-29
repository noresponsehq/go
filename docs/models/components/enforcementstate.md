# EnforcementState

The account's state in the workspace's identity provider just before the action: onboarding, provisional (created, not yet activated), active, restricted or rejected.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
)

value := components.EnforcementStateOnboarding

// Open enum: custom values can be created with a direct type cast
custom := components.EnforcementState("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `EnforcementStateOnboarding`  | onboarding                    |
| `EnforcementStateProvisional` | provisional                   |
| `EnforcementStateActive`      | active                        |
| `EnforcementStateRestricted`  | restricted                    |
| `EnforcementStateRejected`    | rejected                      |