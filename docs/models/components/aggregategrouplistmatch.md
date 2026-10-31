# AggregateGroupListMatch

The list action captured with the first answer. none means no matching entry was found. For retained_first_answers an absent field means the older assessment did not capture list state; this remains a separate group.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v10/models/components"
)

value := components.AggregateGroupListMatchBlock

// Open enum: custom values can be created with a direct type cast
custom := components.AggregateGroupListMatch("custom_value")
```


## Values

| Name                           | Value                          |
| ------------------------------ | ------------------------------ |
| `AggregateGroupListMatchBlock` | block                          |
| `AggregateGroupListMatchAllow` | allow                          |
| `AggregateGroupListMatchNone`  | none                           |