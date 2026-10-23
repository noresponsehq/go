# IntegrationHookPoint

pre_registration runs before the user exists; post_login runs after a successful sign-in.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v5/models/components"
)

value := components.IntegrationHookPointPreRegistration

// Open enum: custom values can be created with a direct type cast
custom := components.IntegrationHookPoint("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `IntegrationHookPointPreRegistration` | pre_registration                      |
| `IntegrationHookPointPostLogin`       | post_login                            |