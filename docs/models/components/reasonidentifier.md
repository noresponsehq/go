# ReasonIdentifier

The submitted identifier or client context field a reason concerns; absent when it concerns the whole assessment. Client context is a claimed hint, not a device fingerprint.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v4/models/components"
)

value := components.ReasonIdentifierEmail

// Open enum: custom values can be created with a direct type cast
custom := components.ReasonIdentifier("custom_value")
```


## Values

| Name                       | Value                      |
| -------------------------- | -------------------------- |
| `ReasonIdentifierEmail`    | email                      |
| `ReasonIdentifierUsername` | username                   |
| `ReasonIdentifierPhone`    | phone                      |
| `ReasonIdentifierName`     | name                       |
| `ReasonIdentifierIP`       | ip                         |
| `ReasonIdentifierContext`  | context                    |