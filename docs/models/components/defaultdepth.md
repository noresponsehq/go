# DefaultDepth

The depth a request for this flow that names none is served at. Defaults: full for account_creation and purchase, quick for login.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v8/models/components"
)

value := components.DefaultDepthQuick

// Open enum: custom values can be created with a direct type cast
custom := components.DefaultDepth("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `DefaultDepthQuick` | quick               |
| `DefaultDepthFull`  | full                |