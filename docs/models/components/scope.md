# Scope

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v7/models/components"
)

value := components.ScopeAssess

// Open enum: custom values can be created with a direct type cast
custom := components.Scope("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `ScopeAssess`       | assess              |
| `ScopeHistory`      | history             |
| `ScopeDelete`       | delete              |
| `ScopeWebhooks`     | webhooks            |
| `ScopeDetailed`     | detailed            |
| `ScopeIntegrations` | integrations        |