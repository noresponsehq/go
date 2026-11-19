# WrittenProfileStatus

queued: the profile is being written. written: summary, statements and leads hold it. failed: it could not be written; error_code says why and no content is returned.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v16/models/components"
)

value := components.WrittenProfileStatusQueued

// Open enum: custom values can be created with a direct type cast
custom := components.WrittenProfileStatus("custom_value")
```


## Values

| Name                          | Value                         |
| ----------------------------- | ----------------------------- |
| `WrittenProfileStatusQueued`  | queued                        |
| `WrittenProfileStatusWritten` | written                       |
| `WrittenProfileStatusFailed`  | failed                        |