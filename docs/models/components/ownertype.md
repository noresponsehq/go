# OwnerType

The credential namespace of the creator. An external user ID equal to an API-key label does not grant ownership of that key's private view.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v9/models/components"
)

value := components.OwnerTypeUser

// Open enum: custom values can be created with a direct type cast
custom := components.OwnerType("custom_value")
```


## Values

| Name              | Value             |
| ----------------- | ----------------- |
| `OwnerTypeUser`   | user              |
| `OwnerTypeAPIKey` | api_key           |