# EnforcementProvider

The identity provider the action ran in; custom for the workspace's own user store.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
)

value := components.EnforcementProviderClerk

// Open enum: custom values can be created with a direct type cast
custom := components.EnforcementProvider("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `EnforcementProviderClerk`  | clerk                       |
| `EnforcementProviderWorkos` | workos                      |
| `EnforcementProviderAuth0`  | auth0                       |
| `EnforcementProviderCustom` | custom                      |