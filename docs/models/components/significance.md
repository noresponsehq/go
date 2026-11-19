# Significance

One rule ranks every change: identity changes (accounts, breaches, identifiers) are high, profile fields are medium, new posts are low.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v19/models/components"
)

value := components.SignificanceHigh

// Open enum: custom values can be created with a direct type cast
custom := components.Significance("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `SignificanceHigh`   | high                 |
| `SignificanceMedium` | medium               |
| `SignificanceLow`    | low                  |