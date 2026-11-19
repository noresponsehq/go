# Stage

## Example Usage

```go
import (
	"github.com/noresponsehq/go/v17/models/components"
)

value := components.StageLookup

// Open enum: custom values can be created with a direct type cast
custom := components.Stage("custom_value")
```


## Values

| Name                 | Value                |
| -------------------- | -------------------- |
| `StageLookup`        | lookup               |
| `StageLookupTyping`  | lookup_typing        |
| `StagePlatforms`     | platforms            |
| `StageWritten`       | written              |
| `StageLatePlatforms` | late_platforms       |
| `StageLateLookup`    | late_lookup          |
| `StageAiTyping`      | ai_typing            |
| `StageAnalysis`      | analysis             |
| `StageRating`        | rating               |