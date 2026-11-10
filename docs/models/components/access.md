# Access

public profiles were read; a private or locked one exists but could not be read.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.AccessPublic

// Open enum: custom values can be created with a direct type cast
custom := components.Access("custom_value")
```


## Values

| Name            | Value           |
| --------------- | --------------- |
| `AccessPublic`  | public          |
| `AccessPrivate` | private         |
| `AccessLocked`  | locked          |