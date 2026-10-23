# HookPoint

The provider event a provider_hook context was sent for; absent for any other transport.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v3/models/components"
)

value := components.HookPointPreRegistration

// Open enum: custom values can be created with a direct type cast
custom := components.HookPoint("custom_value")
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `HookPointPreRegistration`  | pre_registration            |
| `HookPointPostRegistration` | post_registration           |
| `HookPointPreLogin`         | pre_login                   |
| `HookPointPostLogin`        | post_login                  |
| `HookPointTokenIssuance`    | token_issuance              |