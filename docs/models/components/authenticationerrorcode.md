# AuthenticationErrorCode

Machine codes an HTTP 401 answer can carry.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v8/models/components"
)

value := components.AuthenticationErrorCodeUnauthorized

// Open enum: custom values can be created with a direct type cast
custom := components.AuthenticationErrorCode("custom_value")
```


## Values

| Name                                  | Value                                 |
| ------------------------------------- | ------------------------------------- |
| `AuthenticationErrorCodeUnauthorized` | unauthorized                          |