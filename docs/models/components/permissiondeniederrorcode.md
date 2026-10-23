# PermissionDeniedErrorCode

Machine codes an HTTP 403 answer can carry.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v3/models/components"
)

value := components.PermissionDeniedErrorCodeTenantForbidden

// Open enum: custom values can be created with a direct type cast
custom := components.PermissionDeniedErrorCode("custom_value")
```


## Values

| Name                                       | Value                                      |
| ------------------------------------------ | ------------------------------------------ |
| `PermissionDeniedErrorCodeTenantForbidden` | tenant_forbidden                           |