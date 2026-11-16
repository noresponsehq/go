# Certainty

confirmed: a linked record or the searched identifier states it directly. likely: a strong inference. possible: a weak inference.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v15/models/components"
)

value := components.CertaintyConfirmed

// Open enum: custom values can be created with a direct type cast
custom := components.Certainty("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `CertaintyConfirmed` | confirmed            |
| `CertaintyLikely`    | likely               |
| `CertaintyPossible`  | possible             |