# Agreement

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v8/models/components"
)

value := components.AgreementMatch

// Open enum: custom values can be created with a direct type cast
custom := components.Agreement("custom_value")
```


## Values

| Name                | Value               |
| ------------------- | ------------------- |
| `AgreementMatch`    | match               |
| `AgreementPartial`  | partial             |
| `AgreementMismatch` | mismatch            |