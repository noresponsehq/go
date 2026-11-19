# ScreeningSpendLineRole

What the implementation did. analyser appears only on spend recorded before the engine stopped analysing posts.

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v18/models/components"
)

value := components.ScreeningSpendLineRoleResolver

// Open enum: custom values can be created with a direct type cast
custom := components.ScreeningSpendLineRole("custom_value")
```


## Values

| Name                              | Value                             |
| --------------------------------- | --------------------------------- |
| `ScreeningSpendLineRoleResolver`  | resolver                          |
| `ScreeningSpendLineRoleCollector` | collector                         |
| `ScreeningSpendLineRoleAnalyser`  | analyser                          |