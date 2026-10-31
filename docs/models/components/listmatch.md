# ListMatch

The list action captured at the first answer. Requires revision first. Older answers without captured list state do not match, including none.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.ListMatchAllow
```


## Values

| Name             | Value            |
| ---------------- | ---------------- |
| `ListMatchAllow` | allow            |
| `ListMatchBlock` | block            |
| `ListMatchNone`  | none             |