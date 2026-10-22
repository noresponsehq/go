# HookPoint

The provider event a provider_hook context was sent for. pre_registration runs before the user exists, so subject and the request's account_id may be absent there. Sent only with transport provider_hook.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v2/models/components"
)

value := components.HookPointPreRegistration
```


## Values

| Name                        | Value                       |
| --------------------------- | --------------------------- |
| `HookPointPreRegistration`  | pre_registration            |
| `HookPointPostRegistration` | post_registration           |
| `HookPointPreLogin`         | pre_login                   |
| `HookPointPostLogin`        | post_login                  |
| `HookPointTokenIssuance`    | token_issuance              |