# Role

A built-in capability set. Specialist roles are independent; effective access is returned by GET /v1/workspaces/{workspace_id}/capabilities.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
)

value := components.RoleOwner

// Open enum: custom values can be created with a direct type cast
custom := components.Role("custom_value")
```


## Values

| Name               | Value              |
| ------------------ | ------------------ |
| `RoleOwner`        | owner              |
| `RoleAdmin`        | admin              |
| `RoleDeveloper`    | developer          |
| `RoleViewer`       | viewer             |
| `RoleInvestigator` | investigator       |
| `RoleAuditor`      | auditor            |
| `RoleBilling`      | billing            |