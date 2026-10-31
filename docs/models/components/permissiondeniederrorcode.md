# PermissionDeniedErrorCode

Machine codes an HTTP 403 answer can carry.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.PermissionDeniedErrorCodeWorkspaceForbidden

// Open enum: custom values can be created with a direct type cast
custom := components.PermissionDeniedErrorCode("custom_value")
```


## Values

| Name                                          | Value                                         |
| --------------------------------------------- | --------------------------------------------- |
| `PermissionDeniedErrorCodeWorkspaceForbidden` | workspace_forbidden                           |
| `PermissionDeniedErrorCodeRoleForbidden`      | role_forbidden                                |
| `PermissionDeniedErrorCodeEmailUnverified`    | email_unverified                              |
| `PermissionDeniedErrorCodeOnboardingInReview` | onboarding_in_review                          |
| `PermissionDeniedErrorCodeOnboardingRefused`  | onboarding_refused                            |